#pragma once

// This file contains all enums and constants used for CoAP communication

#define RESPONSE_CODE(class, detail) ((class << 5) | (detail))

// if v < 13 {
//  *n = 0xFF & v
// } else {
//  if v <= 0xFF + 13 {
//    *n = 13
//  } else {
//    *n = 14
//  }
// }
#define COAP_OPTION_DELTA(v, n) (v < 13 ? (*n = (0xFF & v)) : (v <= 0xFF + 13 ? (*n = 13) : (*n = 14)))



enum class PacketType {
  CON = 0,
  NONCON = 1,
  ACK = 2,
  RESET = 3
};

enum class Method {
  GET = 1,
  POST = 2,
  PUT = 3,
  DELETE = 4
};

enum class ResponseCode {
  CREATED = RESPONSE_CODE(2, 1),
  DELETED = RESPONSE_CODE(2, 2),
  VALID = RESPONSE_CODE(2, 3),
  CHANGED = RESPONSE_CODE(2, 4),
  CONTENT = RESPONSE_CODE(2, 5),
  BAD_REQUEST = RESPONSE_CODE(4, 0),
  UNAUTHORIZED = RESPONSE_CODE(4, 1),
  BAD_OPTION = RESPONSE_CODE(4, 2),
  FORBIDDEN = RESPONSE_CODE(4, 3),
  NOT_FOUNT = RESPONSE_CODE(4, 4),
  METHOD_NOT_ALLOWD = RESPONSE_CODE(4, 5),
  NOT_ACCEPTABLE = RESPONSE_CODE(4, 6),
  PRECONDITION_FAILED = RESPONSE_CODE(4, 12),
  REQUEST_ENTITY_TOO_LARGE = RESPONSE_CODE(4, 13),
  UNSUPPORTED_CONTENT_FORMAT = RESPONSE_CODE(4, 15),
  INTERNAL_SERVER_ERROR = RESPONSE_CODE(5, 0),
  NOT_IMPLEMENTED = RESPONSE_CODE(5, 1),
  BAD_GATEWAY = RESPONSE_CODE(5, 2),
  SERVICE_UNAVALIABLE = RESPONSE_CODE(5, 3),
  GATEWAY_TIMEOUT = RESPONSE_CODE(5, 4),
  PROXYING_NOT_SUPPORTED = RESPONSE_CODE(5, 5)
};

enum class OptionNumber {
  IF_MATCH = 1,
  URI_HOST = 3,
  E_TAG = 4,
  IF_NONE_MATCH = 5,
  URI_PORT = 7,
  LOCATION_PATH = 8,
  URI_PATH = 11,
  CONTENT_FORMAT = 12,
  MAX_AGE = 14,
  URI_QUERY = 15,
  ACCEPT = 17,
  LOCATION_QUERY = 20,
  PROXY_URI = 35,
  PROXY_SCHEME = 39
};

enum class ContentType {
  NONE = -1,
  TEXT_PLAIN = 0,
  APPLICATION_LINK_FORMAT = 40,
  APPLICATION_XML = 41,
  APPLICATION_OCTET_STREAM = 42,
  APPLICATION_EXI = 47,
  APPLICATION_JSON = 50,
  APPLICATION_CBOR = 60
};


