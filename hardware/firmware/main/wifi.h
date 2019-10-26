#pragma once

#include "Arduino.h"
#include "log.h"

#define SCAN_RESULT_BATCH_SIZE 5

struct WifiNetwork {
  int32_t RSSI;
  String SSID;
  String BSSID;
  void print() {
    Log::print("SSID: ");
    Log::print(SSID);
    Log::print(", BSSID: ");
    Log::print(BSSID);
    Log::print(", RSSI: ");
    Log::println(RSSI);
  };
};

class Wifi {
  public:
    Wifi();
    ~Wifi();
    int getVisibleNetworks();
    int getVisibleNetworkBatch(WifiNetwork* results, int size, int offset) const;
    int getVisibleNetworkCount() const;
    void connect(const char* SSID, const char* password);
    void connectDot1X(const char* ssid, const char* username, const char* password);
    bool isConnected() const;
  private:
    bool connected;
    int numVisibleNetworks;
};