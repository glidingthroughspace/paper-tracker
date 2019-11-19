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
    while (!Serial.available()) { ; }
    Serial.println(); // To end the first line which contains garbled text
  #endif
}

static void logln() { PRINTLN(); }
template <typename T>
static void logln(T value) { PRINTLN(value); }

template <typename T>
static void log(T value) { PRINT(value); }
