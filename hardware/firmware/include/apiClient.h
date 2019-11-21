#pragma once

#include <WiFiUdp.h>
#include <coap-simple.h>

#include <functional>

class ApiClient {
  public: 
    ApiClient(WiFiUDP& udp);
    bool loop();
    bool start();
    void requestNextAction(std::function<void(void)> callback);
  private:
    Coap coap;
    static void coap_response_callback(CoapPacket &packet, IPAddress ip, int port);
};