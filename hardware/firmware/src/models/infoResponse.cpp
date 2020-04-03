#include <models/infoResponse.hpp>

InfoResponse::InfoResponse(uint8_t batteryPercentage, bool isCharging) {
  m_battery_percentage.value = batteryPercentage;
  m_is_charging.value = isCharging;
}

void InfoResponse::toCBOR(CBORDocument& doc) {
  doc.begin_map(2);
  m_battery_percentage.serialize_to(doc);
  m_is_charging.serialize_to(doc);
}
