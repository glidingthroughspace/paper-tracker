#pragma once

/*
 * CoAP client library for embedded devices using the Arduino framework.
 * This is free software, licensed under the MIT license. See the license file for more information.
 *
 * Note that the focus of this library is not raw performance, but rather ease of use.
 * As such it will use dynamic allocation and STL containers.
 */

#include "Udp.h"
#include "Arduino.h"
#include <vector>
#include <functional>

namespace coap {

constexpr HEADER_SIZE = 10;
constexpr OPTION_HEADER_SIZE = 1;
constexpr PAYLOAD_MARKER = 0xFF;
constexpr DEFAULT_PORT = 5683;

// TODO: Refactor the code so this is not needed anymore
constexpr MAX_CALLBACK = 10;

#define RESPONSE_CODE(class, detail) ((class << 5) | (detail))

// if v < 13 {
// 	*n = 0xFF & v
// } else {
// 	if v <= 0xFF + 13 {
// 		*n = 13
// 	} else {
// 		*n = 14
// 	}
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

class Option {
	public:
    uint8_t number;
    uint8_t length;
    uint8_t *buffer;
};

class Packet {
	public:
		uint8_t type;
		uint8_t code;
		uint8_t *token;
		uint8_t tokenlen;
		uint8_t *payload;
		uint8_t payloadlen;
		uint16_t messageid;

		uint8_t optionnum;
		std::vector<CoAPOption> options;

		void addOption(Option);
		void addOption(OptionNumber, uint8_t length, uint8_t* optiondata);
};

typedef Callback = std::function<void(Packet&, IPAddress, int port);

// TODO: This class needs refactoring
// TODO: I don't think we even need this for the client
class URI {
	private:
		String u[MAX_CALLBACK];
		Callback c[MAX_CALLBACK];
	public:
		CoAPURI();
		void add(Callback call, String url);
		Callback find(String url);
};

class Client {
	private:
		UDP udp;
		Callback resp;
		unsigned int port;

		uint16_t send_packet(Packet &packet, IPAddress ip);
		int parseOption(Option *option, uint16_t *running_delta, uint8_t **buf, size_t buflen);
	public:
		Coap(UDP& udp): udp{udp} {};
		bool start(unsigned int port = COAP_DEFAULT_PORT);
		void set_callback(Callback c);

		uint16_t get(IPAddress ip, const char* url, std::vector<const char*> queryParameters = {});
		uint16_t post(IPAddress ip, const char* url, std::vector<const char*> queryParameters = {}, std::vector<uint8_t> payload = {}, ContentType content_type = ContentType::NONE);
		uint16_t send(IPAddress ip, Method method, const char* url, std::vector<const char*> queryParameters = {}, std::vector<uint8_t> payload = {}, ContentType content_type = ContentType::NONE);

		bool loop();
};


} // namespace coap

