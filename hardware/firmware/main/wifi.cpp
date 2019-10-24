#include "wifi.h"

#include <ESP8266WiFi.h>
#include <WiFiClient.h>
#include <WiFiClientSecure.h>

extern "C" {
  #include "user_interface.h"
  #include "wpa2_enterprise.h"
}

#include "log.h"

Wifi::Wifi() : connected(false) {
  // This saves time when connecting, and isn't needed, see https://www.instructables.com/id/ESP8266-Pro-Tips/
  WiFi.persistent(false);
};

Wifi::~Wifi() {
   if (connected) {
     WiFi.disconnect();
   }
};

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
    result.SSID = WiFi.SSID(offset + i).c_str();
    result.RSSI = WiFi.RSSI(offset + i);
    result.BSSID = WiFi.BSSIDstr(offset + i).c_str();
    results[i] = result;
    result.print();
  }
  return i;
}

int Wifi::getVisibleNetworkCount() const {
  return numVisibleNetworks;
}

void Wifi::connect(const char* SSID, const char* password) {
  Log::print("Connecting to WiFi network with SSID ");
  Log::print(SSID);
  Log::println("...");
  WiFi.begin(SSID, password);
  while (WiFi.status() != WL_CONNECTED) {
    Log::print('.');
    delay(500);
  }
  Log::println();
  Log::println("Connected");
  Log::print("IP address: ");
  Log::println(WiFi.localIP());
  connected = true;
}

void Wifi::connectDot1X(const char* ssid, const char* username, const char* password) {
  // Setting ESP into STATION mode only (no AP mode or dual mode)
  wifi_set_opmode(STATION_MODE);
  struct station_config wifi_config;
  memset(&wifi_config, 0, sizeof(wifi_config));
  strcpy((char*)wifi_config.ssid, ssid);
  wifi_station_set_config(&wifi_config);
  wifi_station_clear_cert_key();
  wifi_station_clear_enterprise_ca_cert();
  wifi_station_set_wpa2_enterprise_auth(1);
  wifi_station_set_enterprise_identity((uint8_t*)username, strlen(username));
  wifi_station_set_enterprise_username((uint8_t*)username, strlen(username));
  wifi_station_set_enterprise_password((uint8_t*)password, strlen(password));
  wifi_station_connect();

  // Wait for connection AND IP address from DHCP
  Log::println();
  Log::print("Connecting to 802.1X network with SSID ");
  Log::print(ssid);
  Log::println("...");
  while (WiFi.status() != WL_CONNECTED) {
    delay(500);
    Log::print(".");
  }
  Log::println();
  Log::println("Connected");
  Log::print("IP address: ");
  Log::println(WiFi.localIP());
  connected = true;
}