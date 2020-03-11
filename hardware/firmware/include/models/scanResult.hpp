#pragma once

#include <log.hpp>
#include <serialization/cbor/CBORValue.hpp>
#include <serialization/cbor/CBORDocument.hpp>

class ScanResult {
  public:
    ScanResult(): RSSI{"rssi"}, BSSID{"bssid"}, SSID{"ssid"} {};
    ScanResult(int32_t rssi, const String& bssid, const String& ssid): RSSI{"rssi", rssi}, BSSID{"bssid", bssid}, SSID{"ssid", ssid} {};

    void print() {
      log("SSID: ");
      log(SSID.get());
      log(", BSSID: ");
      log(BSSID.get());
      log(", RSSI: ");
      logln(static_cast<int32_t>(RSSI.value));
    }

    void toCBOR(CBORDocument& cbor) {
      logln("Serializing scan result:");
      print();
      cbor.begin_map(3);
      RSSI.serialize_to(cbor);
      BSSID.serialize_to(cbor);
      SSID.serialize_to(cbor);
    }
  private:
    CBORInt32 RSSI;
    CBORString BSSID;
    CBORString SSID;
};
