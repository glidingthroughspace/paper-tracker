#include <storage.hpp>

#include <EEPROM.h>

// Offset to be able to move all values to a new location to circumvent wear and the aid in
// development. All values are relative to this location.
constexpr int EEPROM_OFFSET = 0x00;

uint16_t Storage::get(Value value) {
  auto msb = EEPROM.read(EEPROM_OFFSET + (uint8_t)value);
  auto lsb = EEPROM.read(EEPROM_OFFSET + (uint8_t)value + 1);
  return (msb << 8) + lsb;
}

bool Storage::exists(Value value) {
  return (Storage::get(value) != 0xFFFF);
}

bool Storage::set(Value value, uint16_t newValue) {
  // Since reading does not wear the EEPROM, but writing does, check if the current value equals the
  // new value first.
  if (Storage::get(value) == newValue) {
    return false;
  }
  EEPROM.write(EEPROM_OFFSET + (uint8_t)value, (newValue) >> 8);
  EEPROM.write(EEPROM_OFFSET + (uint8_t)value + 1, (newValue) & 0xFF);
  return true;
}
