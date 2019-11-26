#pragma once

#include <log.h>
#include <CBOR.h>
#include <CBOR_parsing.h>
#include <CBOR_streams.h>

namespace cbor = ::qindesign::cbor;

enum class CommandType {
	SEND_TRACKING_INFO = 0,
	SIGNAL_LOCATION    = 1,
	SLEEP              = 2,
};

typedef enum {
  kCOMMAND = "Command",
  kSLEEP_TIME = "SleepTimeSec",
} CommandFields;

struct Command {
  uint16_t sleepTimeSec;
  CommandType type;


  static bool isValidType(uint64_t type) {
    return (type <= 2);
  }

  bool fromCBOR(uint8_t* bytes, size_t byteslen) {
    cbor::BytesStream bs{bytes, byteslen};
    cbor::Reader cbor{bs};

    if (!cbor.isWellFormed()) {
      logln("Malformed CBOR data while parsing Command");
      return false;
    }
    // isWellFormed() advances the current position in the stream
    bs.reset();

    uint64_t mapLenght;
    bool mapIsIndefinite;
    if (!cbor::expectMap(cbor, &mapLength, &mapIsIndefinite)) {
      logln("CBOR response is not a map");
      return;
    }
    auto dataType = cbor.readDataType();
    logln((int) dataType);
    int64_t commandType = cbor.getInt();
    log("Command is ");
    logln((int32_t) commandType);
    if (!Command::isValidType(commandType)) {
      logln("Found invalid command number");
      return false;
    }
    type = static_cast<CommandType>(commandType);

    uint64_t sleepTimeSecLarge = cbor.getUnsignedInt();
    if (sleepTimeSecLarge > (2^16)) {
      logln("Sleep time in seconds was more than 16 bit integer");
      return false;
    }
    sleepTimeSec = static_cast<uint16_t>(sleepTimeSecLarge);
    return true;
  }
};

