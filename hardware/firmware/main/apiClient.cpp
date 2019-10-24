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
    Log::println("Networks:");
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
                          + networkCount * JSON_OBJECT_SIZE(3);  // Contents of the array of found networks
  DynamicJsonDocument doc(capacity);

  // TODO: Fill those with proper content
  doc["trackerID"] = "abcdef";
  doc["done"] = false;

  JsonArray wifiNetworks = doc.createNestedArray("wifiNetworks");

  for (int i = 0; i < networkCount; i++) {
    Log::print("Serializing network with SSID ");
    Log::print(wifiNetworkBuf[i].SSID);
    JsonObject network = wifiNetworks.createNestedObject();
    if (network.isNull()) {
      Log::println("Object is null");
    }
    network["SSID"] = wifiNetworkBuf[i].SSID;
    network["RSSI"] = wifiNetworkBuf[i].RSSI;
    network["BSSID"] = wifiNetworkBuf[i].BSSID;
  }

  return doc;
}