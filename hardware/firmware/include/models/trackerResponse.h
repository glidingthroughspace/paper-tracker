#pragma once

#include <models/scanResult.h>
#include <serialization/cbor/CBORDocument.h>
#include <serialization/cbor/CBORValue.h>
#include <serialization/cbor/CBORArray.h>

template <size_t m_scan_result_count>
class TrackerResponse {
  public:
    CBORUint8 batteryPercentage{"BatteryPercentage"};
    ScanResult scanResults[m_scan_result_count];

    void toCBOR(uint8_t* buffer, size_t bufferSize) {
      // TODO: Proper size for the cbor document
    // ScanResult is about 60 bytes in CBOR, add some leeway
      StaticCBORDocument<m_scan_result_count * 60 + 30> doc;
      doc.addValue(batteryPercentage);

      StaticCBORArray<m_scan_result_count, ScanResult> scanResultsCBOR{"ScanResults"};
      for (auto i = 0; i < m_scan_result_count; i++) {
        scanResultsCBOR[i] = scanResults[i];
      }

      doc.addValue(scanResultsCBOR);
      memcpy(doc.serialize(), buffer, bufferSize);
    }
};
