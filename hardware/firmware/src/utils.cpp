#include <utils.hpp>

namespace utils {
  unsigned long to_millis(seconds value) {
    return value * 1000;
  }
}
