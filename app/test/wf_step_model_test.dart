import 'dart:convert';

import 'package:flutter_test/flutter_test.dart';
import 'package:paper_tracker/model/workflow.dart';

void main() {
  const stepOrig =
      '{ "id": 2, "label": "Test Step", "room_ids": [ 3 ], "options": { "Test 1": [ { "id": 4, "label": "Test 1a", "room_ids": [ 7 ] } ] } }';

  test('WFStep deserialization', () async {
    var step = WFStep.fromJson(json.decode(stepOrig));

    expect(step.id, 2);
    expect(step.label, "Test Step");
    expect(step.roomIDs, [3]);
    expect(step.options.length, 1);
    expect(step.options.keys.first, "Test 1");
    expect(step.options["Test 1"].length, 1);
    expect(step.options["Test 1"][0].id, 4);
    expect(step.options["Test 1"][0].label, "Test 1a");
    expect(step.options["Test 1"][0].roomIDs, [7]);
  });

  test('WFStep serialization', () async {
    var step = WFStep(
      id: 2,
      label: "Test Step",
      roomIDs: [3],
      options: {
        "Test 1": [
          WFStep(id: 4, label: "Test 1a", roomIDs: [7])
        ],
      },
    );

    var stepString = json.encode(step.toJSON());

    expect(json.decode(stepString), json.decode(stepOrig));
  });
}
