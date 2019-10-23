#pragma once

#include "Arduino.h"
#include "vector"

struct ScanResult {
  int32_t signalStrength;
  String SSID;
  String MAC;
};

class Wifi {
  public:
    Wifi();
    ~Wifi();
    std::vector<ScanResult> getAvailableNetworks();
    int connect(const char* SSID, const char* password);
    bool isConnected() const;
  private:
    void printNetworks(const std::vector<ScanResult>& networks);
    bool connected;
};