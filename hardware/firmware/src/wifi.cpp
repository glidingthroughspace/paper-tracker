#include <wifi.hpp>

#include <log.hpp>
#include <power.hpp>

#include <cstdio>
#include <cstring>

#ifndef NDEBUG
#define WIFI_CONNECTION_DELAY 250
#else
#define WIFI_CONNECTION_DELAY 10
#endif

#include "esp_wpa2.h"

WIFI::WIFI() {
  // This improves performance. WiFi credentials are hard-coded when flashing,
  // therefore storing them again in flash does not make sense.
  WiFi.persistent(false);
  WiFi.disconnect(true);
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
  WiFi.mode(WIFI_STA);
  WiFi.begin(ssid, password);
  return connectLoop();
}

bool WIFI::connect(const char* ssid, const char* username, const char* password) {
  log("Connecting to WiFi SSID ");
  log(ssid);
  log(" with user ");
  logln(username);
  WiFi.mode(WIFI_STA);

  esp_wifi_sta_wpa2_ent_set_identity((uint8_t*)username, strlen(username));
  esp_wifi_sta_wpa2_ent_set_username((uint8_t*)username, strlen(username));
  esp_wifi_sta_wpa2_ent_set_password((uint8_t*)password, strlen(password));

  esp_wpa2_config_t config = WPA2_CONFIG_INIT_DEFAULT();

  if (esp_wifi_sta_wpa2_ent_enable(&config)) {
    logln("Failed to enable WPA2");
    return false;
  }

  logln("Initialized wifi config");
  WiFi.begin(ssid);
  return connectLoop();
}

bool WIFI::connectLoop() {
  unsigned int counter = 0;
  while(WiFi.status() != WL_CONNECTED) {
    log(WiFi.status());
    delay(WIFI_CONNECTION_DELAY);
    counter++;
    // TODO: We might increase this timeout after field-tests
    // Try for 10 seconds
    if (counter > 10000 / WIFI_CONNECTION_DELAY) {
      logln();
      logln("Connection failed, trying again in 10 seconds");
      Power::deep_sleep_for_seconds(10);
    }
  }
  logln();
  log("Connected, IP address is: ");
  logln(WiFi.localIP());
  udp.begin(LOCAL_UDP_PORT);
  return true;
}

uint8_t WIFI::getVisibleNetworkCount() const {
  return visibleNetworkCount;
}

uint8_t WIFI::scanVisibleNetworks() {
  logln("Scanning for networks...");
  visibleNetworkCount = WiFi.scanNetworks();
  log("Found ");
  log(visibleNetworkCount);
  logln(" networks in reach");
  return visibleNetworkCount;
}

std::vector<ScanResult> WIFI::getAllVisibleNetworks() {
  scanVisibleNetworks();
  std::vector<ScanResult> scanResults(0);
  for (auto i = 0; i < visibleNetworkCount; i++) {
    scanResults.emplace_back(WiFi.RSSI(i), WiFi.BSSIDstr(i), WiFi.SSID(i));
  }
  return scanResults;
}
