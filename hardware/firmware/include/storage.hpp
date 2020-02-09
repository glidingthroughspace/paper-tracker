#pragma once

#include <cstdint>

/**
 * This class is a wrapper around the Arduino EEPROM library.
 * It adds an additional function that checks if a value exists in the EEPROM.
 * Since values that were never written contain the value 0xFF, this library assumes uint16s that
 * are 0xFFFF to be non-existant.
 */
class Storage {
  public:
    // Note that a value takes up 2 bytes, so every value's index must increase by 2;
    enum class Value {
      TRACKER_ID = 0,
    };
  public:
    static uint16_t get(Value);
    static bool exists(Value);
    static bool set(Value, uint16_t);
};
