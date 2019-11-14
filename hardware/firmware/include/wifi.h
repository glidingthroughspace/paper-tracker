#pragma once

#include <scanResult.h>
#include <ESP8266WiFi.h>

#define SCAN_RESULT_BUFFER_SIZE 5

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

    uint8_t scanVisibleNetworks();
    uint8_t getVisibleNetworkCount() const;
    /**
     * Fills the given buffer with network scan results, starting at the index given in startAt.
     * 
     * @returns The numer of networks filled into the buffer. If there aren't enough networks
     * available, not all of the buffer might get filled.
     */
    uint8_t getVisibleNetworks(uint8_t startAt, ScanResult* buffer, uint8_t bufferSize);

  private:
    bool connectLoop();
    uint8_t visibleNetworkCount;
};