#pragma once

#include "Arduino.h"

#ifndef NDEBUG
#define PRINTLN(x) Serial.println(x)
#define PRINT(x) Serial.print(x)
#else
#define PRINTLN(x)
#define PRINT(x)
#endif

/**
 * Abstraction around Serial for disabling Serial output when not in debug mode
 */
struct Log {
  static void initSerial(const unsigned long baudRate) {
    #ifndef NDEBUG
      Serial.begin(baudRate);
    #endif
  }

  static void println() { PRINTLN(); }
  static void println(const Printable& value) { PRINTLN(value); }
  static void println(const char* value) { PRINTLN(value); }
  static void println(const char value) { PRINTLN(value); }
  static void println(const StringSumHelper& value) { PRINTLN(value); }
  static void println(const int32_t value) { PRINTLN(value); }
  static void println(const uint32_t value) { PRINTLN(value); }
  static void println(const __FlashStringHelper* value) { PRINTLN(value); }
  static void println(const unsigned long value) { PRINTLN(value); }

  static void print(const Printable& value) { PRINT(value); }
  static void print(const char* value) { PRINT(value); }
  static void print(const char value) { PRINT(value); }
  static void print(const StringSumHelper& value) { PRINT(value); }
  static void print(const int32_t value) { PRINT(value); }
  static void print(const uint32_t value) { PRINT(value); }
  static void print(const __FlashStringHelper* value) { PRINT(value); }
  static void print(const unsigned long value) { PRINT(value); }

};
