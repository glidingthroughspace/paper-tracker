import 'dart:convert';

import 'package:flutter_test/flutter_test.dart';
import 'package:paper_tracker/model/workflow.dart';

void main() {
  const stepOrig =
      '{ "id": 2, "label": "Test Step", "room_id": 3, "options": { "Test 1": [ { "id": 4, "label": "Test 1a", "room_id": 7 } ] } }';

  test('WFStep deserialization', () async {
    var step = WFStep.fromJson(json.decode(stepOrig));

    expect(step.id, 2);
    expect(step.label, "Test Step");
    expect(step.roomID, 3);
    expect(step.options.length, 1);
    expect(step.options.keys.first, "Test 1");
    expect(step.options["Test 1"].length, 1);
    expect(step.options["Test 1"][0].id, 4);
    expect(step.options["Test 1"][0].label, "Test 1a");
    expect(step.options["Test 1"][0].roomID, 7);
  });

  test('WFStep serialization', () async {
    var step = WFStep(
      id: 2,
      label: "Test Step",
      roomID: 3,
      options: {
        "Test 1": [WFStep(id: 4, label: "Test 1a", roomID: 7)],
      },
    );

    var stepString = json.encode(step.toJSON());

    expect(json.decode(stepString), json.decode(stepOrig));
  });
}
