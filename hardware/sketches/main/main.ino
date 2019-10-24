// To disable debugging, remove this line (or comment it out). The compiler will set NDEBUG automatically
#undef NDEBUG


#include <Arduino.h>

#include <ArduinoJson.h>

#include "wifi.h"
#include "log.h"
#include "apiClient.h"
#include "credentials.h"

#define MILLISECONDS 1000
#define SECONDS 1000 * MILLISECONDS 
#define MINUTES 60 * SECONDS

#ifndef NDEBUG
#define SLEEP_TIME 15 * SECONDS
#else
#define SLEEP_TIME 15 * MINUTES
#endif

Wifi wifi;
ApiClient apiClient;

void setup()
{
  Log::initSerial(115200);
  Log::println("Starting...");
  if (WIFI_IS_DOT1X == 1) {
    wifi.connectDot1X(WIFI_SSID, WIFI_USERNAME, WIFI_PASSWORD);
  } else {
    wifi.connect(WIFI_SSID, WIFI_PASSWORD);
  }

  apiClient.getVisbleNetworks(wifi);
  delay(100);
  apiClient.getVisbleNetworks(wifi);

  pinMode(2, OUTPUT);
  Log::println("Sleeping for 15 seconds");
  ESP.deepSleep(SLEEP_TIME);
}

void loop() {
  digitalWrite(2, HIGH);
  delay(250);
  digitalWrite(2, LOW);
  delay(250);
}
