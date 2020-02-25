#include <log.hpp>

void initSerial(const unsigned long baudRate) {
  #ifndef NDEBUG
    Serial.begin(baudRate);
    Serial.println(); // To end the first line which contains garbled text
  #endif
}

void logln() { PRINTLN(); }
