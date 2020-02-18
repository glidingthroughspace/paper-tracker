#include <serialization/cbor/CBORValue.hpp>

// TODO: Move all other implementations from the header file
// These were needed because of using CBORDocument

void CBORValue::write_key_to(CBORDocument& cbor) {
  cbor.begin_text(m_keylen);
  cbor.write_bytes(m_key, m_keylen);
}

void CBORUint64::serialize_to(CBORDocument& cbor) {
  write_key_to(cbor);
  cbor.write_uint(value);
}

void CBORUint16::serialize_to(CBORDocument& cbor) {
  write_key_to(cbor);
  cbor.write_uint(value);
}

void CBORUint8::serialize_to(CBORDocument& cbor) {
  write_key_to(cbor);
  cbor.write_uint(value);
}

void CBORInt32::serialize_to(CBORDocument& cbor) {
  write_key_to(cbor);
  cbor.write_int(value);
}

void CBORInt16::serialize_to(CBORDocument& cbor) {
  write_key_to(cbor);
  cbor.write_int(value);
}

void CBORString::serialize_to(CBORDocument& cbor) {
  write_key_to(cbor);
  cbor.begin_text(m_value.length());
  cbor.write_bytes(m_value.c_str(), m_value.length());
}
