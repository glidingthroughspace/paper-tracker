#pragma once

#include <WiFiUdp.h>
#include <coap.hpp>
#include <IPAddress.h>
#include <models/command.hpp>
#include <utils.hpp>

#include <map>
#include <vector>
#include <functional>

using coap_callback = std::function<void(coap::Packet&, bool timed_out)>;
using seconds = uint8_t;
constexpr utils::time::seconds DEFAULT_REQUEST_TIMEOUT(10);

class ApiClient {
  public:
    ApiClient(WiFiUDP* udp, IPAddress serverIP);
    bool loop();
    bool start();
    void requestNextCommand(uint16_t trackerID, std::function<void(Command&)> callback);
    void requestTrackerID(std::function<void(int16_t)> callback);
    static bool isErrorResponse(const coap::Packet& packet);
    void writeTrackingData(uint16_t trackerID, std::vector<uint8_t> scanResults, std::function<void(void)> callback);
    void writeInfoResponse(uint16_t trackerID, std::vector<uint8_t> infoResponse);

    struct Callback {
      utils::time::seconds timeout;
      utils::time::timestamp request_started_at;
      bool response_received;
      coap_callback function;
    };
  private:
    static std::map<uint16_t, Callback> callbacks;
    void storeCallback(uint16_t messageID, coap_callback, utils::time::seconds timeout = DEFAULT_REQUEST_TIMEOUT);
    coap::Client coap;
    IPAddress serverIP;
    static void coap_response_callback(coap::Packet &packet, IPAddress ip, int port);
    std::vector<char*> getTrackerIDQueryParam(char* buffer, size_t bufferlen, uint16_t trackerID);
    void update_request_timeouts();
};
