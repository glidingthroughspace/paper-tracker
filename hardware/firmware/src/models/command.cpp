#include <models/command.h>

#include <log.h>

#include <CBOR_parsing.h>
#include <CBOR_streams.h>

bool Command::fromCBOR(uint8_t* buffer, size_t bufferSize) {
  cbor::BytesStream bs{buffer, bufferSize};
  cbor::Reader cbor{bs};

  bool parsedType = false;
  bool parsedSleepTime = false; 

  if (!cbor.isWellFormed()) {
    logln("Malformed CBOR data while parsing Command");
    return false;
  }
  // isWellFormed() advances the current position in the stream, so we have to reset it
  bs.reset();

  // We're reading once to skip the map. Next value should be of type text
  cbor::DataType nextType = cbor.readDataType();
  do {
    nextType = cbor.readDataType(); 
    if (nextType != cbor::DataType::kText) {
      logln("Expected a key, but got none");
      log("Data type is ");
      logln(static_cast<uint8_t>(nextType));
      if (parsedType && parsedSleepTime) {
        logln("Ignoring extra data, since all needed values have been parsed");
        return true;
      }
      return false;
    }
    auto bytesAvailable = cbor.bytesAvailable();
    char key[bytesAvailable + 1];
    cbor.readBytes((uint8_t*) key, bytesAvailable);
    key[bytesAvailable] = '\0';
    if (strcmp(key, kCOMMAND) == 0) {
      if (!parseType(cbor))
        return false;
      parsedType = true;
    } else if (strcmp(key, kSLEEP_TIME) == 0) {
      if (!parseSleepTime(cbor))
        return false;
      parsedSleepTime = true;
    } else {
      log("Command data has unknown key ");
      logln(key);
    }
  } while (nextType != cbor::DataType::kEOS);

  return parseType && parsedSleepTime;
}

uint16_t Command::getSleepTimeInSeconds() const {
  return sleepTimeSec;
}

CommandType Command::getType() const {
  return type;
}

bool Command::parseType(cbor::Reader& cbor) {
  if (cbor.readDataType() != cbor::DataType::kUnsignedInt) {
    logln("Expected an unsigned int when reading command type, but got something else");
    return false;
  }
  auto commandType = cbor.getUnsignedInt();
  if (!isValidType(commandType)) {
    logln("Found unknown command number");
    return false;
  }
  type = static_cast<CommandType>(commandType);
  return true;
}

bool Command::parseSleepTime(cbor::Reader& cbor) {
  if (cbor.readDataType() != cbor::DataType::kUnsignedInt) {
    logln("Expected an unsigned int when reading sleep time, but got something else");
    return false;
  }
  uint64_t sleepTimeSecLarge = cbor.getUnsignedInt();
  if (sleepTimeSecLarge > (2^16)) {
    logln("Sleep time in seconds was more than 16 bit integer");
    return false;
  }
  sleepTimeSec = static_cast<uint16_t>(sleepTimeSecLarge);
  return true;
}