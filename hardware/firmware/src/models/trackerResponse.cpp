#include <models/trackerResponse.hpp>

TrackerResponse::TrackerResponse(uint8_t batteryPercentage, bool isCharging, uint64_t resultID, uint8_t batchCount, std::vector<ScanResult> scanResults) {
  m_battery_percentage.value = batteryPercentage;
  m_is_charging.value = isCharging;
  m_batch_count.value = batchCount;
  m_result_id.value = resultID;
  m_scan_results = std::move(scanResults);
}

void TrackerResponse::toCBOR(CBORDocument& doc) {
  // 1 (battery percentage) + 1 (scan results array)
  doc.begin_map(5);
  m_battery_percentage.serialize_to(doc);
  m_is_charging.serialize_to(doc);
  m_result_id.serialize_to(doc);
  m_batch_count.serialize_to(doc);
  doc.begin_array(m_scan_results.size(), "scan_results");
  for (auto i = 0; i < m_scan_results.size(); i++) {
    m_scan_results[i].toCBOR(doc);
  }
}
