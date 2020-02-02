#include <serialization/cbor/CBORDocument.hpp>

std::vector<uint8_t> CBORDocument::serialize() {
  return m_buffer;
}

void CBORDocument::begin_array(size_t size) {
  m_cbor.beginArray(size);
}

void CBORDocument::begin_array(size_t size, const char* key) {
  auto keylen = strlen(key);
  begin_text(keylen);
  write_bytes(key, keylen);
  m_cbor.beginArray(size);
}

void CBORDocument::begin_map(size_t size) {
  m_cbor.beginMap(size);
}

void CBORDocument::begin_text(size_t size) {
  m_cbor.beginText(size);
}

void CBORDocument::write_bytes(const char* bytes, size_t size) {
  write_bytes((uint8_t*) bytes, size);
}

void CBORDocument::write_bytes(uint8_t* bytes, size_t size) {
  m_cbor.writeBytes(bytes, size);
}

void CBORDocument::write_uint(uint64_t value) {
  m_cbor.writeUnsignedInt(value);
}

void CBORDocument::write_int(int64_t value) {
  m_cbor.writeInt(value);
}

uint8_t* CBORDocument::bytes() {
  return m_buffer.data();
}

size_t CBORDocument::size() const {
  return m_buffer.size();
}
