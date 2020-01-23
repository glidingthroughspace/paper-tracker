#pragma once

#include <log.hpp>
#include <serialization/cbor/CBORValue.hpp>
#include <serialization/cbor/CBORDocument.hpp>

class ScanResult {
  public:
		ScanResult(): RSSI{"RSSI"}, BSSID{"BSSID"}, SSID{"SSID"} {};
		ScanResult(int32_t rssi, const String& bssid, const String& ssid): RSSI{"RSSI", rssi}, BSSID{"BSSID", bssid}, SSID{"SSID", ssid} {};

    void print() {
      log("SSID: ");
      log(SSID.get());
      log(", BSSID: ");
      log(BSSID.get());
      log(", RSSI: ");
      logln(static_cast<uint32_t>(RSSI.value));
    }

    void toCBOR(CBORDocument& cbor) {
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
