#pragma once

#include <WiFiUdp.h>
#include <coap.hpp>
#include <IPAddress.h>
#include <models/command.hpp>

#include <map>
#include <vector>
#include <functional>

using coap_callback = std::function<void(coap::Packet&)>;

class ApiClient {
  public:
    ApiClient(WiFiUDP* udp, IPAddress serverIP);
    bool loop();
    bool start();
    void requestNextCommand(uint16_t trackerID, std::function<void(Command&)> callback);
    void requestTrackerID(std::function<void(uint16_t)> callback);
    static bool isErrorResponse(const coap::Packet& packet);
    void writeTrackingData(uint16_t trackerID, std::vector<uint8_t> scanResults, std::function<void(void)> callback);
  private:
    static std::map<uint16_t, coap_callback> callbacks;
    void storeCallback(uint16_t messageID, coap_callback);
    coap::Client coap;
    IPAddress serverIP;
    static void coap_response_callback(coap::Packet &packet, IPAddress ip, int port);
    std::vector<char*> getTrackerIDQueryParam(uint16_t trackerID);
};
