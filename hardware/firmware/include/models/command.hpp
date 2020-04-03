#pragma once

#include <vector>

#include <serialization/cbor/CBORValue.hpp>
#include <serialization/cbor/CBORParser.hpp>

enum class CommandType {
  SEND_TRACKING_INFO = 0,
  SIGNAL_LOCATION    = 1,
  SLEEP              = 2,
  SEND_INFO          = 3,
  INVALID            = 255,
};

class Command {
  private:
    CBORUint16 sleepTimeSec{"sleep_time_sec"};
    CBORUint8 type{"type"};
    bool isValidType(uint8_t type) const { return (type <= 3); }
    bool parseType(CBORParser&);
  public:
    bool fromCBOR(uint8_t* buffer, size_t bufferSize);
    bool fromCBOR(std::vector<uint8_t>);

    uint16_t getSleepTimeInSeconds() const;
    CommandType getType() const;
    const char* getTypeString() const;
    void print() const;
};
