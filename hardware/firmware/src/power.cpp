#include <power.h>

#include <log.h>

TinyPICO Power::tinypico = TinyPICO();

void Power::print_wakeup_reason() {
  auto wakeup_reason = esp_sleep_get_wakeup_cause();
  switch(wakeup_reason) {
    case ESP_SLEEP_WAKEUP_EXT0 : logln("Wakeup caused by external signal using RTC_IO"); break;
    case ESP_SLEEP_WAKEUP_EXT1 : logln("Wakeup caused by external signal using RTC_CNTL"); break;
    case ESP_SLEEP_WAKEUP_TIMER : logln("Wakeup caused by timer"); break;
    case ESP_SLEEP_WAKEUP_TOUCHPAD : logln("Wakeup caused by touchpad"); break;
    case ESP_SLEEP_WAKEUP_ULP : logln("Wakeup caused by ULP program"); break;
    default : Serial.printf("Wakeup was not caused by deep sleep: %d\n",wakeup_reason); break;
  }
}

void Power::enable_powersavings() {
  Power::tinypico.DotStar_SetPower(false);
}

void Power::deep_sleep_for_seconds(const uint64_t seconds) {
  esp_sleep_enable_timer_wakeup(seconds_to_microseconds(seconds));
  esp_deep_sleep_start();
}

uint64_t Power::seconds_to_microseconds(const uint64_t seconds) {
  return seconds * 1000 * 1000;
}