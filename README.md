# Paper Tracker

The Paper Tracker is a system for tracking documents inside buildings. It uses WiFi Access Points to
calculate the documents' current location and links it to the current step in its workflow.

It consists of a hardware tracker, a backend server and a mobile application for Android and iOS.

## Getting started

You need to get a Tracker (the hardware device), install the backend application on a server and
download the app for your phone.

### The Server

Download the pre-built server application for
[Linux](https://github.com/glidingthroughspace/paper-tracker/releases/latest/download/paper-tracker_linux_x86_64),
place it on a server and run it. It will listen for HTTP requests on port `8080`, so go ahead and
visit `your-server.domain:8080`, which links all the other downloads you'll need.

Note that the server also listens for CoAP requests on port `5688`, so you might need to allow UDP
requests to that port in your servers firewall.

You could also install [Go](https://go.dev) and build the server yourself by running `make` or `make
run` in the `server/` directory. This way, you can run the server on Windows and MacOS as well.

Additionally there is a Dockerfile available in the `server` folder. The docker also uses `8080` as http and `5688` as CoAP port default.
The config file and database are located in the `/config` folder.

For individually configuring the server, you can use the `config_example.toml` as template and rename it to `config.toml`.
Information about the available parameters are available through the `--help` command of the server executable.

### The Tracker

The tracker is built out of a [TinyPICO](https://tinypico.com), which is an ESP32-based development
board and a LiPo battery.
If you have a Tracker, download the Firmware Flasher for
[Linux](https://github.com/glidingthroughspace/paper-tracker/releases/latest/download/flasher-Linux.zip)
or
[Windows](https://github.com/glidingthroughspace/paper-tracker/releases/latest/download/flasher-Windows.zip),
extract the ZIP file and run the flasher application. It will ask you for WiFi credentials and a
path to the firmware directory, which is bundled in the given ZIP file. For flashing, you'll need to
have [PlatformIO Core](https://platformio.org) installed.
After flashing the tracker, it will automatically connect to the server, if the given WiFi
credentials are correct.

### The App

The app is written in Flutter and thus is available for Android and iOS. However, prebuilt apps only
exist for Android. You can download them from Github Releases, an explanation of which download
you'll need is given on the server's download page.
