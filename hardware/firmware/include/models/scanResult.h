#pragma once

#include <log.h>
#include <utils.h>
#include <CBOR.h>
#include <CBOR_streams.h>

namespace cbor = ::qindesign::cbor;

#define SSID_LENGTH 32

#define SCAN_RESULT_SIZE_BYTES sizeof(ScanResult)

struct ScanResult {
	int32_t RSSI;
  uint8_t BSSID[BSSID_LENGTH];
  char SSID[SSID_LENGTH];

  void print() {
    log("SSID: ");
    log(SSID);
    log(", BSSID: ");
    char buf[BSSID_STRING_LENGTH];
    utils::bssid_to_string(BSSID, buf);
    log(buf);
    log(", RSSI: ");
    logln(RSSI);
  }

  bool toCBOR(cbor::Writer& cbor) {
    cbor.writeTag(cbor::kSelfDescribeTag);
    cbor.writeInt(RSSI);
    cbor.beginBytes(BSSID_LENGTH); // Is this correct?
    cbor.writeBytes(BSSID, BSSID_LENGTH);
    cbor.beginText(SSID_LENGTH);
    cbor.writeBytes((uint8_t*) SSID, SSID_LENGTH);
    return true;
  }


  bool toCBOR(uint8_t* buffer, size_t bufferSize) {
    cbor::BytesPrint bp{buffer, bufferSize};
    cbor::Writer cbor{bp};
    cbor.writeTag(cbor::kSelfDescribeTag);
    cbor.writeInt(RSSI);
    cbor.beginBytes(BSSID_LENGTH); // Is this correct?
    cbor.writeBytes(BSSID, BSSID_LENGTH);
    cbor.beginText(SSID_LENGTH);
    cbor.writeBytes((uint8_t*) SSID, SSID_LENGTH);
    return true;
  }
};
