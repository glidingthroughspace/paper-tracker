#pragma once

#include <stddef.h>
#include <stdint.h>
#include "./CBORValue.h"

/** 
 * Represents a to-be-serialized CBOR document.
 * Virtual class. 
 * The concrete implementations are StaticCBORDocument and DynamicCBORDocument 
 */
class CBORDocument {
  public:
    virtual size_t size() const = 0;
    virtual uint8_t* serialize() = 0;
    virtual void addValue(const CBORValue&) = 0;
  private:
    CBORDocument();
    cbor::BytesPrint bp;
    cbor::Writer cbor;
};

template <size_t m_size>
class StaticCBORDocument : public CBORDocument {
  public:
    StaticCBORDocument();
    const size_t size() const { return m_size; };
    const uint8_t* serialize();
  private:
    uint8_t buf[m_size];
};
