#include <models/newIDResponse.hpp>

#include <log.hpp>
#include <serialization/cbor/CBORParser.hpp>

bool NewIDResponse::fromCBOR(uint8_t* buffer, size_t bufferSize) {
  auto cbor = CBORParser(buffer, bufferSize);
  if (!cbor.isWellformedModel()) {
    logln("[NewIDResponse] Malformed CBOR data while parsing NewIDResponse");
    return false;
  }

  auto deserializedID = false;

  while (cbor.advance()) {
    auto key = cbor.findNextKey();
    if (key == nullptr) {
      logln("[NewIDResponse] Unexpected token in CBOR input, continuing with next token");
      continue;
    }
    if (id.matchesKey(key)) {
      deserializedID = id.deserializeFrom(cbor);
    }
  }
  return deserializedID;
}

bool NewIDResponse::fromCBOR(std::vector<uint8_t> data) {
  return fromCBOR(data.data(), data.size());
}

uint16_t NewIDResponse::getID() const {
  return id.value;
}
