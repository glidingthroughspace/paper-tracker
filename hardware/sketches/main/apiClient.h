#pragma once

#include "Arduino.h"
#include "wifi.h"
#include <ArduinoJson.h>

// TODO: Decide how large the buffer should be
#define WIFI_NETWORK_BUFFER_SIZE 5

class ApiClient {
  public:
    ApiClient();
    void getVisbleNetworks(Wifi& wifiClient);
    JsonDocument getVisibleNetworksBatchAsJSON(int networkCount);
  private:
    WifiNetwork wifiNetworkBuf[WIFI_NETWORK_BUFFER_SIZE];
};