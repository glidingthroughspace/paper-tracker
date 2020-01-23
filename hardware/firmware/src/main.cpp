#include <Arduino.h>
#include <WiFiUdp.h>
#include <IPAddress.h>

#include <log.hpp>
#include <models/scanResult.hpp>
#include <models/trackerResponse.hpp>
#include <serialization/cbor/CBORDocument.hpp>
#include <wifi.hpp>
#include <apiClient.hpp>
#include <power.hpp>

#include <credentials.hpp>

constexpr uint64_t ONE_SECOND_IN_MICROSECONDS = 1000 * 1000;

WIFI wifi;
ApiClient apiClient(wifi.getUDP(), IPAddress(192,168,43,91));

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
      auto scanResults = wifi.getAllVisibleNetworks();
      TrackerResponse trackerResponse{100, scanResults};
			CBORDocument cborDocument;
      trackerResponse.toCBOR(cborDocument);
			auto bytes = cborDocument.bytes();
			auto size = cborDocument.size();
			logln();
			for (auto i = 0; i < size; i++) {
				log(bytes[i]);
				log(" ");
			}
			logln();
			apiClient.writeTrackingData(cborDocument.bytes(), cborDocument.size(), [] () {
          logln("Sent scan results to server");
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
