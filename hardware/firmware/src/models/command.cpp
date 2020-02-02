#include <models/command.hpp>

#include <log.hpp>
#include <serialization/cbor/CBORParser.hpp>

bool Command::fromCBOR(uint8_t* buffer, size_t bufferSize) {
  auto cbor = CBORParser(buffer, bufferSize);

  bool parsedType = false;
  bool parsedSleepTime = false;

  if (!cbor.isWellformedModel()) {
    logln("Malformed CBOR data while parsing Command");
    return false;
  }

  while(cbor.advance()) {
    auto key = cbor.findNextKey();
    if (key == nullptr) {
      logln("Unexpected token in CBOR input, continuing with next token");
      continue;
    }
    if (type.matchesKey(key)) {
      parsedType = parseType(cbor);
    } else if (sleepTimeSec.matchesKey(key)) {
      parsedSleepTime = sleepTimeSec.deserializeFrom(cbor);
    } else {
      log("Command data has unknown key ");
      logln(key);
    }
  }

  return parsedType && parsedSleepTime;
}

bool Command::fromCBOR(std::vector<uint8_t> data) {
  return fromCBOR(data.data(), data.size());
}

uint16_t Command::getSleepTimeInSeconds() const {
  return sleepTimeSec.value;
}

CommandType Command::getType() const {
  return static_cast<CommandType>(type.value);
}

bool Command::parseType(CBORParser& cbor) {
  if (!type.deserializeFrom(cbor)) {
    return false;
  }
  if (!isValidType(type.value)) {
    logln("Found unknown command number");
    type.value = static_cast<uint8_t>(CommandType::INVALID);
    return false;
  }
  return true;
}

const char* Command::getTypeString() const {
  switch (type.value) {
    case (uint8_t)CommandType::SEND_TRACKING_INFO:
      return "SendTrackingInfo";
    case (uint8_t)CommandType::SIGNAL_LOCATION:
      return "SignalLocation";
    case (uint8_t)CommandType::SLEEP:
      return "Sleep";
    default:
      return "INVALID";
  }
}
