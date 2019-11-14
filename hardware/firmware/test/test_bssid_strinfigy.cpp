#include <Arduino.h>
#include <unity.h>
#include <scanResult.h>

// void setUp(void) {
// // set stuff up here
// }

// void tearDown(void) {
// // clean stuff up here
// }

void test_byte_to_hex() {
  char outputBuffer1[2];
  const byte input1(255);
  byte_to_hex(input1, outputBuffer1);
  TEST_ASSERT_EQUAL(outputBuffer1[0], 'F');
  TEST_ASSERT_EQUAL(outputBuffer1[1], 'F');

  char outputBuffer2[2];
  const byte input2(15);
  byte_to_hex(input2, outputBuffer2);
  TEST_ASSERT_EQUAL(outputBuffer2[0], '0');
  TEST_ASSERT_EQUAL(outputBuffer2[1], 'F');

  char outputBuffer3[2];
  const byte input3(16);
  byte_to_hex(input3, outputBuffer3);
  TEST_ASSERT_EQUAL(outputBuffer3[0], '1');
  TEST_ASSERT_EQUAL(outputBuffer3[1], '0');

  char outputBuffer4[2];
  const byte input4(17);
  byte_to_hex(input4, outputBuffer4);
  TEST_ASSERT_EQUAL(outputBuffer4[0], '1');
  TEST_ASSERT_EQUAL(outputBuffer4[1], '1');
}

void test_bssid_to_string() {
  uint8_t BSSID[BSSID_LENGTH]{170, 187, 204, 221, 238, 255};
  ScanResult scanResult;
  memcpy(scanResult.BSSID, BSSID, BSSID_LENGTH);
  char result_bssid[BSSID_STRING_LENGTH];
  scanResult.bssid_to_string(result_bssid);
  TEST_ASSERT_EQUAL_STRING("AA:BB:CC:DD:EE:FF", result_bssid);
}

void setup() {
    // NOTE!!! Wait for >2 secs if board doesn't support software reset via Serial.DTR/RTS
    delay(2000);

    UNITY_BEGIN();    // IMPORTANT LINE!
    RUN_TEST(test_byte_to_hex);
    RUN_TEST(test_bssid_to_string);

    UNITY_END();
}

void loop() { }
