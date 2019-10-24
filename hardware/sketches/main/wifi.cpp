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

};

Wifi::~Wifi() {
   if (connected) {
     WiFi.disconnect();
   }
};

std::vector<ScanResult> Wifi::getAvailableNetworks() {
  std::vector<ScanResult> networks;
  int numberOfNetworks = WiFi.scanNetworks();
  for (int i = 0; i < numberOfNetworks; i++) {
    ScanResult net;
    net.SSID = WiFi.SSID(i);
    net.MAC = WiFi.BSSIDstr(i);
    net.signalStrength = WiFi.RSSI(i);
    networks.push_back(net);
  }

  // the printNetworks function is relatively expensive, so only do this in debug mode
  #ifndef NDEBUG
  printNetworks(networks);
  #endif

  return networks;
}

String padRight(String input, int wantedLength) {
  while (input.length() < wantedLength) {
    input.concat(' ');
  }
  return input;
}

void Wifi::printNetworks(const std::vector<ScanResult>& networks) {
  Log::println("The following networks were found:");
  // Find the SSID with the maximum length
  int maxLength = 0;
  for (auto net : networks) {
    if (maxLength < net.SSID.length()) {
      maxLength = net.SSID.length();
    }
  }
  Log::println(padRight("SSID", maxLength) + '\t' + "MAC             " + '\t' + "RSSID");
  for (auto net: networks) {
    Log::print(padRight(net.SSID, maxLength));
    Log::print('\t');
    Log::print(net.MAC);
    Log::print('\t');
    Log::print(net.signalStrength);
    Log::println();
  }
}

void Wifi::connect(const char* SSID, const char* password) {
  Log::print("Connecting to WiFi network with SSID ");
  Log::print(SSID);
  Log::println("...");
  WiFi.begin(SSID, password);
  while (WiFi.status() != WL_CONNECTED) {
    Log::print('.');
  }
}

void Wifi::connectDot1X(const char* ssid, const char* username, const char* password) {
  // WPA2 Connection starts here
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
  // WPA2 Connection ends here

  // Wait for connection AND IP address from DHCP
  Log::println();
  Log::println("Waiting for connection and IP Address from DHCP");
  while (WiFi.status() != WL_CONNECTED) {
    delay(500);
    Log::print(".");
  }
  Log::println("");
  Log::println("WiFi connected");
  Log::println("IP address: ");
  Log::println(WiFi.localIP());
}