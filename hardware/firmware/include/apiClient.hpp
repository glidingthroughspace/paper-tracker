#pragma once

#include <WiFiUdp.h>
#include <coap/coap-simple.h>
#include <IPAddress.h>
#include <models/command.hpp>
#include <map>

#include <functional>

typedef std::function<void(CoapPacket&)> coap_callback;

class ApiClient {
  public:
    ApiClient(WiFiUDP& udp, IPAddress serverIP);
    bool loop();
    bool start();
    void requestNextCommand(std::function<void(Command&)> callback);
    static bool isErrorResponse(const CoapPacket &packet);
    void writeTrackingData(uint8_t* scanResults, size_t scanResultLen, std::function<void(void)> callback);
  private:
    static std::map<uint16_t, coap_callback> callbacks;
    void storeCallback(uint16_t messageID, coap_callback);
    Coap coap;
    IPAddress serverIP;
    static void coap_response_callback(CoapPacket &packet, IPAddress ip, int port);
};
