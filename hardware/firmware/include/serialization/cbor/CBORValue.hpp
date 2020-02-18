#pragma once

#include <log.hpp>

#include "./CBORParser.hpp"
#include "./CBORDocument.hpp"

#define CAP(val, max) max > val ? max : val

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
  private:
    CBORValue() = delete;
  protected:
    CBORType m_type;
    const char* m_key;
    size_t m_keylen;
    void write_key_to(CBORDocument& cbor);
    CBORValue(CBORType type, const char* key, size_t keylength) : m_type{type}, m_key{key}, m_keylen{keylength} {};
  public:
    const char* key() const { return m_key; };
    CBORType type() const { return m_type; };
    virtual void serialize_to(CBORDocument&) = 0;
    bool matchesKey(const char* key) const { return strcmp(m_key, key) == 0; };
    virtual bool deserializeFrom(CBORParser&) = 0;
    virtual size_t serialized_size_bytes() = 0;
};

struct CBORUint64 : public CBORValue {
  CBORUint64(const char* key) : CBORValue(CBORType::UnsignedInt, key, strlen(key)) {};
  uint64_t value;
  void serialize_to(CBORDocument& cbor);
  bool deserializeFrom(CBORParser& parser) { return parser.readUnsignedInt(value); };
  size_t serialized_size_bytes() {
    if (value < 24) { return 1; }
    else if (value < (1 << 8)) { return 2; }
    else if (value < (1UL << 16)) { return 3; }
    else if (value < (1ULL << 32)) { return 5; }
    else { return 9; }
  }
};

struct CBORUint16 : public CBORValue {
  CBORUint16(const char* key) : CBORValue(CBORType::UnsignedInt, key, strlen(key)) {};
  uint16_t value;
  void serialize_to(CBORDocument& cbor);
  bool deserializeFrom(CBORParser& parser) { return parser.readUnsignedInt(value); };
  size_t serialized_size_bytes() {
    if (value < 24) { return 1; }
    else if (value < (1 << 8)) { return 2; }
    else { return 3; }
  }
};

struct CBORUint8 : public CBORValue {
  CBORUint8(const char* key) : CBORValue(CBORType::UnsignedInt, key, strlen(key)) {};
  CBORUint8(const char* key, uint8_t value) : CBORValue(CBORType::UnsignedInt, key, strlen(key)), value{value} {};
  uint8_t value;
  void serialize_to(CBORDocument& cbor) ;
  bool deserializeFrom(CBORParser& parser) { return parser.readUnsignedInt(value); };
  size_t serialized_size_bytes() {
    if (value < 24) { return 1; }
    else { return 2; }
  }
};

struct CBORInt16 : public CBORValue {
  CBORInt16(const char* key) : CBORValue(CBORType::SignedInt, key, strlen(key)) {};
  CBORInt16(const char* key, int16_t value) : CBORValue(CBORType::SignedInt, key, strlen(key)), value{value} {};
  int16_t value;
  void serialize_to(CBORDocument& cbor);
  bool deserializeFrom(CBORParser& parser) { return parser.readInt(value); };
  size_t serialized_size_bytes() {
    // TODO: Calculate actual size, this is only the maximum size
    return 3;
  }
};

struct CBORInt32 : public CBORValue {
  CBORInt32(const char* key) : CBORValue(CBORType::SignedInt, key, strlen(key)) {};
  CBORInt32(const char* key, int32_t value) : CBORValue(CBORType::SignedInt, key, strlen(key)), value{value} {};
  int32_t value;
  void serialize_to(CBORDocument& cbor);
  bool deserializeFrom(CBORParser& parser) { return parser.readInt(value); };
  size_t serialized_size_bytes() {
    // TODO: Calculate actual size, this is only the maximum size
    return 9;
  }
};

class CBORString : public CBORValue {
  private:
    String m_value;
  public:
    CBORString(const char* key) : CBORValue(CBORType::Text, key, strlen(key)) {};
    CBORString(const char* key, const String& value) : CBORValue(CBORType::Text, key, strlen(key)), m_value{value} {};
    void serialize_to(CBORDocument& cbor);
    bool deserializeFrom(CBORParser& parser) { return parser.readString(m_value); };
    size_t serialized_size_bytes() {
      // 9 bytes for the "header".
      // TODO: 9 bytes is the maximum size, not the actual size
      return 9 + m_value.length();
    }
    void set(const String& value) {
      m_value = value;
    };
    String get() const { return m_value; };
    const size_t length() const { return m_value.length(); };
};

