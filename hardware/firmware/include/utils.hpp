#pragma once

#include <cstdint>

namespace utils {
namespace time {
  using seconds = unsigned int;
  using milliseconds = unsigned long;
  using microseconds = unsigned long long;

  class timestamp {
    public:
      timestamp(const milliseconds v);
      timestamp(const seconds v);
      timestamp();
      friend timestamp operator+(const timestamp& a, const seconds b) { return timestamp(a.value + (b * 1000)); };
      friend timestamp operator+(const timestamp& a, const milliseconds b) { return timestamp(a.value + b); };
      friend const bool operator>(const timestamp& lhs, const timestamp& rhs) { return lhs.value > rhs.value; };
      friend const bool operator<(const timestamp& lhs, const timestamp& rhs) { return lhs.value < rhs.value; };
      const bool is_after(const timestamp& other) const;
      const seconds seconds_to(const timestamp& other) { return seconds((other.value - value) / 1000); };
    private:
      // On Arduino, the current time can be measured by the milliseconds after starting the MCU
      milliseconds value;
  };

  timestamp now();
  milliseconds to_millis(seconds);
  microseconds to_micros(seconds);
  void wait_for(seconds value);
  void wait_for(milliseconds value);
}
}
