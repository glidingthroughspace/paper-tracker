#pragma once

#include <models/scanResult.h>
#include <WiFi.h>
#include <WiFiUdp.h>

#define LOCAL_UDP_PORT 4210

class WIFI {
  public:
    WIFI(size_t scan_results_buffer_size);
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

    /**
     * Retrieves all visible networks in batches (configured in the constructor) and passes each
     * batch to the given callback. Depending on the batch size and the count of networks in the
     * area, the callback will be called multiple times.
     */
    void getAllVisibleNetworks(std::function<void(ScanResult* results, size_t results_length)>);

    WiFiUDP& getUDP();

  private:
    WiFiUDP udp;
    bool connectLoop();
    uint8_t visibleNetworkCount;
    size_t m_scan_result_buffer_size;
    ScanResult* m_scan_results_buffer;
};