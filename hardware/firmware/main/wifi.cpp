#include "wifi.h"

#include <WiFiClient.h>
#include <WiFiClientSecure.h>
#include "log.h"
#include "credentials.h"

Wifi::Wifi() {};

Wifi::~Wifi() {
   if (getStatus() == STATION_GOT_IP) {
    ETS_UART_INTR_DISABLE();
    wifi_station_disconnect();
    ETS_UART_INTR_ENABLE();
    setSTAMode(false);
   }
};

bool Wifi::isConnected() const { 
  return WiFi.isConnected();
}

int Wifi::getVisibleNetworks() {
  numVisibleNetworks = WiFi.scanNetworks();

  #ifndef NDEBUG
  Log::print("Found ");
  Log::print(numVisibleNetworks);
  Log::println(" networks in reach");
  #endif

  return numVisibleNetworks;
}

int Wifi::getVisibleNetworkBatch(WifiNetwork* results, const int size, const int offset) const {
  int i;
  for (i = 0; (i < size) && (offset + i < numVisibleNetworks); i++) {
    WifiNetwork result;
    result.SSID = WiFi.SSID(offset + i);
    result.RSSI = WiFi.RSSI(offset + i);
    result.BSSID = WiFi.BSSIDstr(offset + i);
    results[i] = result;
    result.print();
  }
  return i;
}

int Wifi::getVisibleNetworkCount() const {
  return numVisibleNetworks;
}

void Wifi::connect() {
  if (WIFI_IS_DOT1X == 1) {
    connectDot1X();
  } else {
    connectWPA2();
  }
}

bool Wifi::setMode(uint8 m) const {
  if(wifi_get_opmode() == m){
    return true;
  }

  bool ret = false;

  if (m != WIFI_STA && m != WIFI_AP_STA)
    wifi_station_dhcpc_stop(); // It is save to stop the dhcp server, even if it hasn't been started

  ETS_UART_INTR_DISABLE();
  ret = wifi_set_opmode_current(m);
  ETS_UART_INTR_ENABLE();

  return ret;
}

bool Wifi::setSTAMode(bool enable) const {
  uint8 currentMode = wifi_get_opmode();
  bool isEnabled = ((currentMode & WIFI_STA) != 0);

  if(isEnabled != enable) {
    if(enable) {
      return setMode((currentMode | WIFI_STA));
    } else {
      return setMode((currentMode & (~WIFI_STA)));
    }
  } else {
      return true;
  }

}

void Wifi::connectWPA2() {
  Log::print("Connecting to WiFi network with SSID ");
  Log::print(WIFI_SSID);
  if (WIFI_USE_STATIC_CONFIG == 1) {
    Log::print(" using static IP ");
    Log::print(WIFI_STATIC_IP.toString());
    setStaticConfig(WIFI_STATIC_IP, WIFI_STATIC_GATEWAY, WIFI_STATIC_SUBNET_MASK);
  }
  Log::println("...");
  if(!setSTAMode(true)) {
    return;
  }
  auto passphraseLen = strlen(WIFI_PASSWORD);

  struct station_config conf;
  conf.threshold.authmode = (passphraseLen == 0) ? AUTH_OPEN : AUTH_WPA_PSK;

  if(strlen(WIFI_SSID) == 32)
    memcpy(reinterpret_cast<char*>(conf.ssid), WIFI_SSID, 32); //copied in without null term
  else
    strcpy(reinterpret_cast<char*>(conf.ssid), WIFI_SSID);

  if (passphraseLen == 64) // it's not a passphrase, is the PSK, which is copied into conf.password without null term
    memcpy(reinterpret_cast<char*>(conf.password), WIFI_PASSWORD, 64);
  else
    strcpy(reinterpret_cast<char*>(conf.password), WIFI_PASSWORD);

  conf.threshold.rssi = -127;

  ETS_UART_INTR_DISABLE();
  wifi_station_set_config_current(&conf);
  wifi_station_connect();
  ETS_UART_INTR_ENABLE();

  if(!WIFI_USE_STATIC_CONFIG) {
    wifi_station_dhcpc_start();
  }

  connectLoop();
}

bool Wifi::setStaticConfig(const unsigned long ip, const unsigned long gateway, const unsigned long subnet) const {
  struct ip_info info;
  info.ip.addr = ip;
  info.gw.addr = gateway;
  info.netmask.addr = subnet;

  wifi_station_dhcpc_stop();
  if(!wifi_set_ip_info(STATION_IF, &info))
    return false;

  // TODO: Investigate wether we need this. It is in ESP8266WiFiSTA.cpp
  // #if LWIP_VERSION_MAJOR != 1 && !CORE_MOCK
  //   // trigger address change by calling lwIP-v1.4 api
  //   // (see explanation in ESP8266WiFiSTA.cpp)
  //   netif_set_addr(eagle_lwip_getif(STATION_IF), &info.ip, &info.netmask, &info.gw);
  // #endif

  return true;
}


bool Wifi::setStaticConfig(const unsigned long ip, const unsigned long gateway, const unsigned long subnet, const unsigned long dns) const {
  setStaticConfig(ip, gateway, subnet);

  // TODO: Would we want to set a static DNS server?
}


void Wifi::connectDot1X() {
  Log::print("Connecting to 802.1X network with SSID ");
  Log::print(WIFI_SSID);
  Log::println("...");
  // Setting ESP into STATION mode only (no AP mode or dual mode)
  wifi_set_opmode(STATION_MODE);
  struct station_config wifi_config;
  memset(&wifi_config, 0, sizeof(wifi_config));
  strcpy((char*)wifi_config.ssid, WIFI_SSID);
  wifi_station_set_config(&wifi_config);
  wifi_station_clear_cert_key();
  wifi_station_clear_enterprise_ca_cert();
  wifi_station_set_wpa2_enterprise_auth(1);
  wifi_station_set_enterprise_identity((uint8_t*)WIFI_USERNAME, strlen(WIFI_USERNAME));
  wifi_station_set_enterprise_username((uint8_t*)WIFI_USERNAME, strlen(WIFI_USERNAME));
  wifi_station_set_enterprise_password((uint8_t*)WIFI_PASSWORD, strlen(WIFI_PASSWORD));
  wifi_station_connect();

  connectLoop();
}

station_status_t Wifi::getStatus() const {
  return wifi_station_get_connect_status();
}

void Wifi::connectLoop() {
  while (getStatus() != STATION_GOT_IP) {
    delay(1);
  }
  Log::println();
  Log::println("Connected");
  Log::print("IP address: ");
  struct ip_info ip;
  wifi_get_ip_info(STATION_IF, &ip);
  Log::println(IPAddress(ip.ip.addr));
}
