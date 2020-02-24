import 'dart:convert';

import 'package:flutter_test/flutter_test.dart';
import 'package:paper_tracker/model/tracker.dart';

void main() {
  var zeroTime = DateTime.fromMicrosecondsSinceEpoch(0).toUtc();
  const trackerOrig =
      '{ "id": 2, "label": "Test Tracker", "last_poll": "1970-01-01T00:00:00.000Z", "last_sleep_time": "1970-01-01T00:00:00.000Z", "status": 1 }';

  test('Test tracker deserialization', () async {
    var tracker = Tracker.fromJson(json.decode(trackerOrig));

    expect(tracker.id, 2);
    expect(tracker.label, "Test Tracker");
    expect(tracker.lastPoll, zeroTime);
    expect(tracker.lastSleepTime, zeroTime);
    expect(tracker.status, TrackerStatus.Idle);
  });

  test('Test tracker serialization', () async {
    var tracker =
        Tracker(id: 2, label: "Test Tracker", lastPoll: zeroTime, lastSleepTime: zeroTime, status: TrackerStatus.Idle);

    var trackerString = json.encode(tracker.toJson());

    expect(json.decode(trackerString), json.decode(trackerOrig));
  });
}
