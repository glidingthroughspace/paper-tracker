#pragma once

#include "./CBORParser.h"
#include <log.h>
#define CAP(val, max) max > val ? max : val

class CBORDocument;

enum class CBORType {
  UnsignedInt = 0,
  SignedInt   = 1,
  Bytes       = 2,
  Text        = 3,
  Array       = 4,
  Map         = 5,
  Tag         = 6,
  Float       = 7,
  Double      = 8,
  Boolean     = 9,
  Null        = 10,
  Undefined   = 11,
  Break       = 12,
  SimpleValue = 13,
};

/**
 * Virtual class for objets which are CBOR-serializable.
 */
class CBORSerializable {
  public:
    /**
     * Basic CBOR serialization method. Since a CBORDocument is a map, this should directly write
     * the key(s) and value(s). This should be the preferred way to serialize a CBORValue.
     */
    virtual bool toCBOR(CBORDocument&) = 0;
    CBORSerializable() {};
};

/**
 * This represents a value that is serializable and deserializable from/to CBOR.
 * This is a virtual class, implementations for concrete types are below this.
 */
class CBORValue {
  private:
    CBORValue() = delete;
  protected:
    CBORType m_type;
    const char* m_key;
    size_t m_keylen;
    void writeKeyTo(cbor::Writer& cbor) {
      logln("Writing key");
      cbor.beginText(m_keylen);
      cbor.writeBytes((uint8_t*) m_key, m_keylen);
    };
    CBORValue(CBORType type, const char* key, size_t keylength) : m_type{type}, m_key{key}, m_keylen{keylength} {};
  public:
    const char* key() const { return m_key; };
    CBORType type() const { return m_type; };
    virtual void serializeTo(cbor::Writer&) = 0;
    bool matchesKey(const char* key) const { return strcmp(m_key, key) == 0; };
    virtual bool deserializeFrom(CBORParser&) = 0;
};

struct CBORUint64 : public CBORValue {
  CBORUint64(const char* key) : CBORValue(CBORType::UnsignedInt, key, strlen(key)) {};
  uint64_t value;
  void serializeTo(cbor::Writer& cbor) {
    writeKeyTo(cbor);
    cbor.writeUnsignedInt((uint64_t) value);
  };
  bool deserializeFrom(CBORParser& parser) { return parser.readUnsignedInt(value); };
};

struct CBORUint16 : public CBORValue {
  CBORUint16(const char* key) : CBORValue(CBORType::UnsignedInt, key, strlen(key)) {};
  uint16_t value;
  void serializeTo(cbor::Writer& cbor) {
    writeKeyTo(cbor);
    cbor.writeUnsignedInt((uint64_t) value);
  };
  bool deserializeFrom(CBORParser& parser) { return parser.readUnsignedInt(value); };
};

struct CBORUint8 : public CBORValue {
  CBORUint8(const char* key) : CBORValue(CBORType::UnsignedInt, key, strlen(key)) {};
  uint8_t value;
  void serializeTo(cbor::Writer& cbor) {
    writeKeyTo(cbor);
    cbor.writeUnsignedInt((uint64_t) value);
  };
  bool deserializeFrom(CBORParser& parser) { return parser.readUnsignedInt(value); };
};

struct CBORInt32 : public CBORValue {
  CBORInt32(const char* key) : CBORValue(CBORType::SignedInt, key, strlen(key)) {};
  int32_t value;
  void serializeTo(cbor::Writer& cbor) {
    writeKeyTo(cbor);
    cbor.writeInt((int64_t) value);
  };
  bool deserializeFrom(CBORParser& parser) { return parser.readInt(value); };
};

class CBORString : public CBORValue {
  private:
    String m_value;
  public:
    CBORString(const char* key) : CBORValue(CBORType::Text, key, strlen(key)) {};
    void serializeTo(cbor::Writer& cbor) {
      writeKeyTo(cbor);
      cbor.beginText(m_value.length());
      cbor.writeBytes((uint8_t*)m_value.c_str(), m_value.length());
    };
    bool deserializeFrom(CBORParser& parser) { return parser.readString(m_value); };
    void set(const String& value) {
      m_value = value;
    };
    String get() const { return m_value; };
    const size_t length() const { return m_value.length(); };
};

