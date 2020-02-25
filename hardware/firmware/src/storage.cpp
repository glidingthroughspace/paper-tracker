#include <storage.hpp>

#include <log.hpp>
#include <EEPROM.h>
#include <Preferences.h>

// Offset to be able to move all values to a new location to circumvent wear and the aid in
// development. All values are relative to this location.
constexpr int EEPROM_OFFSET = 0x00;

Preferences* Storage::instance = nullptr;

Preferences* Storage::prefs() {
  if (instance == nullptr) {
    instance = new Preferences();
    instance->begin("paptrack", false);
  }
  return instance;
}

uint16_t Storage::get(const char* key) {
  return (uint16_t) prefs()->getUInt(key, 0);
}

bool Storage::exists(const char* key) {
  // Since we only support 16 bit ints, but prefs uses 32 bits, a value can never be the default
  // value, if it ever has been set.
  return (prefs()->getUInt(key, 0xFFFFFFFF) != 0xFFFFFFFF);
}

bool Storage::set(const char* key, uint16_t value) {
  prefs()->putUInt(key, value);
  return true;
}

void Storage::clear() {
  prefs()->clear();
}
