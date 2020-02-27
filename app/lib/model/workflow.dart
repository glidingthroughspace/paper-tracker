import 'package:json_annotation/json_annotation.dart';
import 'package:material_design_icons_flutter/material_design_icons_flutter.dart';
import 'package:paper_tracker/widgets/dropdown.dart';

part 'workflow.g.dart';

@JsonSerializable()
class WorkflowTemplate extends DropdownCapable {
  static const IconData = MdiIcons.clipboardFileOutline;
  static const CompletedIconData = MdiIcons.clipboardCheckOutline;

  @JsonKey(name: "id")
  int id;
  @JsonKey(name: "label")
  String label;
  @JsonKey(name: "steps")
  List<WFStep> steps;

  WorkflowTemplate({this.id, this.label, this.steps});

  factory WorkflowTemplate.fromJson(Map<String, dynamic> json) => _$WorkflowTemplateFromJson(json);
  Map<String, dynamic> toJson() => _$WorkflowTemplateToJson(this);
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

@JsonSerializable()
class WorkflowExec {
  static const IconData = MdiIcons.clipboardTextPlayOutline;

  @JsonKey(name: "id")
  int id;
  @JsonKey(name: "label")
  String label;
  @JsonKey(name: "template_id")
  int templateID;
  @JsonKey(name: "tracker_id")
  int trackerID;
  @JsonKey(name: "completed")
  bool completed;
  @JsonKey(name: "started_on")
  DateTime startedOn;
  @JsonKey(name: "completed_on")
  DateTime completedOn;
  @JsonKey(name: "current_step_id")
  int currentStepID;
  @JsonKey(name: "step_infos", toJson: _stepInfosToJSON)
  Map<int, ExecStepInfo> stepInfos;

  WorkflowExec(
      {this.id,
      this.label,
      this.templateID,
      this.trackerID,
      this.completed,
      this.startedOn,
      this.completedOn,
      this.currentStepID,
      this.stepInfos});

  factory WorkflowExec.fromJson(Map<String, dynamic> json) => _$WorkflowExecFromJson(json);
  Map<String, dynamic> toJSON() => _$WorkflowExecToJson(this);
}

dynamic _stepInfosToJSON(Map<int, ExecStepInfo> stepInfos) {
  return stepInfos?.map((k, e) => MapEntry(k.toString(), e.toJSON()));
}

@JsonSerializable()
class ExecStepInfo {
  @JsonKey(name: "decision")
  String decision;
  @JsonKey(name: "started_on")
  DateTime startedOn;
  @JsonKey(name: "completed_on")
  DateTime completedOn;

  ExecStepInfo({this.decision, this.startedOn, this.completedOn});

  factory ExecStepInfo.fromJson(Map<String, dynamic> json) => _$ExecStepInfoFromJson(json);
  Map<String, dynamic> toJSON() => _$ExecStepInfoToJson(this);
}
