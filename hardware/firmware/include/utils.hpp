#pragma once

#include <cstdint>

namespace utils {
namespace time {
  using seconds = unsigned int;
  using milliseconds = unsigned long;
  using timestamp = milliseconds; // On Arduino, the current time can be measured by the milliseconds after starting the MCU
  timestamp current();
  milliseconds to_millis(seconds);
  bool current_time_is_after(timestamp);
  void wait_for_seconds(seconds value);
}
}
