
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
  wifi.getAvailableNetworks();
  wifi.connectDot1X(EDUROAM_SSID, EDUROAM_USERNAME, EDUROAM_PASSWORD);

  pinMode(2, OUTPUT);
}

void loop() {
  digitalWrite(2, HIGH);
  delay(250);
  digitalWrite(2, LOW);
  delay(250);
}
