#pragma once

#include <serialization/cborUtils.h>

enum class CommandType {
	SEND_TRACKING_INFO = 0,
	SIGNAL_LOCATION    = 1,
	SLEEP              = 2,
};

class Command {
  private:
    uint16_t sleepTimeSec;
    const char* k_sleepTimeSec = "SleepTimeSec";
    CommandType type;
    const char* k_type = "Command";
    bool isValidType(uint8_t type) const { return (type <= 2); }
    bool parseType(CBORDocument&);
    bool parseSleepTime(CBORDocument&);
  public:
    bool fromCBOR(uint8_t* buffer, size_t bufferSize);

    uint16_t getSleepTimeInSeconds() const;
    CommandType getType() const;
};

