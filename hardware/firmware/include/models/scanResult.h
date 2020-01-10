#pragma once

#include <log.h>
#include <utils.h>
#include <serialization/cbor/CBORValue.h>
#include <serialization/cbor/CBORDocument.h>

#define SCAN_RESULT_SIZE_BYTES sizeof(ScanResult)

class ScanResult : public CBORSerializable {
  public:
    CBORInt32 RSSI{"RSSI"};
    CBORCString BSSID{"BSSID"};
    CBORCString SSID{"SSID"};

    void print() {
      log("SSID: ");
      log(SSID.get());
      log(", BSSID: ");
      log(BSSID.get());
      log(", RSSI: ");
      logln(static_cast<uint32_t>(RSSI.value));
    }

    bool toCBOR(CBORDocument& cbor) {
      cbor.addValue(RSSI);
      cbor.addValue(BSSID);
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
