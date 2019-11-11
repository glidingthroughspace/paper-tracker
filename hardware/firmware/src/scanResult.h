#pragma once

#include <Arduino.h>
#include "log.h"

struct ScanResult {
	int32_t RSSI;
  // FIMXE: Change this to a char* to optimize allocations
  String BSSID;
  String SSID;

  void print() {
    log("SSID: ");
    log(SSID);
    log(", BSSID: ");
    log(BSSID);
    log(", RSSI: ");
    logln(RSSI);
  }
};