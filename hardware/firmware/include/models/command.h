#pragma once

#include <CBOR.h>

namespace cbor = ::qindesign::cbor;

enum class CommandType {
	SEND_TRACKING_INFO = 0,
	SIGNAL_LOCATION    = 1,
	SLEEP              = 2,
};

static const char* kCOMMAND = "Command";
static const char* kSLEEP_TIME = "SleepTimeSec";

class Command {
  private:
    uint16_t sleepTimeSec;
    CommandType type;
    bool isValidType(uint64_t type) const { return (type <= 2); }
    bool parseType(cbor::Reader&);
    bool parseSleepTime(cbor::Reader&);
  public:
    bool fromCBOR(uint8_t* buffer, size_t bufferSize);

    uint16_t getSleepTimeInSeconds() const;
    CommandType getType() const;
};

