#pragma once

#ifndef SCAN_RESULT_BUFFER_SIZE
#define SCAN_RESULT_BUFFER_SIZE 5 // FIXME: Make this a compile-time parameter
#endif

#include <models/scanResult.h>
#include <serialization/cbor/CBORDocument.h>
#include <serialization/cbor/CBORValue.h>
#include <serialization/cbor/CBORArray.h>

struct TrackerResponse {
  CBORUint8 batteryPercentage{"BatteryPercentage"};
  ScanResult scanResults[SCAN_RESULT_BUFFER_SIZE];

  void toCBOR(uint8_t* buffer, size_t bufferSize) {
    // TODO: Proper size for the cbor document
    StaticCBORDocument<1000> doc;
    doc.addValue(batteryPercentage);

    StaticCBORArray<SCAN_RESULT_BUFFER_SIZE, ScanResult> scanResultsCBOR{"ScanResults"};
    for (auto i = 0; i < SCAN_RESULT_BUFFER_SIZE; i++) {
      scanResultsCBOR[i] = scanResults[i];
    }

    doc.addValue(scanResultsCBOR);
    memcpy(doc.serialize(), buffer, bufferSize);
  }
};
