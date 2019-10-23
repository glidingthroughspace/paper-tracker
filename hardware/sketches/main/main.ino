
#include <Arduino.h>
#include "wifi.h"
#include "log.h"
#include "credentials.h"

#undef NDEBUG

Wifi wifi;

void setup()
{
  Log::initSerial(115200);
  Log::println("Starting...");
  // wifi.connect(WIFI_SSID, WIFI_PASSWORD);
  auto networks = wifi.getAvailableNetworks();
}

void loop() {}
