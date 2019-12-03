#pragma once

#include <functional>
#include <CBOR.h>
#include <CBOR_parsing.h>
#include <CBOR_streams.h>

#define MAX_KEY_LENGHT 64

namespace cbor = ::qindesign::cbor;

class CBORDocument {
  public:
    static CBORDocument fromBuffer(uint8_t* buf, size_t buflen);
    bool advance();
    bool isWellformedModel();
    char* findNextKey();
    bool readUnsignedInt(uint64_t& target);
    bool readUnsignedInt(uint32_t& target);
    bool readUnsignedInt(uint16_t& target);
    bool readUnsignedInt(uint8_t& target);
  private:
    bool isAtEOF() const;
    CBORDocument(cbor::BytesStream, cbor::Reader);
    cbor::BytesStream bs;
    cbor::Reader reader;
    cbor::DataType nextDataType;
    char currentKey[MAX_KEY_LENGHT];
};
