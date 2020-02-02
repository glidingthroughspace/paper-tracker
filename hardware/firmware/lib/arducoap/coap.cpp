#include "./coap.hpp"

#include <utility>

#include "Arduino.h"

namespace coap {

bool Client::start(unsigned int port) {
	this->port = port;
	return udp->begin(port);
}

void Client::set_callback(Callback c) {
	resp = c;
}

uint16_t Client::get(IPAddress ip, const char* url, std::vector<const char*> queryParameters) {
	return send(ip, Method::GET, url, queryParameters);
}

uint16_t Client::post(IPAddress ip, const char *url, std::vector<const char*> queryParameters, std::vector<uint8_t> payload, ContentType content_type) {
	return send(ip, Method::POST, url, queryParameters, payload, content_type);
}

uint16_t Client::send(IPAddress ip, Method method, const char* url, std::vector<const char*> queryParameters, std::vector<uint8_t> payload, ContentType content_type) {

	// Packet intitialization
	Packet packet;
	packet.type = (uint8_t)PacketType::CON;
	packet.code = (uint8_t)method;
	packet.token = nullptr; // TODO: Read up on what the tokens are in CoAP
	packet.tokenlen = 0;
	packet.payload = payload;
	packet.optionnum = 0;
	packet.messageid = rand();

	// Set URI_HOST and URI_PATH
	String ipaddress = ip.toString();
	packet.addOption(OptionNumber::URI_HOST, (uint8_t)ipaddress.length(), (uint8_t*)ipaddress.c_str());
	// Parse the url
	unsigned int idx = 0;
	for (unsigned int i = 0; i < strlen(url); i++) {
		if (url[i] == '/') {
			packet.addOption(OptionNumber::URI_PATH, i-idx, (uint8_t*)(url + idx));
			idx = i + 1;
		}
	}
	if (idx <= strlen(url)) {
		packet.addOption(OptionNumber::URI_PATH, strlen(url)-idx, (uint8_t*)(url + idx));
	}

	// Set Content-Format option
	uint8_t optionBuffer[2] {0};
	if (content_type != ContentType::NONE) {
		optionBuffer[0] = ((uint16_t)content_type & 0xFF00) >> 8;
		optionBuffer[1] = ((uint16_t)content_type & 0x00FF) ;
		packet.addOption(OptionNumber::CONTENT_FORMAT, 2, optionBuffer);
	}

	// Set the URI-Query option
	if (queryParameters.size() > 0) {
		for (auto it = queryParameters.begin(); it < queryParameters.end(); ++it) {
			packet.addOption(OptionNumber::URI_QUERY, strlen(*it), (uint8_t*)*it);
		}
	}

	// send packet
	return send_packet(packet, ip);
}

uint16_t Client::send_packet(Packet& packet, IPAddress ip) {

	// Initialize the buffer with 4 bytes, since the CoAP header is 4 bytes.
	// Now, manipulation of those first four bytes is easy.
	std::vector<uint8_t> buffer(0);

	// TODO: make use of push_back here
	// make the base CoAP packet header, see https://tools.ietf.org/html/rfc7252#section-3
	buffer.push_back(0x01 << 6);
	buffer[0] |= (packet.type & 0x03) << 4;
	buffer[0] |= (packet.tokenlen & 0x0F);
	buffer.push_back(packet.code);
	buffer.push_back(packet.messageid >> 8);
	buffer.push_back(packet.messageid & 0xFF);



	// make the token
	if (packet.token != nullptr && packet.tokenlen <= 0x0F) {
		for (auto i = 0; i < packet.tokenlen; i++) {
			buffer.push_back(packet.token[i]);
		}
	}



	// make the option header
	uint16_t running_delta = 0;
	for (auto i = 0; i < packet.options.size(); i++)  {
		uint32_t optdelta;
		uint8_t len, delta;

		// TODO: This code is left over from coap-simple. This is really hard to understand.
		optdelta = packet.options[i].number - running_delta;
		COAP_OPTION_DELTA(optdelta, &delta);
		COAP_OPTION_DELTA((uint32_t)packet.options[i].length, &len);

		buffer.push_back(0xFF & (delta << 4 | len));

		if (delta == 13) {
			buffer.push_back(optdelta - 13);
		} else if (delta == 14) {
			buffer.push_back((optdelta - 269) >> 8);
			buffer.push_back(0xFF & (optdelta - 269));
		}

		if (len == 13) {
			buffer.push_back(packet.options[i].length - 13);
		} else if (len == 14) {
			buffer.push_back(packet.options[i].length >> 8);
			buffer.push_back(0xFF & (packet.options[i].length - 269));
		}

		for (auto j = 0; j < packet.options[i].length; j++) {
			buffer.push_back(packet.options[i].buffer[j]);
		}

		running_delta = packet.options[i].number;
	}



	// make payload
	if (packet.payload.size() > 0) {
		buffer.push_back(0xFF); // Flag to indicate start of payload
		buffer.insert(buffer.end(), packet.payload.begin(), packet.payload.end());
	}

	// send the packet
	if (udp->beginPacket(ip, port) == 0) {
		return 0;
	}
	if (udp->write(buffer.data(), buffer.size()) < buffer.size()) {
		return 0;
	}
	if (udp->endPacket() == 0) {
		return 0;
	}

	return packet.messageid;
}

bool Client::loop() {

	// The maximum size of packets we accept from the server.
	// TODO: Use an std::vector here, so this size is dynamic
	constexpr auto BUF_MAX_SIZE = 512;
	uint8_t buffer[BUF_MAX_SIZE];
	int32_t packetlen = udp->parsePacket();

	while (packetlen > 0) {
		packetlen = udp->read(buffer, packetlen >= BUF_MAX_SIZE ? BUF_MAX_SIZE : packetlen);

		Packet packet;

		// parse coap packet header
		if (packetlen < HEADER_SIZE || (((buffer[0] & 0xC0) >> 6) != 1)) {
				packetlen = udp->parsePacket();
				continue;
		}

		packet.type = (buffer[0] & 0x30) >> 4;
		packet.tokenlen = buffer[0] & 0x0F;
		packet.code = buffer[1];
		packet.messageid = 0xFF00 & (buffer[2] << 8);
		packet.messageid |= 0x00FF & buffer[3];

		if (packet.tokenlen == 0)  packet.token = NULL;
		else if (packet.tokenlen <= 8)  packet.token = buffer + 4;
		else {
			packetlen = udp->parsePacket();
			continue;
		}

		// parse packet options/payload
		if (HEADER_SIZE + packet.tokenlen < packetlen) {
			int optionIndex = 0;
			uint16_t delta = 0;
			uint8_t *end = buffer + packetlen;
			uint8_t *p = buffer + HEADER_SIZE + packet.tokenlen;
			// 10 = maximum number of options
			while (optionIndex < 10 && *p != 0xFF && p < end) {
				if (0 != parseOption(&packet.options[optionIndex], &delta, &p, end-p)) return false;
				optionIndex++;
			}
			packet.optionnum = optionIndex;

			if (p+1 < end && *p == 0xFF) {
				std::vector<uint8_t> payload;
				payload.insert(payload.end(), (p + 1), end);
				packet.payload = std::move(payload);
			} else {
				packet.payload = {};
			}
		}

		if (packet.type == (uint8_t)PacketType::ACK) {
			resp(packet, udp->remoteIP(), udp->remotePort());
		}  else {
			// This unexpected
			return false;
		}
		// next packet
		packetlen = udp->parsePacket();
	}

	return true;
}

int Client::parseOption(Option* option, uint16_t *running_delta, uint8_t** buf, size_t buflen) {
	uint8_t *p = *buf;
	uint8_t headlen = 1;
	uint16_t len, delta;

	if (buflen < headlen) return -1;

	delta = (p[0] & 0xF0) >> 4;
	len = p[0] & 0x0F;

	if (delta == 13) {
			headlen++;
			if (buflen < headlen) return -1;
			delta = p[1] + 13;
			p++;
	} else if (delta == 14) {
			headlen += 2;
			if (buflen < headlen) return -1;
			delta = ((p[1] << 8) | p[2]) + 269;
			p+=2;
	} else if (delta == 15) return -1;

	if (len == 13) {
			headlen++;
			if (buflen < headlen) return -1;
			len = p[1] + 13;
			p++;
	} else if (len == 14) {
			headlen += 2;
			if (buflen < headlen) return -1;
			len = ((p[1] << 8) | p[2]) + 269;
			p+=2;
	} else if (len == 15)
			return -1;

	if ((p + 1 + len) > (*buf + buflen))  return -1;
	option->number = delta + *running_delta;
	option->buffer = p+1;
	option->length = len;
	*buf = p + 1 + len;
	*running_delta += delta;

	return 0;
}

void Packet::addOption(OptionNumber num, uint8_t length, uint8_t* optiondata) {
	options.emplace_back((uint8_t)num, length, optiondata);
}

void Packet::addOption(Option option) {
	options.push_back(option);
}

} // namespace coap

