
#include <Arduino.h>
#include "wifi.h"
#include "log.h"

#undef NDEBUG

void setup()
{
  Log::initSerial(115200);
  Log::println("Starting...");
  Wifi wifi;
  wifi.getAvailableNetworks();
  Log::println("Done");
}

void loop() {}
