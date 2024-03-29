#pragma once

/**
 * Abstraction around Serial for disabling Serial output when not in debug mode
 */

#ifndef NDEBUG
  #include "Arduino.h"
  #define PRINTLN(x) Serial.println(x)
  #define PRINT(x) Serial.print(x)
  #define PRINTF(fmt, ...) Serial.printf(fmt, ##__VA_ARGS__)
#else
  #define PRINTLN(x)
  #define PRINT(x)
  #define PRINTF(fmt, ...)
#endif

void initSerial(const unsigned long baudRate);

void logln();
template <typename T>
void logln(T value) { PRINTLN(value); }

template <typename T>
void log(T value) { PRINT(value); }

template <typename... Args>
void logf(const char* fmt, Args... args) { PRINTF(fmt, args...); };

