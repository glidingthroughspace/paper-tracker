#include <utils.h>

namespace utils {
  void byte_to_hex(uint8_t input, char* output) {
    output[1] = hexmap[input & 0x0F];
    output[0] = hexmap[(input & 0xF0) >> 4];
  }

  void bssid_to_string(uint8_t* bssid, char* buf) {
    // A  A  :  B  B  :  C  C  :  D  D  :  E  E  :  F  F
    // 0  1  2  3  4  5  6  7  8  9 10 11 12 13 14 15 16
    buf[2] = buf[5] = buf[8] = buf[11] = buf[14] = ':';
    for (int i = 0; i < BSSID_LENGTH; i++) {
      utils::byte_to_hex(bssid[i], &buf[i * 3]);
    }
    buf[17] = '\0';
  }
}