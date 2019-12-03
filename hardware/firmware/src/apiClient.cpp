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

void ApiClient::requestNextCommand(std::function<void(Command&)> callback) {
  logln("Requesting next action from server");
  // FIXME: It is possible for the server to answer so quickly that the response callback is not
  // registered yet
  uint16_t messageID = coap.get(serverIP, 5688, "tracker/poll", "trackerid=1");
  storeCallback(messageID, [&] (CoapPacket& packet) {
    if (ApiClient::isErrorResponse(packet)) {
      logln("Requesting the next action failed");
      return;
    }
    Command cmd;
    if (!cmd.fromCBOR(packet.payload, packet.payloadlen)) {
      logln("Could not deserialize next command");
      return;
    }

    log("Next Command is ");
    log((uint8_t) cmd.getType());
    log(" and sleep time in seconds is ");
    logln(cmd.getSleepTimeInSeconds());
    callback(cmd);
  });
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
  it->second(packet);
}

void ApiClient::storeCallback(uint16_t messageID, coap_callback callback) {
  if (messageID == 0) {
    logln("Sending the message failed");
    return;
  }
  log("Message has ID ");
  logln(messageID);
  callbacks[messageID] = callback;
}

bool ApiClient::isErrorResponse(const CoapPacket& response) {
  return response.code > RESPONSE_CODE(2, 31);
}