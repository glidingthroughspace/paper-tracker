#pragma once

#include "./CBORParser.h"

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
 * This represents a value that is serializable and deserializable from/to CBOR.
 * This is a virtual class, implementations for concrete types are below this.
 */
class CBORValue {
  protected:
    const CBORType m_type;
    const char* m_key;
    const size_t m_keylen;
    void writeKeyTo(cbor::Writer& cbor) {
      cbor.beginText(m_keylen);
      cbor.writeBytes((uint8_t*) m_key, m_keylen);
    };
    CBORValue(CBORType type, const char* key, size_t keylength) : m_type{type}, m_key{key}, m_keylen{keylength} {};
  public:
    const char* key() const { return m_key; };
    const CBORType type() const { return m_type; };
    virtual void serializeTo(cbor::Writer&) = 0;
    bool matchesKey(const char* key) const { return strcmp(m_key, key) == 0; };
    virtual bool deserializeFrom(CBORParser&) = 0;
};

struct CBORUint64 : public CBORValue {
  CBORUint64(const char* key) : CBORValue{CBORType::UnsignedInt, key, strlen(key)} {};
  uint64_t value;
  void serializeTo(cbor::Writer& cbor) {
    writeKeyTo(cbor);
    cbor.writeUnsignedInt((uint64_t) value);
  };
  bool deserializeFrom(CBORParser& parser) { return parser.readUnsignedInt(value); };
};

struct CBORUint16 : public CBORValue {
  CBORUint16(const char* key) : CBORValue{CBORType::UnsignedInt, key, strlen(key)} {};
  uint16_t value;
  void serializeTo(cbor::Writer& cbor) {
    writeKeyTo(cbor);
    cbor.writeUnsignedInt((uint64_t) value);
  };
  bool deserializeFrom(CBORParser& parser) { return parser.readUnsignedInt(value); };
};

struct CBORUint8 : public CBORValue {
  CBORUint8(const char* key) : CBORValue{CBORType::UnsignedInt, key, strlen(key)} {};
  uint8_t value;
  void serializeTo(cbor::Writer& cbor) {
    writeKeyTo(cbor);
    cbor.writeUnsignedInt((uint64_t) value);
  };
  bool deserializeFrom(CBORParser& parser) { return parser.readUnsignedInt(value); };
};

struct CBORInt32 : public CBORValue {
  CBORInt32(const char* key) : CBORValue{CBORType::SignedInt, key, strlen(key)} {};
  int32_t value;
  void serializeTo(cbor::Writer& cbor) {
    writeKeyTo(cbor);
    cbor.writeInt((int64_t) value);
  };
  bool deserializeFrom(CBORParser& parser) { return parser.readInt(value); };
}

#define CBOR_MAX_STRING_LENGTH 1024
class CBORCString : public CBORValue {
  private:
    char* m_value;
    size_t m_length;
  public:
    CBORCString(const char* key) : CBORValue{CBORType::Text, key, strlen(key)} {};
    void serializeTo(cbor::Writer& cbor) {
      writeKeyTo(cbor);
      cbor.beginText(m_length);
      cbor.writeBytes((uint8_t*)m_value, m_length);
    };
    bool deserializeFrom(CBORParser& parser) { return parser.readCString(m_value, m_length); };
    void set(char* value) {
      m_value = strndup(value, CBOR_MAX_STRING_LENGTH);
      m_length = strlen(m_value);
    };
    const char* get() const { return m_value; };
    const size_t length() const { return m_length; };
};