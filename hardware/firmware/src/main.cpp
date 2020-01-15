#include <Arduino.h>
#include <WiFiUdp.h>
#include <IPAddress.h>

#include <log.h>
#include <models/scanResult.h>
#include <models/trackerResponse.h>
#include <wifi.h>
#include <apiClient.h>
#include <power.h>

#include <credentials.h>

#define SCAN_RESULT_BUFFER_SIZE 5
// FIXME: This number is not correct
#define SCAN_RESULT_MESSAGE_OVERHEAD 100

constexpr uint64_t ONE_SECOND_IN_MICROSECONDS = 1000 * 1000;

WIFI wifi(SCAN_RESULT_BUFFER_SIZE);
ScanResult scanResultBuffer[SCAN_RESULT_BUFFER_SIZE];
ApiClient apiClient(wifi.getUDP(), IPAddress(192,168,43,111));

uint8_t bytes[SCAN_RESULT_BUFFER_SIZE * SCAN_RESULT_SIZE_BYTES + SCAN_RESULT_MESSAGE_OVERHEAD]{0};

void haltIf(bool condition, const char* message);

static void onCommandReceived(Command& command) {
  log("Next Command is ");
  log((uint8_t) command.getType());
  log(" and sleep time in seconds is ");
  logln(command.getSleepTimeInSeconds());

  switch (command.getType()) {
    case CommandType::SLEEP: {
      Power::deep_sleep_for_seconds(command.getSleepTimeInSeconds());
    } break;
    case CommandType::SEND_TRACKING_INFO: {
      wifi.getAllVisibleNetworks([] (ScanResult* scan_results, size_t scan_result_count) {
        TrackerResponse<SCAN_RESULT_BUFFER_SIZE> trackerResponse;
        memcpy(scan_results, trackerResponse.scanResults, scan_result_count);
        trackerResponse.toCBOR(bytes, sizeof(bytes));
        apiClient.writeTrackingData(bytes, sizeof(bytes), [] () {
          logln("Sent scan results to server");
        });
      });
    } break;
    default:
      logln("Unknown command");
  }
}

void setup() {
  Power::enable_powersavings();
  initSerial(115400);
  logln("Starting");

  Power::print_wakeup_reason();

  #ifdef WIFI_USERNAME
  haltIf(!wifi.connect(WIFI_SSID, WIFI_USERNAME, WIFI_PASSWORD), "Failed to connect to WiFi");
  #else
  haltIf(!wifi.connect(WIFI_SSID, WIFI_PASSWORD), "Failed to connect to WiFi");
  #endif

  haltIf(!apiClient.start(), "Failed to start the API client");

  apiClient.requestNextCommand(onCommandReceived);
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
