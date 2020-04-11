#pragma once

#include <TinyPICO.h>
#include <cstdint>

#include <utils.hpp>

class Power {
  public:
    static void print_wakeup_reason();
    static void enable_powersavings();
    static void deep_sleep_for(const utils::time::seconds);
    static uint8_t get_battery_percentage();
    static bool is_charging();
  private:
    static TinyPICO tinypico;
    static uint64_t seconds_to_microseconds(const uint64_t seconds);
};

