#include <power.h>

#include <log.h>

TinyPICO Power::tinypico = TinyPICO();

void Power::print_wakeup_reason() {
  #ifndef NDEBUG
  auto wakeup_reason = esp_sleep_get_wakeup_cause();
  switch(wakeup_reason) {
    case ESP_SLEEP_WAKEUP_TIMER : logln("Wakeup by timer"); break;
    case ESP_SLEEP_WAKEUP_ULP : logln("Wakeup by ULP"); break;
  }
  #endif
}

void Power::enable_powersavings() {
  Power::tinypico.DotStar_SetPower(false);
}

void Power::deep_sleep_for_seconds(const uint64_t seconds) {
  logln("Going to sleep");
  esp_sleep_enable_timer_wakeup(seconds_to_microseconds(seconds));
  esp_deep_sleep_start();
}

uint64_t Power::seconds_to_microseconds(const uint64_t seconds) {
  return seconds * 1000 * 1000;
}