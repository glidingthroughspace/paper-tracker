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
#include <storage.hpp>

#include <credentials.hpp>

WIFI wifi;
ApiClient apiClient(&wifi.getUDP(), IPAddress(172,16,15,230));

void haltIf(bool condition, const char* message);
void sendScanResultsInChunks(std::vector<ScanResult>&);
bool hasTrackerID();

static void onCommandReceived(Command& command) {
  command.print();

  switch (command.getType()) {
    case CommandType::SLEEP:
      Power::deep_sleep_for_seconds(command.getSleepTimeInSeconds());
      break;
    case CommandType::SEND_TRACKING_INFO: {
        auto scanResults = wifi.getAllVisibleNetworks();
        sendScanResultsInChunks(scanResults);
        Power::deep_sleep_for_seconds(command.getSleepTimeInSeconds());
      }
      break;
    default:
      // We already sleep & reset the tracker when deserializing the command, so this should never
      // be reached.
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

  Storage::clear();
  if (!hasTrackerID()) {
    logln("This tracker does not have an ID yet");
    apiClient.requestTrackerID([] (uint16_t newID) {
      Storage::set(Storage::TRACKER_ID, newID);
      apiClient.requestNextCommand(newID, onCommandReceived);
    });
  } else {
    auto id = Storage::get(Storage::TRACKER_ID);
    log("This tracker has id ");
    logln(id);
    apiClient.requestNextCommand(id, onCommandReceived);
  }
}

void loop() {
  apiClient.loop();
}

bool hasTrackerID() {
  return Storage::exists(Storage::TRACKER_ID);
}

void sendScanResultsInChunks(std::vector<ScanResult>& scanResults) {
  constexpr size_t batchSize = 10;
  for (auto i = 0; i < scanResults.size(); i+=batchSize) {
    auto begin = scanResults.begin() + i;
    auto end = (i + batchSize > scanResults.size()) ? scanResults.end() : scanResults.begin() + i + batchSize;
    std::vector<ScanResult> batch(begin, end);

    TrackerResponse trackerResponse{100, batch};
    CBORDocument cborDocument;
    trackerResponse.toCBOR(cborDocument);
    apiClient.writeTrackingData(Storage::get(Storage::TRACKER_ID), cborDocument.serialize(), [] () {
      logln("Sent scan results to server");
    });
  }
}

void haltIf(bool condition, const char* message) {
  if (condition) {
    // TODO: Maybe blink the LED?
    logln("Setup action failed, stalling tracker!");
    while(true) {;}
  }
}
