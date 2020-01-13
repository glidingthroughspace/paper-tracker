#pragma once

#include <stddef.h>
#include <stdint.h>
#include <log.h>

#include <CBOR.h>
#include <CBOR_streams.h>
namespace cbor = ::qindesign::cbor;
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
    virtual void addValue(CBORValue&) = 0;
  protected:
    CBORDocument(uint8_t* buf, size_t buflen): bp{buf, buflen}, cbor{bp} {};
    cbor::BytesPrint bp;
    cbor::Writer cbor;
};

/**
 * Implementation of a CBORDocument with a static buffer size
 */
template <size_t m_size>
class StaticCBORDocument : public CBORDocument {
  public:
    StaticCBORDocument(): CBORDocument{buf, m_size} {
      // TODO: There might be a way to optimize this to not use indefinites (e.g. using a second
      // template parameter)
      cbor.beginIndefiniteMap();
    };
    size_t size() const { return m_size; };
    uint8_t* serialize() {
      cbor.endIndefinite();
      return buf;
    };
    void addValue(CBORValue& value) { value.serializeTo(cbor); };
    void addValue(CBORSerializable& value) { 
      value.toCBOR(*this); 
    };
  private:
    uint8_t buf[m_size];
};
