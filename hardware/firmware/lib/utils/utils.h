#pragma once

#include <stdint.h>

#define BSSID_LENGTH 6
// Two characters per byte + colons between bytes + NULL terminator
#define BSSID_STRING_LENGTH BSSID_LENGTH * 2 + BSSID_LENGTH

namespace utils {
  // Map for converting bytes to hex strings
  constexpr char hexmap[] = {'0', '1', '2', '3', '4', '5', '6', '7',
                            '8', '9', 'A', 'B', 'C', 'D', 'E', 'F'};

  /**
   * Converts a single byte to its hexadecimal representation
   * 
   * @params input The byte to be converted
   * @params output A buffer of size 2 (or more). The first two bytes will be modified
   */
  void byte_to_hex(uint8_t input, char* output);

  /**
   * Converts the BSSID to a C string. 
   * @param buf An 18 bytes long buffer to fill with the BSSID. Use the BSSID_STRING_LENGTH constant
   * to allocate your buffer.
   */
  void bssid_to_string(uint8_t* bssid, char* buf);
}

