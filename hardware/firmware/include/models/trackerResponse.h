#pragma once

#ifndef SCAN_RESULT_BUFFER_SIZE
#define SCAN_RESULT_BUFFER_SIZE 5 // FIXME: Make this a compile-time parameter
#endif

#include <models/scanResult.h>
#include <CBOR.h>
#include <CBOR_streams.h>

struct TrackerResponse {
  uint8_t batteryPercentage;
  ScanResult scanResults[SCAN_RESULT_BUFFER_SIZE];

  void toCBOR(uint8_t* buffer, size_t bufferSize) {
    cbor::BytesPrint bp{buffer, bufferSize};
    cbor::Writer cbor{bp};
    cbor.writeTag(cbor::kSelfDescribeTag);
    cbor.writeUnsignedInt(batteryPercentage);
    cbor.beginArray(SCAN_RESULT_BUFFER_SIZE);
    for (size_t i = 0; i < SCAN_RESULT_BUFFER_SIZE; i++) {
      scanResults[i].toCBOR(cbor);
    }
  }
};
