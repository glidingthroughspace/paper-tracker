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

void haltIf(bool condition, const char* message);

void setup() {
  initSerial(115400);
  logln("Starting");

  const char cmdData[] = { 0xa2, 0x67, 0x43, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x02, 0x6c, 0x53, 0x6c, 0x65, 0x65, 0x70, 0x54, 0x69, 0x6d, 0x65, 0x53, 0x65, 0x63, 0x05, 0x0a};
  
  Command cmd;
  cmd.fromCBOR((uint8_t*)(cmdData), sizeof(cmdData));

  logln("Parsed command");
  
  return;

  #ifdef WIFI_USERNAME
  haltIf(!wifi.connect(WIFI_SSID, WIFI_USERNAME, WIFI_PASSWORD), "Failed to connect to WiFi");
  #else
  haltIf(!wifi.connect(WIFI_SSID, WIFI_PASSWORD), "Failed to connect to WiFi");
  #endif

  haltIf(!apiClient.start(), "Failed to start the API client");

  apiClient.requestNextCommand([] (Command& command) {
    log("Next Command is ");
    log((uint8_t) command.getType());
    log(" and sleep time in seconds is ");
    logln(command.getSleepTimeInSeconds());
  });

  // wifi.scanVisibleNetworks();
  // logln("Scanned for networks");
  // wifi.getVisibleNetworks(0, scanResultBuffer, SCAN_RESULT_BUFFER_SIZE);
  // TrackerResponse trackerResponse{0};
  // memcpy(scanResultBuffer, trackerResponse.scanResults, SCAN_RESULT_BUFFER_SIZE);
  // trackerResponse.toCBOR(bytes, sizeof(bytes));
  // apiClient.writeTrackingData(bytes, sizeof(bytes), [] () {});
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