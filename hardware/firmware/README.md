# Tracker Firmware

This directory contains the firmware for the paper tracker hardware devices.
It is built using [PlatformIO](https://platformio.org), which makes it simple to resolve
dependencies and build it in a CI workflow.

Since the hardware is supposed to be based on an ESP8285 board, this is one of the build targets.
Currently however, we're developing the firmware on the 100% compatible ESP8266 board, namely a
WeMos D1 mini, which is the default build target.

We build the firmware using `make`, since the PlatformIO commands can get quite long. The following
targets are available:

* `build`: Build for the _current_ MCU platform. You can set the platform in the
  `PLATFORM` variable (either `d1_mini` or `esp8285`)
* `build-all`: Build for all defined platforms. This is used in CI
* `upload`: Uploads the code to the current platform. Configure the serial port to use in the
  `SERIALPORT` variable (defaults to `/dev/ttyUSB0`)
* `monitor`: Starts the serial monitor on the given serial port
* `run`: Build and upload the code, start the serial monitor
