// To disable debugging, change 'undef' to 'define'
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
#define SLEEP_TIME 10 * SECONDS
#else
#define SLEEP_TIME 1 * MINUTES
#endif

Wifi wifi;
ApiClient apiClient;

void setup()
{
  Log::initSerial(115200);
  Log::println("Starting...");
  const auto startupMills = millis();
  wifi.connect();
  if (!wifi.isConnected()) {
    Log::println("Could not connect to WiFi!");
  }

  // Log::println("Getting networks");
  // apiClient.getVisbleNetworks(wifi);
  // delay(100);
  // apiClient.getVisbleNetworks(wifi);

  const auto shutdownMillis = millis();
  Log::print("Running for ");
  Log::print(shutdownMillis - startupMills);
  Log::println("ms, going to sleep now");
  ESP.deepSleep(SLEEP_TIME);
}

void loop() {
  digitalWrite(2, HIGH);
  delay(250);
  digitalWrite(2, LOW);
  delay(250);
}
