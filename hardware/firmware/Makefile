PLATFORM?=d1_mini # This can be either esp8285 or d1_mini
SERIALPORT?=/dev/ttyUSB0 # Usually the default port

build:
	pio run -e $(PLATFORM)

upload: build
	pio run -e $(PLATFORM) -t upload --upload-port $(SERIALPORT)

monitor:
	pio device monitor --eol LF -b 115400 -p $(SERIALPORT)

run:
	$(MAKE) upload
	$(MAKE) monitor