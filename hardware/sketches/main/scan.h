#pragma once

#include "Arduino.h"
#include "vector"

struct ScanResult {
  int signalStrength;
  String SSID;
  String macAddress;
};

class Scanner {
  public:
    static std::vector<ScanResult> getAvailableNetworks();
};