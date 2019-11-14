#pragma once

#include <Arduino.h>

// Map for converting bytes to hex strings
constexpr char hexmap[] = {'0', '1', '2', '3', '4', '5', '6', '7',
                           '8', '9', 'A', 'B', 'C', 'D', 'E', 'F'};

/**
 * Converts a single byte to its hexadecimal representation
 * 
 * @params input The byte to be converted
 * @params output A buffer of size 2 (or more). The first two bytes will be modified
 */
void byte_to_hex(uint8_t input, char* output) {
  output[1] = hexmap[input & 0x0F];
  output[0] = hexmap[(input & 0xF0) >> 4];
}