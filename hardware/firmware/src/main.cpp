#include <Arduino.h>

#include "log.h"
#include "scanResult.h"

#include <ArduinoJson.h>

#define RUNS 250

void serializeData(size_t);

void setup() {
  initSerial(115400);
  logln("Starting");
  auto start = millis();
  for (size_t i = 0; i < RUNS; i++) {
    serializeData(i);
  }
  auto end = millis();
  logln();
  log("Serialized ");
  log(RUNS);
  log(" scan results in ");
  log(end - start);
  logln("ms");
}

void serializeData(size_t count) {
  DynamicJsonDocument doc(JSON_OBJECT_SIZE(3));
  ScanResult res{-count, "AA:BB:CC:DD:EE", "MyWifi"};
  res.print();
  doc["RSSI"] = res.RSSI;
  doc["BSSID"] = res.BSSID;
  doc["SSID"] = res.SSID;
  serializeJson(doc, Serial);
}

void loop() { }