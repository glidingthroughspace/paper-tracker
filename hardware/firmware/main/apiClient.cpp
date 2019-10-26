#include "apiClient.h"
#include "log.h"

ApiClient::ApiClient() : wifiNetworkBuf({}) {
}

void ApiClient::getVisbleNetworks(Wifi& wifiClient) {
  wifiClient.getVisibleNetworks();

  int offset = 0;
  int count;

  while (true) {
    count = wifiClient.getVisibleNetworkBatch(wifiNetworkBuf, WIFI_NETWORK_BUFFER_SIZE, offset);
    for (int i = 0; i < count; i++) {
      wifiNetworkBuf[i].print();
    }
    Log::print("Got ");
    Log::print(count);
    Log::println(" networks");
    if (count == 0) {
      break;
    }
    auto doc = getVisibleNetworksBatchAsJSON(count);
    // TODO: Serialize to an HTTP connection
    #ifndef NDEBUG
    serializeJson(doc, Serial);
    #endif
    Log::println();
    offset += count;
  }
}

JsonDocument ApiClient::getVisibleNetworksBatchAsJSON(int networkCount) {
  const size_t capacity = JSON_OBJECT_SIZE(3) // The outer object
                          + JSON_ARRAY_SIZE(networkCount)  // The array of visible networks
                          // FIXME: For some reason, the strings are too large for the objects. As
                          // a workaround, we just allocate three times the object size 
                          + (networkCount * 3) * JSON_OBJECT_SIZE(3);  // Contents of the array of found networks
  DynamicJsonDocument doc(capacity);

  // TODO: Fill those with proper content
  doc["trackerID"] = "abcdef";
  doc["done"] = false;

  JsonArray wifiNetworks = doc.createNestedArray("networks");

  for (int i = 0; i < networkCount; i++) {
    Log::print("Serializing network: ");
    wifiNetworkBuf[i].print();
    JsonObject network = wifiNetworks.createNestedObject();
    if (network.isNull()) {
      Log::println("Object is null");
      Log::print("Free heap size: ");
      Log::println(ESP.getFreeHeap());
    }
    if (!network["SSID"].set(wifiNetworkBuf[i].SSID)) {
      Log::println("Failed to add SSID");
    }
    network["RSSI"] = wifiNetworkBuf[i].RSSI;
    if (!network["BSSID"].set(wifiNetworkBuf[i].BSSID)) {
      Log::println("Failed to add BSSID");
    }
  }

  Log::print("Document size: ");
  Log::print(doc.memoryUsage());
  Log::print(", capacity: ");
  Log::println(doc.capacity());

  return doc;
}