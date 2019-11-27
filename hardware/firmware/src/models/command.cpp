#include <models/command.h>

#include <log.h>

#include <CBOR_parsing.h>
#include <CBOR_streams.h>

bool Command::fromCBOR(uint8_t* buffer, size_t bufferSize) {
  cbor::BytesStream bs{buffer, bufferSize};
  cbor::Reader cbor{bs};

  if (!cbor.isWellFormed()) {
    logln("Malformed CBOR data while parsing Command");
    return false;
  }
  // isWellFormed() advances the current position in the stream, so we have to reset it
  bs.reset();

  uint64_t mapLength;
  bool mapIsIndefinite;
  if (!cbor::expectMap(cbor, &mapLength, &mapIsIndefinite)) {
    logln("CBOR response is not a map");
    return false;
  }

  auto nextType = cbor.readDataType(); 
  do {
    if (nextType != cbor::DataType::kText) {
      logln("Expected a key, but got none");
      return false;
    }
    auto bytesAvailable = cbor.bytesAvailable();
    char key[bytesAvailable + 1];
    cbor.readBytes((uint8_t*) key, bytesAvailable);
    key[bytesAvailable] = '\0';
    if (strcmp(key, kCOMMAND) == 0) {
      if (!parseType(cbor)) {
        return false;
      }
    } else if (strcmp(key, kSLEEP_TIME) == 0) {
      if (!parseSleepTime(cbor)) {
        return false;
      }
    } else {
      log("Command data has unknown key ");
      logln(key);
      return false;
    }
    nextType = cbor.readDataType();
  } while (nextType != cbor::DataType::kEOS);

  return true;
}

uint16_t Command::getSleepTimeInSeconds() const {
  return sleepTimeSec;
}

CommandType Command::getType() const {
  return type;
}

bool Command::parseType(cbor::Reader& cbor) {
  if (cbor.readDataType() != cbor::DataType::kUnsignedInt) {
    logln("Expected an unsigned int, but got none");
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
  uint64_t sleepTimeSecLarge = cbor.getUnsignedInt();
  if (sleepTimeSecLarge > (2^16)) {
    logln("Sleep time in seconds was more than 16 bit integer");
    return false;
  }
  sleepTimeSec = static_cast<uint16_t>(sleepTimeSecLarge);
  return true;
}