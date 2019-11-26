#include <apiClient.h>

#include <log.h>

std::map<uint16_t, coap_callback> ApiClient::callbacks;

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

void ApiClient::requestNextAction(coap_callback callback) {
  logln("Requesting next action from server");
  uint16_t messageID = coap.get(serverIP, 5688, "tracker/poll", "trackerid=1");
  storeCallback(messageID, callback);
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

  auto it = callbacks.find(packet.messageid);
  if (it == callbacks.end()) {
    log("No callback registered for message with ID ");
    logln(packet.messageid);
    return;
  }


}

void ApiClient::storeCallback(uint16_t messageID, coap_callback callback) {
  if (messageID == 0) {
    logln("Sending the message failed");
    return;
  }
  log("Message has ID");
  logln(messageID);
  callbacks[messageID] = callback;
}

bool ApiClient::isErrorResponse(const CoapPacket& response) {
  return response.code > RESPONSE_CODE(2, 31);
}