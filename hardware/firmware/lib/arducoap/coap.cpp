#include "./coap.hpp"

namespace coap {

bool Client::start(unsigned int port = DEFAULT_PORT) {
	this->port = port;
	return udp.begin(port);
}

void Client::set_callback(Callback c) {
	resp = c;
}

uint16_t Client::get(IPAddress ip, const char* url, std::vector<const char*> queryParameters = {}) {
	return send(ip, Method::GET, url, queryParameters);
}

uint16_t Client::post(IPAddress ip, const char *url, std::vector<const char*> queryParameters = {}, std::vector<uint8_t> payload = {}, ContentType content_type = ContentType::NONE) {
	return send(ip, Method::POST, url, queryParameters, payload, content_type);
}

uint16_t Client::send(IPAddress ip, Method method, const char* url, std::vector<const char*> queryParameters = {}, std::vector<uint8_t> payload = {}, ContentType content_type = ContentType::NONE) {
	// Packet intitialization
	Packet packet;
	packet.type = (uint8_t)PacketType::CON;
	packet.code = (uint8_t)method;
	packet.token = nullptr; // TODO: Read up on what the tokens are in CoAP
	packet.tokenlen = 0;
	packet.payload = payload.data();
	packet.payloadlen = payload.size();
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
	std::vector<uint8_t> buffer(4);

	// make the base CoAP packet header, see https://tools.ietf.org/html/rfc7252#section-3
	buffer.data()[0] = 0x01 << 6;
	buffer.data()[0] |= (packet.type & 0x03) << 4;
	buffer.data()[0] |= (packet.tokenlen & 0x0F);
	buffer.data()[1] = packet.code;
	buffer.data()[3] = packet.messageid >> 8;
	buffer.data()[4] = packet.messageid & 0xFF;

	// make the token
	if (packet.token != nullptr && packet.tokenlen <= 0x0F) {
		for (auto i = 0; i < packet.tokenlen; i++) {
			buffer.push_back(packet.token[i]);
		}
	}

	// make the option header
	uint16_t running_delta = 0;
	for (int i = 0; i < packet.options.size(); i++)  {
		uint32_t optdelta;
		uint8_t len, delta;

		optdelta = packet.options[i].number - running_delta;
		COAP_OPTION_DELTA(optdelta, &delta);
		COAP_OPTION_DELTA((uint32_t)packet.options[i].length, &len);

		*p++ = (0xFF & (delta << 4 | len));
		if (delta == 13) {
			*p++ = (optdelta - 13);
			packetSize++;
		} else if (delta == 14) {
			*p++ = ((optdelta - 269) >> 8);
			*p++ = (0xFF & (optdelta - 269));
				packetSize+=2;
		} if (len == 13) {
			*p++ = (packet.options[i].length - 13);
			packetSize++;
		} else if (len == 14) {
			*p++ = (packet.options[i].length >> 8);
			*p++ = (0xFF & (packet.options[i].length - 269));
			packetSize+=2;
		}

		memcpy(p, packet.options[i].buffer, packet.options[i].length);
		p += packet.options[i].length;
		packetSize += packet.options[i].length + 1;
		running_delta = packet.options[i].number;


	}

	// make payload
	if (packet.payloadlen > 0) {
		if ((packetSize + 1 + packet.payloadlen) >= BUF_MAX_SIZE) {
			return 0;
		}
		*p++ = 0xFF;
		memcpy(p, packet.payload, packet.payloadlen);
		packetSize += 1 + packet.payloadlen;
	}

	// send the packet
	udp.beginPacket(ip, port);
	udp.write(buffer.data(), buffer.size());
	udp.endPacket();

	return packet.messageid;
}

} // namespace coap

