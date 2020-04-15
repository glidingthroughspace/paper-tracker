#pragma once

#include <TinyPICO.h>
#include <cstdint>

#include <utils.hpp>

class Power {
  public:
    static void print_wakeup_reason();
    static void enable_powersavings();
    // halt_for is a higher-level wrapper around deep_sleep_for and utils::time::wait_for, which
    // chooses to deep-sleep or wait depending on the given halt time. If the sleep time is 0, it
    // returns immediately.
    static void halt_for(const utils::time::seconds);
    static void deep_sleep_for(const utils::time::seconds);
    static uint8_t get_battery_percentage();
    static bool is_charging();
  private:
    static constexpr utils::time::seconds MIN_SECONDS_FOR_DEEP_SLEEP = 10;
    static TinyPICO tinypico;
    static uint64_t seconds_to_microseconds(const uint64_t seconds);
};

