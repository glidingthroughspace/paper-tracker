#pragma

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

bool isValidCommandType(uint8_t type) {
  return (type <= 2);
}

struct Command {
  uint16_t sleepTimeSec;
  CommandType type;

  bool fromCBOR(uint8_t bytes, Command* command) {
    // FIXME: sizeof() might not work here
    cbor::BytesStream bs{bytes, sizeof(bytes)};
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
    if (!expectUnsignedInt(cbor, &command->sleepTimeSec))
      return false;
    uint8_t commandType;
    if (!expectUnsignedInt(cbor, commandType))
      return false;
    if (!isValidCommandType(commandType))
      return false;
    &command->type = commandType;

    return true;
    }
  }
}

