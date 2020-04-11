#include <apiClient.hpp>

#include <string.h>

#include <log.hpp>
#include <power.hpp>
#include <types.hpp>
#include <utils.hpp>

#include <models/newIDResponse.hpp>

std::map<uint16_t, ApiClient::Callback> ApiClient::callbacks;

ApiClient::ApiClient(WiFiUDP* udp, IPAddress serverIP)
  : coap(udp), serverIP(serverIP) {

}

bool ApiClient::start() {
  coap.set_callback(coap_response_callback);
  return coap.start(5688);
}

bool ApiClient::loop() {
  update_request_timeouts();
  return coap.loop();
}

void ApiClient::requestTrackerID(std::function<void(int16_t)> callback) {
  logln("[Api] Requesting new tracker ID from server");
  auto message_id = coap.post(serverIP, "tracker/new");
  storeCallback(message_id, [callback] (coap::Packet& packet, bool timed_out) {
    if (timed_out) {
      logln("[Api] Requesting a new tracker ID timed out. Cannot continue operation, going to sleep");
      Power::deep_sleep_for(utils::time::seconds(10));
      return;
    }
    if (ApiClient::isErrorResponse(packet)) {
      logln("[Api] Requesting a new tracker ID failed. Cannot continue operation, sleeping 10 seconds");
      Power::deep_sleep_for(utils::time::seconds(10));
      return;
    }
    NewIDResponse nid;
    if (!nid.fromCBOR(packet.payload)) {
      logln("[Api] Could not deserialize new tracker ID. Cannot continue operation, sleeping 10 seconds");
      Power::deep_sleep_for(utils::time::seconds(10));
      return;
    }
    callback(nid.getID());
  }, 20); // This is a one-time operation, so we can cope with a longer timeout to make it more likely that the action succeeds
}

void ApiClient::requestNextCommand(uint16_t trackerID, std::function<void(Command&)> callback) {
  logln("[Api] Requesting next action from server");
  // FIXME: It is (in theory) possible for the server to answer so quickly that the response
  // callback is not registered yet. This is highly unlikely though.
  char param[15];
  auto qp = getTrackerIDQueryParam(param, 15, trackerID);
  uint16_t messageID = coap.get(serverIP, "tracker/poll", qp);
  storeCallback(messageID, [callback] (coap::Packet& packet, bool timed_out) {
    if (timed_out) {
      logln("[Api] Requesting the next action timed out, going to sleep");
      Power::deep_sleep_for(utils::time::seconds(10));
      return;
    }
    if (ApiClient::isErrorResponse(packet)) {
      logln("[Api] Requesting the next action failed, going to sleep for 10 seconds");
      Power::deep_sleep_for(utils::time::seconds(10));
      return;
    }
    Command cmd;
    if (!cmd.fromCBOR(packet.payload)) {
      logln("[Api] Could not deserialize next command, going to sleep for 10 seconds");
      Power::deep_sleep_for(utils::time::seconds(10));
      return;
    }

    callback(cmd);
  });
}


void ApiClient::writeTrackingData(uint16_t trackerID, std::vector<uint8_t> scanResults, std::function<void(void)> callback) {
  logf("[Api] Posting scan results to server, sending %d scan result bytes\n", scanResults.size());
  char param[15];
  auto qp = getTrackerIDQueryParam(param, 15, trackerID);
  auto msgID = coap.post(serverIP, "tracker/tracking", qp, scanResults, ContentType::APPLICATION_CBOR);
  storeCallback(msgID, [callback] (coap::Packet& packet, bool timed_out) {
    if (timed_out) {
      logln("[Api] Request to send tracking data timed out, continuing");
      callback();
    }
    if (ApiClient::isErrorResponse(packet)) {
      logf("[Api] Failed to send tracking data: %d\n", packet.code);
      return;
    }
    logln("[Api] Got response to tracking data request");

    callback();
  });
}

void ApiClient::writeInfoResponse(uint16_t trackerID, std::vector<uint8_t> infoResponse) {
  logln("[Api] Posting battery information to server");
  char param[15];
  auto qp = getTrackerIDQueryParam(param, 15, trackerID);
  auto msgID = coap.post(serverIP, "tracker/status", qp, infoResponse, ContentType::APPLICATION_CBOR);
  storeCallback(msgID, [] (coap::Packet& packet, bool timed_out) {
    if (timed_out) {
      logln("[Api] Request to send battery info timed out, continuing");
    }
    if (ApiClient::isErrorResponse(packet)) {
      logln("[Api] Failed to info response data");
      logln(packet.code);
      return;
    }
    logln("[Api] Sent battery info");
  });
}



void ApiClient::coap_response_callback(coap::Packet& packet, IPAddress ip, int port) {
  logf("[Api] Got a CoAP response for request with ID %d, payload is: ", packet.messageid);

  char p[packet.payload.size() + 1];
  memcpy(p, packet.payload.data(), packet.payload.size());
  p[packet.payload.size()] = '\0';
  logln(p);

  auto it = callbacks.find(packet.messageid);
  if (it == callbacks.end()) {
    logf("[Api] No callback registered for message with ID %d\n", packet.messageid);
    return;
  }
  it->second.response_received = true;
  it->second.function(packet, false);
  logf("[Api] Unregistering callback with ID %d\n", packet.messageid);
  callbacks.erase(it);
}

void ApiClient::storeCallback(uint16_t messageID, coap_callback callback, utils::time::seconds timeout) {
  if (messageID == 0) {
    logln("[Api] Sending the message failed");
    return;
  }
  logf("[Api] Message has ID %d\n", messageID);
  callbacks[messageID] = Callback{
    timeout,
    utils::time::now(),
    false,
    callback,
  };
}

bool ApiClient::isErrorResponse(const coap::Packet& response) {
  return response.code > RESPONSE_CODE(2, 31);
}

std::vector<char*> ApiClient::getTrackerIDQueryParam(char* buffer, size_t bufferlen, uint16_t trackerID) {
  std::vector<char*> queryParams;
  snprintf(buffer, bufferlen, "trackerid=%d", trackerID);
  queryParams.push_back(buffer);
  return queryParams;
}

void ApiClient::update_request_timeouts() {
  for (auto& it : callbacks) {
    logf("[Api] There is a callback stored for message ID %d, which times out in %ds\n", it.first,
        utils::time::now().seconds_to(it.second.request_started_at + it.second.timeout));
  }
  for (auto& it : callbacks) {
    auto& cb = it.second;
    if (cb.response_received == true) {
      logf("[Api] Already received a response for request with ID %d, removing from callback list\n", it.first);
      callbacks.erase(it.first);
    } else if (utils::time::now().is_after(cb.request_started_at + cb.timeout)) {
      logf("[Api] Request with ID %d timed out\n", it.first);
      coap::Packet empty_packet;
      cb.function(empty_packet, true);
      callbacks.erase(it.first);
    }
  }
}
