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

void initSerial(const unsigned long baudRate);

void logln();
template <typename T>
void logln(T value) { PRINTLN(value); }

template <typename T>
void log(T value) { PRINT(value); }

