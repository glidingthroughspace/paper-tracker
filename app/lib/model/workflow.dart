import 'package:flutter/material.dart';
import 'package:json_annotation/json_annotation.dart';

part 'workflow.g.dart';

@JsonSerializable()
class Workflow {
  static const IconData = Icons.rotate_left;

  @JsonKey(name: "id")
  int id;
  @JsonKey(name: "label")
  String label;
  @JsonKey(name: "is_template")
  bool isTemplate;
  @JsonKey(name: "steps")
  List<WFStep> steps;

  Workflow({this.id});

  factory Workflow.fromJson(Map<String, dynamic> json) => _$WorkflowFromJson(json);
  Map<String, dynamic> toJson() => _$WorkflowToJson(this);
}

@JsonSerializable()
class WFStep {
  @JsonKey(name: "id")
  int id;
  @JsonKey(name: "label")
  String label;
  @JsonKey(name: "room_id")
  int roomID;
  @JsonKey(name: "options", includeIfNull: false)
  Map<String, List<WFStep>> options;

  WFStep({this.id, this.label, this.roomID, this.options});

  factory WFStep.fromJson(Map<String, dynamic> json) => _$WFStepFromJson(json);
  Map<String, dynamic> toJSON() => _$WFStepToJson(this);
}
