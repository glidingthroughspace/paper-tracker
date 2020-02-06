#pragma once

#include <Arduino.h>

#include <CBOR.h>
#include <CBOR_parsing.h>
#include <CBOR_streams.h>

#define MAX_KEY_LENGTH 64

namespace cbor = ::qindesign::cbor;

class CBORParser {
  public:
    CBORParser(uint8_t* buf, size_t buflen);
    bool advance();
    bool isWellformedModel();
    char* findNextKey();
    bool readUnsignedInt(uint64_t& target);
    bool readUnsignedInt(uint32_t& target);
    bool readUnsignedInt(uint16_t& target);
    bool readUnsignedInt(uint8_t& target);
    bool readInt(int64_t& target);
    bool readInt(int32_t& target);
    bool readCString(char* target, size_t& target_length);
    bool readString(String& target);
  private:
    bool isAtEOF() const;
    cbor::BytesStream bs;
    cbor::Reader reader;
    cbor::DataType nextDataType;
    char currentKey[MAX_KEY_LENGTH];
};