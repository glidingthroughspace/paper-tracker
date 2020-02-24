import 'package:json_annotation/json_annotation.dart';
import 'package:paper_tracker/model/workflow.dart';

part 'createStepRequest.g.dart';

@JsonSerializable()
class CreateStepRequest {
  @JsonKey(name: "prev_step_id")
  int previousStepID;
  @JsonKey(name: "decision_label")
  String decisionLabel;
  @JsonKey(name: "step", toJson: _stepToJSON)
  WFStep step;

  CreateStepRequest({this.previousStepID, this.decisionLabel, this.step});

  factory CreateStepRequest.fromJson(Map<String, dynamic> json) => _$CreateStepRequestFromJson(json);
  Map<String, dynamic> toJson() => _$CreateStepRequestToJson(this);
}

dynamic _stepToJSON(WFStep step) {
  return step.toJSON();
}
