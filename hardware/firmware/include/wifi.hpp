#pragma once

#include <models/scanResult.hpp>
#include <vector>
#include <WiFi.h>
#include <WiFiUdp.h>

#define LOCAL_UDP_PORT 4210

class WIFI {
  public:
    WIFI();
    ~WIFI();

    /**
     * Connect to a WPA2-secured network
     */
    bool connect(const char* ssid, const char* password);
    /**
     * Connect to a 802.1X network. Note that this uses the security settings of eduroam.
     */
    bool connect(const char* ssid, const char* username, const char* password);

    uint8_t getVisibleNetworkCount() const;

    /**
     * Retrieves all visible networks in batches (configured in the constructor) and passes each
     * batch to the given callback. Depending on the batch size and the count of networks in the
     * area, the callback will be called multiple times.
     */
    std::vector<ScanResult> getAllVisibleNetworks();

    WiFiUDP& getUDP();

  private:
    WiFiUDP udp;
    uint8_t visibleNetworkCount;
    bool connectLoop();
    uint8_t scanVisibleNetworks();
};
