import 'dart:convert';

import 'package:flutter/cupertino.dart';
import 'package:flutter_test/flutter_test.dart';
import 'package:paper_tracker/model/tracker.dart';

void main() {
  var zeroTime = DateTime.fromMicrosecondsSinceEpoch(0).toUtc();
  const trackerOrig = '{ "ID": 2, "Label": "Test Tracker", "LastPoll": "1970-01-01T00:00:00.000Z", "LastSleepTime": "1970-01-01T00:00:00.000Z" }';

  test('Test tracker deserialization', () async {
    var tracker = Tracker.fromJson(json.decode(trackerOrig));

    expect(tracker.id, 2);
    expect(tracker.label, "Test Tracker");
    expect(tracker.lastPoll, zeroTime);
    expect(tracker.lastSleepTime, zeroTime);
  });

  test('Test tracker serialization', () async {
    var tracker = Tracker(id: 2, label: "Test Tracker", lastPoll: zeroTime, lastSleepTime: zeroTime);

    var trackerString = json.encode(tracker);

    expect(json.decode(trackerString), json.decode(trackerOrig));
  });
}
