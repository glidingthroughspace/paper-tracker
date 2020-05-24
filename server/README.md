# Paper-Tracker Server

For compiling and running with go use the following commands:
```
make run
```
There is also a Dockerfile available.

To configure use the `--help` command or use the example config file and rename it to `config.toml`.

In case you don't have an actual tracker use the following to create a new one:
```
make new-tracker
```
For this to work `coap-client`, `cbor-diag` and `jq` need to be installed.
If you configured a different COAP Port, you need to also adapt the port in the Makefile.

After deployment you can visit the server under its HTTP port to download the app and flasher tool.
