#include <utils.hpp>

#include <Arduino.h>

namespace utils {
namespace time {
  timestamp current() {
    return millis();
  }

  milliseconds to_millis(seconds value) {
    return value * 1000;
  }

  bool current_time_is_after(timestamp ts) {
    return millis() > ts;
  }
}
}
