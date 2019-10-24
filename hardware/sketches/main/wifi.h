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
    void connect(const char* SSID, const char* password);
    void connectDot1X(const char* ssid, const char* username, const char* password);
    bool isConnected() const;
  private:
    void printNetworks(const std::vector<ScanResult>& networks);
    bool connected;
};