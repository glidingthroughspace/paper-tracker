#pragma once

#include <models/scanResult.hpp>
#include <serialization/cbor/CBORDocument.hpp>
#include <log.hpp>
#include <vector>
#include <cstdint>

class InfoResponse {
  public:
    InfoResponse(uint8_t batteryPercentage, bool isCharging);
    void toCBOR(CBORDocument&);
  private:
    CBORUint8 m_battery_percentage{"battery_percentage"};
    CBORBool m_is_charging{"is_charging"};
};
