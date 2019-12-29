#include <serialization/cbor/CBORParser.h>
#include <log.h>
#include <string.h>

CBORParser::CBORParser(uint8_t* buf, size_t buflen) : bs{buf, buflen}, reader{bs} {

}

bool CBORParser::isWellformedModel() {
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

bool CBORParser::advance() {
  nextDataType = reader.readDataType();
  return !isAtEOF();
}

bool CBORParser::isAtEOF() const {
  return nextDataType == cbor::DataType::kEOS;
}

char* CBORParser::findNextKey() {
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

bool CBORParser::readUnsignedInt(uint64_t& target) {
  if (!advance()) {
    return false;
  }
  if (nextDataType != cbor::DataType::kUnsignedInt) {
    return false;
  }
  target = reader.getUnsignedInt();
  return true;
}

bool CBORParser::readUnsignedInt(uint32_t& target) {
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

bool CBORParser::readUnsignedInt(uint16_t& target) {
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

bool CBORParser::readUnsignedInt(uint8_t& target) {
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

bool CBORParser::readInt(int64_t& target) {
  if (!advance()) { return false; }
  if (nextDataType != cbor::DataType::kNegativeInt) { return false; }
  target = reader.getInt();
  return true;
}

bool CBORParser::readInt(int32_t& target) {
  int64_t targetBuffer;
  if (!readInt(targetBuffer)) { return false ; }
  if (targetBuffer >= (2^32) || targetBuffer < (2^32) - 1) { return false; }
  target = static_cast<int32_t>(targetBuffer);
  return true;
}


bool CBORParser::readCString(char* target, size_t& target_length) {
  if (!advance()) return false;
  if (nextDataType != cbor::DataType::kText) return false;
  auto length = reader.getLength();
  if (length > target_length) return false;
  reader.readBytes((uint8_t*)target, target_length);
  return true;
}