#pragma once

#include <serialization/cborUtils.h>

enum class CommandType {
	SEND_TRACKING_INFO = 0,
	SIGNAL_LOCATION    = 1,
	SLEEP              = 2,
  INVALID            = 255,
};

class Command {
  private:
    CBORValue<uint16_t> sleepTimeSec = { "SleepTimeSec", 0 };
    CBORValue<uint8_t> type = { "Command", static_cast<uint8_t>(CommandType::SLEEP) };
    bool isValidType(uint8_t type) const { return (type <= 2); }
    bool parseType(CBORDocument&);
  public:
    bool fromCBOR(uint8_t* buffer, size_t bufferSize);

    uint16_t getSleepTimeInSeconds() const;
    CommandType getType() const;
};

