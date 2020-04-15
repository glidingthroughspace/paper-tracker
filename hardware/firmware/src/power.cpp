#include <power.hpp>

#include <log.hpp>

TinyPICO Power::tinypico = TinyPICO();

// The battery voltage to consider as 0% charged. Anything lower than that will return negative (or
// rolled over) battery percentages. This is taken from the battery's specification.
constexpr float min_battery_voltage = 3.0;
// Maximum voltage considered as 100% battery. The PMIC typically stop charging at around 4.19. The
// percentage will be clipped to 100 at that point.
constexpr float max_battery_voltage = 4.18;

void Power::print_wakeup_reason() {
  #ifndef NDEBUG
  auto wakeup_reason = esp_sleep_get_wakeup_cause();
  switch(wakeup_reason) {
    case ESP_SLEEP_WAKEUP_TIMER : logln("[Power] Wakeup by timer"); break;
    case ESP_SLEEP_WAKEUP_ULP : logln("[Power] Wakeup by ULP"); break;
    default: break;
  }
  #endif
}

void Power::enable_powersavings() {
  Power::tinypico.DotStar_SetPower(false);
}

void Power::halt_for(const utils::time::seconds secs) {
  if (secs == utils::time::seconds(0)) {
    logln("[Power] Not sleeping, since sleep time is 0s");
    return;
  }
  if (secs > Power::MIN_SECONDS_FOR_DEEP_SLEEP) {
    deep_sleep_for(secs);
  } else {
    logf("[Power] Waiting for %ds\n", secs);
    utils::time::wait_for(secs);
  }
}

void Power::deep_sleep_for(const utils::time::seconds secs) {
  logf("[Power] Going to sleep for %ds\n", secs);
  esp_deep_sleep(utils::time::to_micros(secs));
}

uint64_t Power::seconds_to_microseconds(const uint64_t seconds) {
  return seconds * 1000 * 1000;
}

uint8_t Power::get_battery_percentage() {
  auto voltage = tinypico.GetBatteryVoltage();
  uint8_t percentage = (uint8_t)((voltage - min_battery_voltage) / (max_battery_voltage - min_battery_voltage) * 100);
  // Clip to 100%
  percentage = percentage < 100 ? percentage : 100;
  logf("[Power] The battery is %d%% charged (%fV)\n", percentage, voltage);
  return percentage;
}

bool Power::is_charging() {
  auto charging = tinypico.IsChargingBattery();
  logf("[Power] The battery is %s\n", charging ? "charging" : "discharging");
  return charging;
}
