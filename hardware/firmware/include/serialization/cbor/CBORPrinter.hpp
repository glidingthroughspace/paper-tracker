#pragma once
// Implementation of the Print interface for use with dynamic allocation instead of old C-style
// buffers.
// See CBOR_Streams.h of libcbor for the reference implementation.

#include <Print.h>

#include <vector>

class CBORPrinter : public Print {
  public:
    CBORPrinter(std::vector<uint8_t>& buffer) : m_buffer{buffer} {};
    ~CBORPrinter() = default;

    // Writes a byte to the buffer.
    size_t write(uint8_t b) override {
      m_buffer.push_back(b);
      return 1;
    };
  private:
    std::vector<uint8_t>& m_buffer;
};
