#pragma once

#include "Arduino.h"

/**
 * Abstraction around Serial for disabling Serial output when not in debug mode
 */
class Log {
  public:
    static void initSerial(const unsigned long baudRate) {
      #ifndef NDEBUG
        Serial.begin(baudRate);
      #endif
    }
    static void println(const Printable& value) {
      #ifndef NDEBUG
        Serial.println(value);
      #endif
    }
    static void println(const char* value) {
      #ifndef NDEBUG
        Serial.println(value);
      #endif
    }
    static void println(const char value) {
      #ifndef NDEBUG
        Serial.println(value);
      #endif
    }
    static void println(const StringSumHelper& value) {
      #ifndef NDEBUG
        Serial.println(value);
      #endif
    }
    static void println() {
      #ifndef NDEBUG
        Serial.println();
      #endif
    }

    static void print(const Printable& value) {
      #ifndef NDEBUG
        Serial.print(value);
      #endif
    }
    static void print(const char* value) {
      #ifndef NDEBUG
        Serial.print(value);
      #endif
    }
    static void print(const char value) {
      #ifndef NDEBUG
        Serial.print(value);
      #endif
    }
    static void print(const StringSumHelper& value) {
      #ifndef NDEBUG
        Serial.print(value);
      #endif
    }

};
