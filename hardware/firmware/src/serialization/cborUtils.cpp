#include <serialization/cborUtils.h>
#include <log.h>

CBORDocument::CBORDocument(uint8_t* buf, size_t buflen) : bs{buf, buflen}, reader{bs} {

}

bool CBORDocument::isWellformedModel() {
  if (!reader.isWellFormed()) {
    return false;
  }
  // The reader.isWellFormed() call advances the buffer pointer. We'll need to reset it to return to
  // the beginning.
  bs.reset();

  auto type = reader.readDataType();
  // All our models should be sent/received in a map
  if (type != cbor::DataType::kMap) {
    return false;
  }
  // Don't reset the buffer here. We'd like to start reading the next available type after this call
  return true;
}

bool CBORDocument::advance() {
  nextDataType = reader.readDataType();
  return !isAtEOF();
}

bool CBORDocument::isAtEOF() const {
  return nextDataType == cbor::DataType::kEOS;
}

char* CBORDocument::findNextKey() {
  if (nextDataType != cbor::DataType::kText) {
    logln("Expected a key, but got none");
    log("Data type is ");
    logln(static_cast<uint8_t>(nextDataType));
    return nullptr;
  }
  auto bytesAvailable = reader.bytesAvailable();
  if (bytesAvailable > MAX_KEY_LENGTH - 1) {
    logln("Key is too large to read");
    return nullptr;
  }
  reader.readBytes((uint8_t*) currentKey, bytesAvailable);
  currentKey[bytesAvailable] = '\0';
  return currentKey;
}

bool CBORDocument::readUnsignedInt(uint64_t& target) {
  if (!advance()) {
    return false;
  }
  if (nextDataType != cbor::DataType::kUnsignedInt) {
    return false;
  }
  target = reader.getUnsignedInt();
  return true;
}

bool CBORDocument::readUnsignedInt(uint32_t& target) {
  uint64_t targetBuffer;
  if (!readUnsignedInt(targetBuffer)) {
    return false;
  }
  if (targetBuffer > (2^32)) {
    return false;
  }
  target = static_cast<uint32_t>(targetBuffer);
  return true;
}

bool CBORDocument::readUnsignedInt(uint16_t& target) {
  uint64_t targetBuffer;
  if (!readUnsignedInt(targetBuffer)) {
    return false;
  }
  if (targetBuffer > (2^16)) {
    return false;
  }
  target = static_cast<uint16_t>(targetBuffer);
  return true;
}

bool CBORDocument::readUnsignedInt(uint8_t& target) {
  uint64_t targetBuffer;
  if (!readUnsignedInt(targetBuffer)) {
    return false;
  }
  if (targetBuffer > (2^8)) {
    return false;
  }
  target = static_cast<uint8_t>(targetBuffer);
  return true;
}

template <typename T>
void CBORValue<T>::writeKeyTo(cbor::Writer& cbor) {
  cbor.beginText(keylen);
  cbor.writeBytes((uint8_t*)(m_key), keylen);
}

template <typename T>
void CBORValue<T>::serializeTo(cbor::Writer& cbor) {
  log("Cannot serialize value with key ");
  log(key);
  logln(" because it is of an unknown type");
}

// Specializations of the serializeTo method follow

template <>
void CBORValue<uint8_t>::serializeTo(cbor::Writer& cbor) {
  writeKeyTo(cbor);
  cbor.writeUnsignedInt(value);
}

template <>
void CBORValue<uint16_t>::serializeTo(cbor::Writer& cbor) {
  writeKeyTo(cbor);
  cbor.writeUnsignedInt(value);
}

template <>
void CBORValue<uint32_t>::serializeTo(cbor::Writer& cbor) {
  writeKeyTo(cbor);
  cbor.writeUnsignedInt(value);
}

template <>
void CBORValue<uint64_t>::serializeTo(cbor::Writer& cbor) {
  writeKeyTo(cbor);
  cbor.writeUnsignedInt(value);
}

template <>
void CBORValue<int>::serializeTo(cbor::Writer& cbor) {
  writeKeyTo(cbor);
  cbor.writeInt(value);
}

template <>
void CBORValue<const char*>::serializeTo(cbor::Writer& cbor) {
  writeKeyTo(cbor);
  auto valuelen = strlen(value);
  cbor.beginText(valuelen);
  cbor.writeBytes((uint8_t*)(value), valuelen);
}

// Generic implementation

template <typename T>
bool CBORValue<T>::deserializeFrom(CBORDocument& cbor) {
  log("Cannot deserialize value with key ");
  log(key);
  logln(" because it is of an unknown type");
  return false;
}

// Specializations

template <>
bool CBORValue<uint64_t>::deserializeFrom(CBORDocument& cbor) {
  if (!cbor.readUnsignedInt(value)) {
    logln("Expected a 64 bit unsigned int when reading sleep time, but got something else");
    return false;
  }
  return true;
}

template <>
bool CBORValue<uint32_t>::deserializeFrom(CBORDocument& cbor) {
  if (!cbor.readUnsignedInt(value)) {
    logln("Expected a 32 bit unsigned int when reading sleep time, but got something else");
    return false;
  }
  return true;
}

template <>
bool CBORValue<uint16_t>::deserializeFrom(CBORDocument& cbor) {
  if (!cbor.readUnsignedInt(value)) {
    logln("Expected a 16 bit unsigned int when reading sleep time, but got something else");
    return false;
  }
  return true;
}

template <>
bool CBORValue<uint8_t>::deserializeFrom(CBORDocument& cbor) {
  if (!cbor.readUnsignedInt(value)) {
    logln("Expected a 8 bit unsigned int when reading sleep time, but got something else");
    return false;
  }
  return true;
}
