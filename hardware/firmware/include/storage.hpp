#pragma once

#include <cstdint>
#include <Preferences.h>

/**
 * This class is a wrapper around the Arduino EEPROM library.
 * It adds an additional function that checks if a value exists in the EEPROM.
 * Since values that were never written contain the value 0xFF, this library assumes uint16s that
 * are 0xFFFF to be non-existant.
 */
class Storage {
  public:
    static constexpr const char* TRACKER_ID = "trackerid";
  public:
    // TODO: Make this private
    static Preferences* instance;
    static Preferences* prefs();
    static uint16_t get(const char* key);
    static bool exists(const char* key);
    static bool set(const char* key, uint16_t);
    static void clear();
};
