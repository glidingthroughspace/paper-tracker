#pragma once

#include <functional>
#include <string.h>
#include <CBOR.h>
#include <CBOR_parsing.h>
#include <CBOR_streams.h>

#define MAX_KEY_LENGTH 64

namespace cbor = ::qindesign::cbor;

class CBORDocument {
  public:
    CBORDocument(uint8_t* buf, size_t buflen);
    bool advance();
    bool isWellformedModel();
    char* findNextKey();
    bool readUnsignedInt(uint64_t& target);
    bool readUnsignedInt(uint32_t& target);
    bool readUnsignedInt(uint16_t& target);
    bool readUnsignedInt(uint8_t& target);
  private:
    bool isAtEOF() const;
    cbor::BytesStream bs;
    cbor::Reader reader;
    cbor::DataType nextDataType;
    char currentKey[MAX_KEY_LENGTH];
};

/**
 * This represents a value that is serializable and deserializable from/to CBOR.
 * The template is specialized for all known types that can directly be serialized (and the other
 * way around).
 * For unknown types, serialization and deserialization will print an error and do nothing else.
 */
template <typename T>
class CBORValue {
  private:
    const char* m_key;
    const size_t keylen;
    void writeKeyTo(cbor::Writer&);
  public:
    CBORValue(const char* key, T value): m_key{key}, keylen{strlen(key)} {};
    const char* key() const { return m_key; };
    T value;
    void serializeTo(cbor::Writer&);
    bool matchesKey(const char* key) const { return strcmp(m_key, key) == 0; };
    // This is a "default" operation. In some cases, you'll want to validate the output. This
    // function only checks if the data type fits.
    bool deserializeFrom(CBORDocument&);
};
