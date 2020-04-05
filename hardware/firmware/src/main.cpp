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
        Power::deep_sleep_for_seconds(command.getSleepTimeInSeconds());
      }
      break;
    case CommandType::SEND_TRACKING_INFO: {
        auto scanResults = wifi.getAllVisibleNetworks();
        sendScanResultsInChunks(scanResults);
        if (command.getSleepTimeInSeconds() > 0 && command.getSleepTimeInSeconds() <= 10) {
          utils::time::wait_for_seconds(command.getSleepTimeInSeconds());
        } else if (command.getSleepTimeInSeconds() > 10) {
          Power::deep_sleep_for_seconds(command.getSleepTimeInSeconds());
        } else {
          logln("Not sleeping, since sleep time is 0");
          requestNextCommand();
        }
      }
      break;
    case CommandType::SEND_INFO: {
        sendStatusInformation();
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
  randomSeed(analogRead(0));
  Power::enable_powersavings();
  initSerial(115400);
  logln("Starting");

  Power::print_wakeup_reason();

  #ifndef NDEBUG
  Power::get_battery_percentage();
  Power::is_charging();
  pinMode(4, INPUT_PULLDOWN);
  auto should_clear_storage = digitalRead(4);
  if (should_clear_storage) {
    logln("Clearing storage since pin 4 is pulled high");
    Storage::clear();
  }
  #endif

  #ifdef WIFI_USERNAME
  haltIf(!wifi.connect(WIFI_SSID, WIFI_USERNAME, WIFI_PASSWORD), "Failed to connect to WiFi");
  #else
  haltIf(!wifi.connect(WIFI_SSID, WIFI_PASSWORD), "Failed to connect to WiFi");
  #endif

  haltIf(!apiClient.start(), "Failed to start the API client");

  if (!hasTrackerID()) {
    logln("This tracker does not have an ID yet");
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
  log("This tracker has id ");
  logln(id);
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
  log("Current resultID is ");
  logln(static_cast<uint32_t>(resultID)); // rand() only returns an int anyway, so this is safe
  for (auto i = 0; i < scanResults.size(); i+=batchSize) {
    auto begin = scanResults.begin() + i;
    auto end = (i + batchSize > scanResults.size()) ? scanResults.end() : scanResults.begin() + i + batchSize;
    std::vector<ScanResult> batch(begin, end);

    auto percent = Power::get_battery_percentage();
    auto charging = Power::is_charging();
    log("Battery percentage is ");
    log(percent);
    log(" and charging state is ");
    logln(charging);
    TrackerResponse trackerResponse{percent, charging, resultID, batchCount, batch};
    CBORDocument cborDocument;
    trackerResponse.toCBOR(cborDocument);
    apiClient.writeTrackingData(Storage::get(Storage::TRACKER_ID), cborDocument.serialize(), [] () {
      logln("Sent scan results to server");
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
    logln("Setup action failed, stalling tracker!");
    while(true) {;}
  }
}
