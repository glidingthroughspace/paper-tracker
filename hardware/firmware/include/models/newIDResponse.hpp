#pragma once

#include <vector>

#include <serialization/cbor/CBORValue.hpp>
#include <serialization/cbor/CBORParser.hpp>

class NewIDResponse {
  private:
    CBORUint16 id{"id"};
  public:
    bool fromCBOR(uint8_t* buffer, size_t bufferSize);
    bool fromCBOR(std::vector<uint8_t>);

    uint16_t getID() const;
};

