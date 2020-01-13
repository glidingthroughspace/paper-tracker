#pragma once

#include <log.h>
#include <utils.h>
#include <serialization/cbor/CBORValue.h>
#include <serialization/cbor/CBORDocument.h>

#define SCAN_RESULT_SIZE_BYTES sizeof(ScanResult)

class ScanResult : public CBORSerializable {
  public:
    CBORInt32 RSSI{"RSSI"};    // 4 byte + 4 byte = 8 byte
    CBORString BSSID{"BSSID"}; // 5 byte + 17 byte = 22 byte
    CBORString SSID{"SSID"};   // 4 byte + 16 byte = 20 byte
                               // = 50 byte

    void print() {
      log("SSID: ");
      log(SSID.get());
      log(", BSSID: ");
      log(BSSID.get());
      log(", RSSI: ");
      logln(static_cast<uint32_t>(RSSI.value));
    }

    bool toCBOR(CBORDocument& cbor) {
      logln("serializing bssid");
      cbor.addValue(BSSID);
      logln("serializing rssi");
      cbor.addValue(RSSI);
      logln("serializing ssid");
      cbor.addValue(SSID);
      return true;
    }

    bool toCBOR(cbor::Writer& cbor) {
      RSSI.serializeTo(cbor);
      BSSID.serializeTo(cbor);
      SSID.serializeTo(cbor);
      return true;
    }
  
};
