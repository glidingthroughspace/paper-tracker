PLATFORM?=tinypico # This can be either tinypico (the default) or desktop
SERIALPORT?=/dev/ttyUSB0 # Usually the default port

build:
	pio run -e $(PLATFORM)

build-all:
	pio run

clean:
	pio run -t clean

upload: build
	pio run -e $(PLATFORM) -t upload --upload-port $(SERIALPORT)

monitor:
	pio device monitor --eol LF -b 115400 -p $(SERIALPORT)

test:
	$(MAKE) test-desktop

test-desktop:
	pio test -e desktop

run:
	$(MAKE) upload
	$(MAKE) monitor

.PHONY: test
