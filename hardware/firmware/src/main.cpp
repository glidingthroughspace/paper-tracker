#include <Arduino.h>
#include <WiFiUdp.h>
#include <IPAddress.h>

#include <log.h>
#include <models/scanResult.h>
#include <models/trackerResponse.h>
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
  haltIf(!wifi.connect(WIFI_SSID, WIFI_USERNAME, WIFI_PASSWORD), "Failed to connect to WiFi");
  #else
  haltIf(!wifi.connect(WIFI_SSID, WIFI_PASSWORD), "Failed to connect to WiFi");
  #endif

  haltIf(!apiClient.start(), "Failed to start the API client");

  apiClient.requestNextAction([] (CoapPacket& responsePacket) {
    if (ApiClient::isErrorResponse(responsePacket)) {
      logln("Request for next action returned an error");
      logln(responsePacket.code);
      return;
    }
    // Deserialize the action, do something with it
  });

  wifi.scanVisibleNetworks();
  logln("Scanned for networks");
  wifi.getVisibleNetworks(0, scanResultBuffer, SCAN_RESULT_BUFFER_SIZE);
  TrackerResponse trackerResponse{0};
  apiClient.requestNextAction([] () {});
  memcpy(scanResultBuffer, trackerResponse.scanResults, SCAN_RESULT_BUFFER_SIZE);
  trackerResponse.toCBOR(bytes, sizeof(bytes));
  apiClient.writeTrackingData(bytes, sizeof(bytes), [] () {});
}

void loop() {
  apiClient.loop();
}

void haltIf(bool condition, const char* message) {
  if (condition) {
    // TODO: Maybe blink the LED?
    logln("Failed to start CoAP client! Stalling Tracker!");
    while(true) {;}
  }
}