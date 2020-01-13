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
  WiFi.disconnect();
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
  esp_wifi_sta_wpa2_ent_set_identity((uint8_t*)username, strlen(username));
  esp_wifi_sta_wpa2_ent_set_username((uint8_t*)username, strlen(username));
  esp_wifi_sta_wpa2_ent_set_password((uint8_t*)password, strlen(password));

  esp_wpa2_config_t config = WPA2_CONFIG_INIT_DEFAULT();
  WiFi.mode(WIFI_STA);

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
    log(WiFi.status());
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
    result.RSSI.value = WiFi.RSSI(i);
    result.BSSID.set(WiFi.BSSIDstr(i));
    result.SSID.set(WiFi.SSID(i));
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
