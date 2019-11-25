#include <apiClient.h>

#include <log.h>

ApiClient::ApiClient(WiFiUDP& udp, IPAddress serverIP) 
  : coap(udp), serverIP(serverIP) {

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
  coap.get(serverIP, 5688, "tracker/poll", "trackerid=1");
}


void ApiClient::writeTrackingData(uint8_t* scanResults, size_t scanResultLen, std::function<void(void)> callback) {
  logln("Posting scan results to server");
  log("Sending ");
  log(scanResultLen);
  logln(" scan result bytes");
  auto msgID = coap.post(serverIP, 5688, "tracker/tracking", scanResults, scanResultLen, "trackerid=1");
  log("Message ID is ");
  logln(msgID);
}

void ApiClient::coap_response_callback(CoapPacket &packet, IPAddress ip, int port) {
  logln("Got a CoAP response, payload is: ");
  
  char p[packet.payloadlen + 1];
  memcpy(p, packet.payload, packet.payloadlen);
  p[packet.payloadlen] = '\0';
  
  logln(p);
}