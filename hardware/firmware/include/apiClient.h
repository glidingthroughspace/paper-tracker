#pragma once

#include <WiFiUdp.h>

#include <coap-simple.h>

class ApiClient {
  public: 
    ApiClient(WiFiUDP& udp);
    bool loop();
    bool start();
  private:
    Coap coap;
    static void coap_response_callback(CoapPacket &packet, IPAddress ip, int port);
};