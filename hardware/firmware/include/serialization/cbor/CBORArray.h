#pragma once

#include "./CBORDocument.h"
#include "./CBORValue.h"
#include <log.h>

#include <type_traits>

class CBORArray : public CBORValue {
  protected: 
    CBORArray(char* key) : CBORValue{CBORType::Array, key, strlen(key)} {};
  private:
    CBORArray() = delete;
  public:
    virtual size_t size() const = 0;
};

template <size_t m_size, typename T, typename = std::enable_if<std::is_base_of<CBORSerializable, T>::value>>
class StaticCBORArray : public CBORArray {
  private:
    T arr[m_size];
  public:
    void serializeTo(cbor::Writer& cbor) {
      writeKeyTo(cbor);
      cbor.beginArray(m_size);
      for (auto i = 0; i < m_size; i++) {
        // FIXME: Find out the CBOR size somehow
        StaticCBORDocument<100> document;
        document.addValue(arr[i]);
        auto buf = document.serialize();
        auto written = cbor.write(buf, document.size());
        if (written < document.size()) {
          logln("StaticCBORDocument was too small to serialize CBORSerializable in array!");
          continue;
        }
      }
    };
    bool deserializeFrom(CBORParser& parser) {
      // TODO: This is a no-op for now. We don't currently need to deserialize arrays
      return false;
    };
    StaticCBORArray(char* key) : CBORArray{key} {};
    size_t size() const { return m_size; };
    T& operator[](int index) { return arr[index]; };
};
