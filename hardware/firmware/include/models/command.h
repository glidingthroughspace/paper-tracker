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


struct Command {
  uint16_t sleepTimeSec;
  CommandType type;


  static bool isValidType(uint64_t type) {
    return (type <= 2);
  }

  bool fromCBOR(uint8_t* bytes, size_t byteslen) {
    cbor::BytesStream bs{bytes, byteslen};
    cbor::Reader cbor{bs};

    // First check if things are well-formed
    if (!cbor.isWellFormed()) {
      logln("Malformed CBOR data while parsing Command");
      return false;
    }

    // Assume the data starts with the "Self-describe CBOR" tag and continues
    // with an array consisting of one boolean item and one text item.
    if (!expectValue(cbor, cbor::DataType::kTag, cbor::kSelfDescribeTag))
      return false;
    uint64_t sleepTimeSecLarge;
    if (!expectUnsignedInt(cbor, &sleepTimeSecLarge))
      return false;
    if (sleepTimeSecLarge > (2^16)) {
      logln("Sleep time in seconds was more than 16 bit integer");
      return false;
    }
    sleepTimeSec = (uint16_t) sleepTimeSecLarge;
    uint64_t commandType;
    if (!expectUnsignedInt(cbor, &commandType))
      return false;
    if (!Command::isValidType(commandType))
      return false;
    type = (CommandType) commandType;

    return true;
  }
};

