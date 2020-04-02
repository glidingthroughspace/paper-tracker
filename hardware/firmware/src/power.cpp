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
    case ESP_SLEEP_WAKEUP_TIMER : logln("Wakeup by timer"); break;
    case ESP_SLEEP_WAKEUP_ULP : logln("Wakeup by ULP"); break;
    default: break;
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

uint8_t Power::get_battery_percentage() {
  auto voltage = tinypico.GetBatteryVoltage();
  log("The battery voltage is: ");
  log(voltage);
  uint8_t percentage = (uint8_t)((voltage - min_battery_voltage) / (max_battery_voltage - min_battery_voltage) * 100);
  // Clip to 100%
  percentage = percentage < 100 ? percentage : 100;
  log(", this equals a percentage of: ");
  logln(percentage);
  return percentage;
}

bool Power::is_charging() {
  auto charging = tinypico.IsChargingBattery();
  log("We are charging: ");
  logln(charging);
  return charging;
}
