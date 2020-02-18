#include <apiClient.hpp>

#include <log.hpp>
#include <power.hpp>
#include <types.hpp>
#include <string.h>

#include <models/newIDResponse.hpp>

std::map<uint16_t, coap_callback> ApiClient::callbacks;

ApiClient::ApiClient(WiFiUDP* udp, IPAddress serverIP)
  : coap(udp), serverIP(serverIP) {

}

bool ApiClient::start() {
  coap.set_callback(coap_response_callback);
  return coap.start(5688);
}

bool ApiClient::loop() {
  return coap.loop();
}

void ApiClient::requestTrackerID(std::function<void(int16_t)> callback) {
  logln("Requesting new tracker ID from server");
  auto message_id = coap.post(serverIP, "tracker/new");
  storeCallback(message_id, [callback] (coap::Packet& packet) {
    if (ApiClient::isErrorResponse(packet)) {
      logln("Requesting a new tracker ID failed. Cannot continue operation, sleeping 10 seconds");
      Power::deep_sleep_for_seconds(10);
      return;
    }
    NewIDResponse nid;
    if (!nid.fromCBOR(packet.payload)) {
      logln("Could not deserialize new tracker ID. Cannot continue operation, sleeping 10 seconds");
      Power::deep_sleep_for_seconds(10);
      return;
    }
    callback(nid.getID());
  });
}

void ApiClient::requestNextCommand(uint16_t trackerID, std::function<void(Command&)> callback) {
  logln("Requesting next action from server");
  // FIXME: It is (in theory) possible for the server to answer so quickly that the response
  // callback is not registered yet. This is highly unlikely though.
  char param[15];
  auto qp = getTrackerIDQueryParam(param, 15, trackerID);
  uint16_t messageID = coap.get(serverIP, "tracker/poll", qp);
  storeCallback(messageID, [callback] (coap::Packet& packet) {
    if (ApiClient::isErrorResponse(packet)) {
      logln("Requesting the next action failed, going to sleep for 10 seconds");
      Power::deep_sleep_for_seconds(10);
      return;
    }
    Command cmd;
    if (!cmd.fromCBOR(packet.payload)) {
      logln("Could not deserialize next command, going to sleep for 10 seconds");
      Power::deep_sleep_for_seconds(10);
      return;
    }

    callback(cmd);
  });
}


void ApiClient::writeTrackingData(uint16_t trackerID, std::vector<uint8_t> scanResults, std::function<void(void)> callback) {
  logln("Posting scan results to server");
  log("Sending ");
  log(scanResults.size());
  logln(" scan result bytes");
  char param[15];
  auto qp = getTrackerIDQueryParam(param, 15, trackerID);
  auto msgID = coap.post(serverIP, "tracker/tracking", qp, scanResults, ContentType::APPLICATION_CBOR);
  storeCallback(msgID, [callback] (coap::Packet& packet) {
    if (ApiClient::isErrorResponse(packet)) {
      logln("Failed to send tracking data");
      logln(packet.code);
      return;
    }
    logln("Got response");

    callback();
  });
}

void ApiClient::coap_response_callback(coap::Packet& packet, IPAddress ip, int port) {
  logln("Got a CoAP response, payload is: ");

  char p[packet.payload.size() + 1];
  memcpy(p, packet.payload.data(), packet.payload.size());
  p[packet.payload.size()] = '\0';
  logln(p);

#if 1
  Serial.println();
  for (auto c : packet.payload) {
    Serial.print(c);
    Serial.print(" ");
  }
  Serial.println();
#endif

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

bool ApiClient::isErrorResponse(const coap::Packet& response) {
  return response.code > RESPONSE_CODE(2, 31);
}

std::vector<char*> ApiClient::getTrackerIDQueryParam(char* buffer, size_t bufferlen, uint16_t trackerID) {
  std::vector<char*> queryParams;
  snprintf(buffer, bufferlen, "trackerid=%d", trackerID);
  queryParams.push_back(buffer);
  logln("Quey parameters:");
  for (auto qp : queryParams) {
    logln(qp);
  }
  return queryParams;
}
