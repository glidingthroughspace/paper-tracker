#pragma once

#include <TinyPICO.h>
#include <cstdint>

class Power {
  public:
    static void print_wakeup_reason();
    static void enable_powersavings();
    static void deep_sleep_for_seconds(const uint64_t seconds);
  private:
    static TinyPICO tinypico;
    static uint64_t seconds_to_microseconds(const uint64_t seconds);
};

