#pragma once

/**
 * Abstraction around Serial for disabling Serial output when not in debug mode
 */

#include "Arduino.h"

#ifndef NDEBUG
#define PRINTLN(x) Serial.println(x)
#define PRINT(x) Serial.print(x)
#else
#define PRINTLN(x)
#define PRINT(x)
#endif


static void initSerial(const unsigned long baudRate) {
  #ifndef NDEBUG
    Serial.begin(baudRate);
    Serial.println(); // To end the first line which contains garbled text
  #endif
}

static void logln() { PRINTLN(); }
static void logln(const Printable& value) { PRINTLN(value); }
static void logln(const char* value) { PRINTLN(value); }
static void logln(const char value) { PRINTLN(value); }
static void logln(const StringSumHelper& value) { PRINTLN(value); }
static void logln(const int32_t value) { PRINTLN(value); }
static void logln(const uint32_t value) { PRINTLN(value); }
static void logln(const __FlashStringHelper* value) { PRINTLN(value); }
static void logln(const unsigned long value) { PRINTLN(value); }

static void log(const Printable& value) { PRINT(value); }
static void log(const char* value) { PRINT(value); }
static void log(const char value) { PRINT(value); }
static void log(const StringSumHelper& value) { PRINT(value); }
static void log(const int32_t value) { PRINT(value); }
static void log(const uint32_t value) { PRINT(value); }
static void log(const __FlashStringHelper* value) { PRINT(value); }
static void log(const unsigned long value) { PRINT(value); }
