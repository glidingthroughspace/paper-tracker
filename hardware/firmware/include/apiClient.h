#pragma once

#include <WiFiUdp.h>
#include <coap-simple.h>
#include <IPAddress.h>

#include <functional>

class ApiClient {
  public: 
    ApiClient(WiFiUDP& udp, IPAddress serverIP);
    bool loop();
    bool start();
    void requestNextAction(std::function<void(void)> callback);
  private:
    IPAddress serverIP;
    Coap coap;
    uint16_t onNextActionMessageID;
    function<void(void)> onNextActionCallback;
    static void coap_response_callback(CoapPacket &packet, IPAddress ip, int port);
};