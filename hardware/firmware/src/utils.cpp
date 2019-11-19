#include <utils.h>

namespace utils {
  void byte_to_hex(uint8_t input, char* output) {
    output[1] = hexmap[input & 0x0F];
    output[0] = hexmap[(input & 0xF0) >> 4];
  }
}