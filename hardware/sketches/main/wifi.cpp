#include "wifi.h"

#include <ESP8266WiFi.h>
#include <WiFiClient.h>
#include <WiFiClientSecure.h>

#include "log.h"


Wifi::Wifi() : connected(false) {

};

Wifi::~Wifi() {
   if (connected) {
     WiFi.disconnect();
   }
};

std::vector<ScanResult> Wifi::getAvailableNetworks() {
  std::vector<ScanResult> networks;
  int numberOfNetworks = WiFi.scanNetworks();
  for (int i = 0; i < numberOfNetworks; i++) {
    ScanResult net;
    net.SSID = WiFi.SSID(i);
    net.MAC = WiFi.BSSIDstr(i);
    net.signalStrength = WiFi.RSSI(i);
    networks.push_back(net);
  }

  // the printNetworks function is relatively expensive, so only do this in debug mode
  #ifndef NDEBUG
  printNetworks(networks);
  #endif

  return networks;
}

String padRight(String input, int wantedLength) {
  while (input.length() < wantedLength) {
    input.concat(' ');
  }
  return input;
}

void Wifi::printNetworks(const std::vector<ScanResult>& networks) {
  Log::println("The following networks were found:");
  // Find the SSID with the maximum length
  int maxLength = 0;
  for (auto net : networks) {
    if (maxLength < net.SSID.length()) {
      maxLength = net.SSID.length();
    }
  }
  Log::println(padRight("SSID", maxLength) + '\t' + "MAC             " + '\t' + "RSSID");
  for (auto net: networks) {
    Log::print(padRight(net.SSID, maxLength));
    Log::print('\t');
    Log::print(net.MAC);
    Log::print('\t');
    Log::print(net.signalStrength);
    Log::println();
  }
}

int Wifi::connect(const char* SSID, const char* password) {
  Log::print("Connecting to WiFi network with SSID ");
  Log::print(SSID);
  Log::println("...");
  WiFi.begin(SSID, password);
  while (WiFi.status() != WL_CONNECTED) {
    Log::print('.');
  }
  

}