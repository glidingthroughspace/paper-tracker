#include <models/command.h>

#include <log.h>
#include <serialization/cborUtils.h>

bool Command::fromCBOR(uint8_t* buffer, size_t bufferSize) {
  auto cbor = CBORDocument(buffer, bufferSize);

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

    if (strcmp(key, k_type) == 0) {
      parsedType = parseType(cbor);
    } else if (strcmp(key, k_sleepTimeSec) == 0) {
      parsedSleepTime = parseSleepTime(cbor);
    } else {
      log("Command data has unknown key ");
      logln(key);
    }
  }

  return parsedType && parsedSleepTime;
}

uint16_t Command::getSleepTimeInSeconds() const {
  return sleepTimeSec;
}

CommandType Command::getType() const {
  return type;
}

bool Command::parseType(CBORDocument& cbor) {
  uint8_t commandType;
  if (!cbor.readUnsignedInt(commandType)) {
    logln("Expected an 8 bit usigned int when reading the command type, but got something else");
    return false;
  }
  if (!isValidType(commandType)) {
    logln("Found unknown command number");
    return false;
  }
  type = static_cast<CommandType>(commandType);
  return true;
}

bool Command::parseSleepTime(CBORDocument& cbor) {
  if (!cbor.readUnsignedInt(sleepTimeSec)) {
    logln("Expected a 16 bit unsigned int when reading sleep time, but got something else");
    return false;
  }
  return true;
}