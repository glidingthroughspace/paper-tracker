#pragma once

#include <models/scanResult.hpp>
#include <serialization/cbor/CBORDocument.hpp>
#include <log.hpp>
#include <vector>
#include <cstdint>

class TrackerResponse {
  public:
    TrackerResponse(uint8_t batteryPercentage, bool isCharging, uint64_t resultID, uint8_t batchCount, std::vector<ScanResult> scanResults);
    void toCBOR(CBORDocument&);
  private:
    CBORUint8 m_battery_percentage{"battery_percentage"};
    CBORBool m_is_charging{"is_charging"};
    // An ID for the result. This should be the same in every batch of the result.
    CBORUint64 m_result_id{"result_id"};
    // Indicates how many batches are in the response in total (how many the server should expect)
    CBORUint8 m_batch_count{"result_batch_count"};
    std::vector<ScanResult> m_scan_results;
};
