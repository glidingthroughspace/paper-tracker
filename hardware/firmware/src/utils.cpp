#include <utils.hpp>

#include <Arduino.h>

namespace utils {
namespace time {
  timestamp now() {
    return timestamp(millis());
  }

  milliseconds to_millis(seconds value) {
    return value * 1000;
  }

  void wait_for(seconds value) {
    timestamp end_time = now() + value;
    while (!now().is_after(end_time)) {
      yield();
    }
  }

  void wait_for(milliseconds value) {
    timestamp end_time = now() + value;
    while (!now().is_after(end_time)) {
      yield();
    }
  }

  timestamp::timestamp(const milliseconds v) : value{v} {}
  timestamp::timestamp(const seconds v) : value{v * 1000} {}
  const bool timestamp::is_after(const timestamp &other) const {
    return value > other.value;
  }
}
}
