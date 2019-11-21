#include <apiClient.h>
#include <IPAddress.h>

#include <log.h>

ApiClient::ApiClient(WiFiUDP& udp) : coap(udp) {

}

bool ApiClient::start() {
  coap.response(coap_response_callback);
  return coap.start();
}

bool ApiClient::loop() {
  return coap.loop();
}

void ApiClient::requestNextAction(std::function<void(void)> callback) {
  logln("Requesting next action from server");
  coap.get(IPAddress(192, 168, 43, 91), 5688, "tracker/poll?trackerid=1");
}

void ApiClient::coap_response_callback(CoapPacket &packet, IPAddress ip, int port) {
  logln("Got a CoAP response, payload is: ");
  
  char p[packet.payloadlen + 1];
  memcpy(p, packet.payload, packet.payloadlen);
  p[packet.payloadlen] = '\0';
  
  logln(p);
}