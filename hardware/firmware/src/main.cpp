#include <Arduino.h>
#include <WiFiUdp.h>
#include <IPAddress.h>

#include <log.h>
#include <scanResult.h>
#include <wifi.h>
#include <apiClient.h>

#include <credentials.h>

#define SCAN_RESULT_BUFFER_SIZE 5
// FIXME: This number is not correct
#define SCAN_RESULT_MESSAGE_OVERHEAD 100


WIFI wifi;
ScanResult scanResultBuffer[SCAN_RESULT_BUFFER_SIZE];
ApiClient apiClient(wifi.getUDP(), IPAddress(192,168,43,111));

uint8_t bytes[SCAN_RESULT_BUFFER_SIZE * SCAN_RESULT_SIZE_BYTES + SCAN_RESULT_MESSAGE_OVERHEAD]{0};

void fail_loop(const char* message);

void setup() {
  initSerial(115400);
  logln("Starting");

  #ifdef WIFI_USERNAME
  if (!wifi.connect(WIFI_SSID, WIFI_USERNAME, WIFI_PASSWORD)) {
    fail_loop("Failed to connect to WiFi");
  }
  #else
  if (!wifi.connect(WIFI_SSID, WIFI_PASSWORD)) {
    fail_loop("Failed to connect to WiFi");
  }
  #endif

  if (!apiClient.start()) {
    fail_loop("Failed to start the API client");
  }

  apiClient.requestNextAction([] () {});

}

void loop() {
  apiClient.loop();
}

void fail_loop(const char* message) {
  // TODO: Maybe blink the LED?
  logln("Failed to start CoAP client! Stalling Tracker!");
  while(true) {;}
}