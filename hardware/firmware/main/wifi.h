#pragma once

#include "Arduino.h"
#include "log.h"

// ESP8266 SDK
extern "C" {
  #include "user_interface.h"
  #include "wpa2_enterprise.h"
}

#define SCAN_RESULT_BATCH_SIZE 5

struct WifiNetwork {
	int32_t RSSI;
	String SSID;
	String BSSID;
	void print() {
		Log::print("SSID: ");
		Log::print(SSID);
		Log::print(", BSSID: ");
		Log::print(BSSID);
		Log::print(", RSSI: ");
		Log::println(RSSI);
	};
};

class Wifi {
	public:
		Wifi();
		~Wifi();
		int getVisibleNetworks();
		int getVisibleNetworkBatch(WifiNetwork* results, int size, int offset) const;
		int getVisibleNetworkCount() const;
		void connect();
		void connectWPA2();
		void connectDot1X();
		bool isConnected() const;
		station_status_t getStatus() const;
	private:
		bool setStaticConfig(const unsigned long ip, const unsigned long gateway, const unsigned long subnet) const;
		bool setStaticConfig(const unsigned long ip, const unsigned long gateway, const unsigned long subnet, const unsigned long dns) const;
		void connectLoop();
		bool setSTAMode(bool enable) const;
		bool setMode(uint8 m) const;
		int numVisibleNetworks;
};