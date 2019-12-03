#include <serialization/cborUtils.h>
#include <log.h>

CBORDocument CBORDocument::fromBuffer(uint8_t* buf, size_t buflen) {
  cbor::BytesStream bs{buf, buflen};
  cbor::Reader reader{bs};
  return CBORDocument(bs, reader);
}

CBORDocument::CBORDocument(cbor::BytesStream bs, cbor::Reader reader) : bs{bs}, reader{reader} {

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
  if (bytesAvailable > MAX_KEY_LENGHT - 1) {
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