#include <wifi.hpp>

#include <cstdio>
#include <cstring>

#include <log.hpp>
#include <power.hpp>
#include <utils.hpp>

#ifndef NDEBUG
#define WIFI_CONNECTION_DELAY 250
#else
#define WIFI_CONNECTION_DELAY 10
#endif

constexpr auto CONNECTION_TIMEOUT_SECONDS = 15;

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
  logf("[WiFi] Connecting to WiFi SSID %s\n", ssid);
  WiFi.mode(WIFI_STA);
  WiFi.begin(ssid, password);
  return connectLoop();
}

bool WIFI::connect(const char* ssid, const char* username, const char* password) {
  logf("[WiFi] Connecting to WiFi SSID %s with user %s\n", ssid, username);
  WiFi.mode(WIFI_STA);

  esp_wifi_sta_wpa2_ent_set_identity((uint8_t*)username, strlen(username));
  esp_wifi_sta_wpa2_ent_set_username((uint8_t*)username, strlen(username));
  esp_wifi_sta_wpa2_ent_set_password((uint8_t*)password, strlen(password));

  esp_wpa2_config_t config = WPA2_CONFIG_INIT_DEFAULT();

  if (esp_wifi_sta_wpa2_ent_enable(&config)) {
    logln("[WiFi] Failed to enable WPA2");
    return false;
  }

  logln("[WiFi] Initialized wifi config");
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
    if (counter > (CONNECTION_TIMEOUT_SECONDS * 1000) / WIFI_CONNECTION_DELAY) {
      logln();
      logln("[WiFi] Connection timeout reached");
      Power::deep_sleep_for(utils::time::seconds(10));
    }
  }
  logln();
  logf("[WiFi] Connected, IP address is %s\n", WiFi.localIP().toString().c_str());
  udp.begin(LOCAL_UDP_PORT);
  return true;
}

uint8_t WIFI::getVisibleNetworkCount() const {
  return visibleNetworkCount;
}

uint8_t WIFI::scanVisibleNetworks() {
  logln("[WiFi] Scanning for networks...");
  visibleNetworkCount = WiFi.scanNetworks();
  logf("[WiFi] Found %d access points in reach\n", visibleNetworkCount);
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
