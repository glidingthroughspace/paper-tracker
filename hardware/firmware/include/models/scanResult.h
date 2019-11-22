#pragma once

#include <log.h>
#include <utils.h>
#include <CBOR.h>
#include <CBOR_streams.h>

namespace cbor = ::qindesign::cbor;

#define SSID_LENGTH 32
#define BSSID_LENGTH 6
// Two characters per byte + colons between bytes + NULL terminator
#define BSSID_STRING_LENGTH BSSID_LENGTH * 2 + BSSID_LENGTH

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
    bssid_to_string(buf);
    log(buf);
    log(", RSSI: ");
    logln(RSSI);
  }

  /**
   * Converts the BSSID to a C string. 
   * @param buf An 18 bytes long buffer to fill with the BSSID. Use the BSSID_STRING_LENGTH constant
   * to allocate your buffer.
   */
  void bssid_to_string(char* buf) {
    // A  A  :  B  B  :  C  C  :  D  D  :  E  E  :  F  F
    // 0  1  2  3  4  5  6  7  8  9 10 11 12 13 14 15 16
    buf[2] = buf[5] = buf[8] = buf[11] = buf[14] = ':';
    for (int i = 0; i < BSSID_LENGTH; i++) {
      utils::byte_to_hex(BSSID[i], &buf[i * 3]);
    }
    buf[17] = '\0';
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
