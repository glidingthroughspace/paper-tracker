#include <models/trackerResponse.hpp>

TrackerResponse::TrackerResponse(uint8_t batteryPercentage, bool isCharging, bool isLastBatch, std::vector<ScanResult> scanResults) {
  m_battery_percentage.value = batteryPercentage;
  m_is_charging.value = isCharging;
  m_is_last_batch.value = isLastBatch;
  m_scan_results = std::move(scanResults);
}

void TrackerResponse::toCBOR(CBORDocument& doc) {
  // 1 (battery percentage) + 1 (scan results array)
  doc.begin_map(4);
  m_battery_percentage.serialize_to(doc);
  m_is_charging.serialize_to(doc);
  m_is_last_batch.serialize_to(doc);
  doc.begin_array(m_scan_results.size(), "scan_results");
  for (auto i = 0; i < m_scan_results.size(); i++) {
    m_scan_results[i].toCBOR(doc);
  }
}
