#pragma once

#include <models/scanResult.hpp>
#include <serialization/cbor/CBORDocument.hpp>
#include <log.hpp>
#include <vector>
#include <cstdint>

class TrackerResponse {
  public:
    TrackerResponse(uint8_t batteryPercentage, std::vector<ScanResult> scanResults);
		void toCBOR(CBORDocument&);
	private:
    CBORUint8 m_battery_percentage;
    std::vector<ScanResult> m_scan_results;
};
