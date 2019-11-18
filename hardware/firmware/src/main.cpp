#undef NDEBUG

#include <Arduino.h>

#include <log.h>
#include <scanResult.h>
#include <wifi.h>


#include <credentials.h>

#define SCAN_RESULT_BUFFER_SIZE 5
// FIXME: This number is not correct
#define SCAN_RESULT_MESSAGE_OVERHEAD 100


WIFI wifi;
ScanResult scanResultBuffer[SCAN_RESULT_BUFFER_SIZE];

uint8_t bytes[SCAN_RESULT_BUFFER_SIZE * SCAN_RESULT_SIZE_BYTES + SCAN_RESULT_MESSAGE_OVERHEAD]{0};

void setup() {
  initSerial(115400);
  logln("Starting");

  if (!wifi.connect(WIFI_SSID, WIFI_USERNAME, WIFI_PASSWORD)) {
    // TODO: Indicate that the connection failed. Maybe blink the LED?
  }

}

void loop() {
}