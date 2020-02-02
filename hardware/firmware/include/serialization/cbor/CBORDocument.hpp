#pragma once

#include <cstddef>
#include <cstdint>
#include <log.hpp>

#include <CBOR.h>
#include <CBOR_streams.h>
namespace cbor = ::qindesign::cbor;

#include "./CBORPrinter.hpp"

class CBORValue;
class CBORSerializable;

/**
 * Represents a to-be-serialized CBOR document.
 * Virtual class.
 * The concrete implementations are StaticCBORDocument and DynamicCBORDocument
 */
class CBORDocument {
  public:
    CBORDocument() : m_printer{m_buffer}, m_cbor{m_printer} {};
    std::vector<uint8_t> serialize();
    void begin_array(size_t size);
    void begin_array(size_t size, const char* key);
    void begin_map(size_t size);
    void begin_map(size_t size, const char* key);
    void begin_text(size_t size);
    void write_bytes(const char* bytes, size_t size);
    void write_bytes(uint8_t* bytes, size_t size);
    void write_int(int64_t value);
    void write_uint(uint64_t value);
    uint8_t* bytes();
    size_t size() const;
  private:
    std::vector<uint8_t> m_buffer;
    CBORPrinter m_printer;
    cbor::Writer m_cbor;
};
