#include <Arduino.h>

#include "log.h"
#include "scanResult.h"

#include <CBOR.h>
#include <CBOR_streams.h>

namespace cbor = ::qindesign::cbor;

#define RUNS 250

void serializeData(size_t);

// FIXME: This should have a better size
uint8_t bytes[255]{0};
cbor::BytesStream bs{bytes, sizeof(bytes)};
cbor::BytesPrint bp{bytes, sizeof(bytes)};

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
  cbor::Writer cbor{bp};
  ScanResult res{ -count, "AA:BB:CC:DD:EE", "MyWifi"};

  bp.reset();

  cbor.writeTag(cbor::kSelfDescribeTag);
  cbor.beginArray(2);
  cbor.writeInt(res.RSSI);
  cbor.beginText(res.BSSID.length());
  cbor.writeBytes((const uint8_t *) res.BSSID.c_str(), res.BSSID.length());
  cbor.beginText(res.SSID.length());
  cbor.writeBytes((const uint8_t *) res.SSID.c_str(), res.SSID.length());

  logln(cbor.getWriteSize());
}

void loop() { }