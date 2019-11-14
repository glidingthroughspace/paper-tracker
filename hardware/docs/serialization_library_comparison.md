# Serialization library comparison

The following object has been serialized 250 times for the tests:

```cpp
ScanResult res{ -count, "AA:BB:CC:DD:EE", "MyWifi"};
```

which translates to

```json
{
  "RSSI": <-count>,
  "BSSID": "AA:BB:CC:DD:EE",
  "SSID": "MyWifi"
}
```

| Library     | Serialization time (ms) | Result size (bytes) |
| ----------- | ----------------------- | ------------------- |
| ArduinoJson | 460                     | 38                  |
| libCBOR     | 77                      | 28                  |