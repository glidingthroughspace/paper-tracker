#pragma once

/*
 * CoAP client library for embedded devices using the Arduino framework.
 * This is free software, licensed under the MIT license. See the license file for more information.
 *
 * Note that the focus of this library is not raw performance, but rather ease of use.
 * As such it will use dynamic allocation and STL containers.
 */

#include "./types.hpp"

#include "Udp.h"
#include "Arduino.h"
#include <vector>
#include <functional>

namespace coap {

constexpr auto HEADER_SIZE = 10;
constexpr auto OPTION_HEADER_SIZE = 1;
constexpr auto PAYLOAD_MARKER = 0xFF;
constexpr unsigned int DEFAULT_PORT = 5683;

// TODO: Refactor the code so this is not needed anymore
constexpr auto MAX_CALLBACK = 10;

class Option {
  public:
    Option(uint8_t number, uint8_t length, uint8_t* buffer) : number{number}, length{length}, buffer{buffer} {};
    uint8_t number;
    uint8_t length;
    uint8_t* buffer;
};

class Packet {
  public:
    uint8_t type;
    uint8_t code;
    uint8_t* token;
    uint8_t tokenlen;
    std::vector<uint8_t> payload;
    uint16_t messageid;

    uint8_t optionnum;
    std::vector<Option> options;

    void addOption(Option);
    void addOption(OptionNumber, uint8_t length, uint8_t* optiondata);
};

using Callback = std::function<void(Packet&, IPAddress, int)>;

// TODO: This class needs refactoring
// TODO: I don't think we even need this for the client
class URI {
  private:
    String u[MAX_CALLBACK];
    Callback c[MAX_CALLBACK];
  public:
    URI();
    void add(Callback call, String url);
    Callback find(String url);
};

class Client {
  private:
    UDP* udp;
    Callback resp;
    unsigned int port;

    uint16_t send_packet(Packet &packet, IPAddress ip);
    int parseOption(Option *option, uint16_t *running_delta, uint8_t **buf, size_t buflen);
  public:
    Client(UDP* udp): udp{udp} {};
    bool start(unsigned int port = DEFAULT_PORT);
    void set_callback(Callback c);

    uint16_t get(IPAddress ip, const char* url, std::vector<const char*> queryParameters = {});
    uint16_t post(IPAddress ip, const char* url, std::vector<const char*> queryParameters = {}, std::vector<uint8_t> payload = {}, ContentType content_type = ContentType::NONE);
    uint16_t send(IPAddress ip, Method method, const char* url, std::vector<const char*> queryParameters = {}, std::vector<uint8_t> payload = {}, ContentType content_type = ContentType::NONE);

    bool loop();
};


} // namespace coap

