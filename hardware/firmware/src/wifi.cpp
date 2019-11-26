#include <wifi.h>
#include <ESP8266WiFi.h>
#include <log.h>
#include <stdio.h>
#include <string.h>

#ifndef NDEBUG
#define WIFI_CONNECTION_DELAY 250
#else
#define WIFI_CONNECTION_DELAY 10
#endif

extern "C" {
  #include "user_interface.h"
  #include "wpa2_enterprise.h"
}

WIFI::WIFI() {
  // This improves performance. WiFi credentials are hard-coded when flashing,
  // therefore storing them again in flash does not make sense.
  WiFi.persistent(false);
}

WIFI::~WIFI() {
  WiFi.disconnect();
}

WiFiUDP& WIFI::getUDP() {
  return udp;
}

bool WIFI::connect(const char* ssid, const char* password) {
  log("Connecting to WiFi SSID ");
  logln(ssid);
  WiFi.begin(ssid, password);
  return connectLoop();
}

bool WIFI::connect(const char* ssid, const char* username, const char* password) {
  log("Connecting to WiFi SSID ");
  log(ssid);
  log(" with user ");
  logln(username);
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

  return connectLoop();
}

bool WIFI::connectLoop() {
  while(WiFi.status() != WL_CONNECTED) {
    log('.');
    delay(WIFI_CONNECTION_DELAY);
  }
  logln();
  log("Connected, IP address is: ");
  logln(WiFi.localIP());
  udp.begin(LOCAL_UDP_PORT);
  return true;
}

uint8_t WIFI::getVisibleNetworks(uint8_t startAt, ScanResult* buffer, uint8_t bufferSize) {
  const uint8_t networkCount = getVisibleNetworkCount();
  int i;
  for (i = startAt; i < startAt + bufferSize && i < networkCount; i++) {
    ScanResult result{};
    result.RSSI = WiFi.RSSI(i);
    memcpy(result.BSSID, WiFi.BSSID(i), BSSID_LENGTH);
    strcpy(result.SSID, WiFi.SSID(i).c_str());
    buffer[i - startAt] = result;
  }
  return i - startAt;
}

uint8_t WIFI::getVisibleNetworkCount() const {
  return visibleNetworkCount;
}

uint8_t WIFI::scanVisibleNetworks() {
  logln("Scanning for networks...");
  visibleNetworkCount = WiFi.scanNetworks();
  return visibleNetworkCount;
}
