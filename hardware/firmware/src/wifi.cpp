#include <wifi.h>
#include <log.h>
#include <stdio.h>
#include <string.h>

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
  esp_wifi_sta_wpa2_ent_set_identity((uint8_t*)username, strlen(username));
  esp_wifi_sta_wpa2_ent_set_username((uint8_t*)username, strlen(username));
  esp_wifi_sta_wpa2_ent_set_password((uint8_t*)password, strlen(password));

  esp_wpa2_config_t config;
  config.crypto_funcs = &g_wifi_default_wpa2_crypto_funcs;
  logln("Initialized wifi config");
  if (esp_wifi_sta_wpa2_ent_enable(&config)) {
  logln("Failed to enable WPA2");
  return false;
  }

  WiFi.begin(ssid);

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
  visibleNetworkCount = WiFi.scanNetworks();
  return visibleNetworkCount;
}
