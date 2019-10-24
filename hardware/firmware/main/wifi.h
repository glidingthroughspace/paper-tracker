#pragma once

#include "Arduino.h"

#define SCAN_RESULT_BATCH_SIZE 5

struct WifiNetwork {
  int32_t RSSI;
  String SSID;
  String BSSID;
};

class Wifi {
  public:
    Wifi();
    ~Wifi();
    int getVisibleNetworks();
    int getVisibleNetworkBatch(WifiNetwork* results, int size, int offset) const;
    void connect(const char* SSID, const char* password);
    void connectDot1X(const char* ssid, const char* username, const char* password);
    bool isConnected() const;
  private:
    bool connected;
    int numVisibleNetworks;
};