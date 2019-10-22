#include <ESP8266WiFi.h>
#include <WiFiClient.h>
#include <WiFiClientSecure.h>

#include <Arduino.h>
#define LED 2

void setup()
{
  Serial.begin(115200);
  Serial.println();

  int numberOfNetworks = WiFi.scanNetworks();

  Serial.println("The following WiFi networks were found:");
  for (int i = 0; i < numberOfNetworks; i++) {
    Serial.print(WiFi.SSID(i));
    Serial.print("\t");
    Serial.print(WiFi.RSSI(i));
    Serial.println();
  }

}

void loop() {}
