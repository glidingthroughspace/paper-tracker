#include <Arduino.h>
#include <WiFiUdp.h>
#include <IPAddress.h>

#include <log.hpp>
#include <models/scanResult.hpp>
#include <models/trackerResponse.hpp>
#include <models/infoResponse.hpp>
#include <serialization/cbor/CBORDocument.hpp>
#include <wifi.hpp>
#include <apiClient.hpp>
#include <power.hpp>
#include <storage.hpp>
#include <cmath>
#include <utils.hpp>

#include <credentials.hpp>

WIFI wifi;
ApiClient apiClient(&wifi.getUDP(), SERVER_IP);

void haltIf(bool condition, const char* message);
void sendScanResultsInChunks(std::vector<ScanResult>&);
void sendStatusInformation();
bool hasTrackerID();
void requestNextCommand();

static void onCommandReceived(Command& command) {
  command.print();

  switch (command.getType()) {
    case CommandType::SLEEP: {
        Power::deep_sleep_for(command.getSleepTime());
      }
      break;
    case CommandType::SEND_TRACKING_INFO: {
        auto scanResults = wifi.getAllVisibleNetworks();
        sendScanResultsInChunks(scanResults);
        if (command.getSleepTime() > seconds(0) && command.getSleepTime() <= seconds(10)) {
          utils::time::wait_for(command.getSleepTime());
        } else if (command.getSleepTime() > seconds(10)) {
          Power::deep_sleep_for(command.getSleepTime());
        } else {
          logln("[Main] Not sleeping, since sleep time is 0");
          requestNextCommand();
        }
      }
      break;
    case CommandType::SEND_INFO: {
        sendStatusInformation();
        Power::deep_sleep_for(command.getSleepTime());
      }
      break;
    default:
      // We already sleep & reset the tracker when deserializing the command, so this should never
      // be reached.
      logln("[Main] Unknown command");
  }
}

void setup() {
  // Initialize random source
  randomSeed(analogRead(0));
  Power::enable_powersavings();
  initSerial(115400);
  logln("[Main] Starting");

  Power::print_wakeup_reason();

  #ifndef NDEBUG
  Power::get_battery_percentage();
  Power::is_charging();
  pinMode(4, INPUT_PULLDOWN);
  auto should_clear_storage = digitalRead(4);
  if (should_clear_storage) {
    logln("[Main] Clearing storage since pin 4 is pulled high");
    Storage::clear();
  }
  #endif

  if (WIFI_USERNAME != nullptr)
    haltIf(!wifi.connect(WIFI_SSID, WIFI_USERNAME, WIFI_PASSWORD), "[Main] Failed to connect to WiFi");
  else
    haltIf(!wifi.connect(WIFI_SSID, WIFI_PASSWORD), "[Main] Failed to connect to WiFi");

  haltIf(!apiClient.start(), "[Main] Failed to start the API client");

  if (!hasTrackerID()) {
    logln("[Main] This tracker does not have an ID yet");
    apiClient.requestTrackerID([] (uint16_t newID) {
      Storage::set(Storage::TRACKER_ID, newID);
      requestNextCommand();
    });
  } else {
    requestNextCommand();
  }
}

void requestNextCommand() {
  auto id = Storage::get(Storage::TRACKER_ID);
  logf("[Main] This tracker has id %d\n", id);
  apiClient.requestNextCommand(id, onCommandReceived);
}

void loop() {
  apiClient.loop();
}

bool hasTrackerID() {
  return Storage::exists(Storage::TRACKER_ID);
}

void sendScanResultsInChunks(std::vector<ScanResult>& scanResults) {
  constexpr size_t batchSize = 10;
  uint8_t batchCount = floor(scanResults.size() / batchSize) + 1;
  uint64_t resultID = random(10, 100000);
  logf("[Main] Current resultID is %d\n", resultID);
  for (auto i = 0; i < scanResults.size(); i+=batchSize) {
    auto begin = scanResults.begin() + i;
    auto end = (i + batchSize > scanResults.size()) ? scanResults.end() : scanResults.begin() + i + batchSize;
    std::vector<ScanResult> batch(begin, end);

    auto percent = Power::get_battery_percentage();
    auto charging = Power::is_charging();
    TrackerResponse trackerResponse{percent, charging, resultID, batchCount, batch};
    CBORDocument cborDocument;
    trackerResponse.toCBOR(cborDocument);
    apiClient.writeTrackingData(Storage::get(Storage::TRACKER_ID), cborDocument.serialize(), [] () {
      logln("[Main] Sent scan results to server");
    });
  }
}

void sendStatusInformation() {
  InfoResponse infoResponse{Power::get_battery_percentage(), Power::is_charging()};
  CBORDocument cborDocument;
  infoResponse.toCBOR(cborDocument);
  apiClient.writeInfoResponse(Storage::get(Storage::TRACKER_ID), cborDocument.serialize());
}

void haltIf(bool condition, const char* message) {
  if (condition) {
    // TODO: Maybe blink the LED?
    logln("[Main] Setup action failed, stalling tracker!");
    while(true) {;}
  }
}
