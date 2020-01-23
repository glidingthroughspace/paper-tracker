#include <models/trackerResponse.hpp>

TrackerResponse::TrackerResponse(uint8_t batteryPercentage, std::vector<ScanResult> scanResults): m_battery_percentage{"BatteryPercentage", batteryPercentage} {
	m_scan_results = std::move(scanResults);
}

void TrackerResponse::toCBOR(CBORDocument& doc) {
	// 1 (battery percentage) + 1 (scan results array)
	doc.begin_map(2);
	m_battery_percentage.serialize_to(doc);
	doc.begin_array(m_scan_results.size(), "ScanResults");
	for (auto i = 0; i < m_scan_results.size(); i++) {
		m_scan_results[i].toCBOR(doc);
	}
}
