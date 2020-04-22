import 'package:flutter/material.dart';
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
  @JsonKey(name: "first_revision_id")
  int firstRevisionID;
  @JsonKey(name: "steps", toJson: _stepsToJson)
  List<WFStep> steps;
  @JsonKey(name: "step_editing_locked")
  bool stepEditingLocked;

  WorkflowTemplate({this.id, this.label, this.firstRevisionID, this.steps, this.stepEditingLocked});

  factory WorkflowTemplate.fromJson(Map<String, dynamic> json) => _$WorkflowTemplateFromJson(json);
  Map<String, dynamic> toJson() => _$WorkflowTemplateToJson(this);
}

dynamic _stepsToJson(List<WFStep> steps) {
  if (steps != null)
    return steps.map((step) => step.toJSON()).toList();
  else
    return null;
}

@JsonSerializable()
class WFStep {
  static const CurrentStepColor = Colors.amber;

  @JsonKey(name: "id")
  int id;
  @JsonKey(name: "label")
  String label;
  @JsonKey(name: "room_ids")
  List<int> roomIDs;
  @JsonKey(name: "options", includeIfNull: false, toJson: _optionsToJson)
  Map<String, List<WFStep>> options;

  WFStep({this.id, this.label, this.roomIDs, this.options});

  factory WFStep.fromJson(Map<String, dynamic> json) => _$WFStepFromJson(json);
  Map<String, dynamic> toJSON() => _$WFStepToJson(this);
}

dynamic _optionsToJson(Map<String, List<WFStep>> options) {
  if (options != null)
    return options.map((decision, steps) => MapEntry(decision, _stepsToJson(steps)));
  else
    return null;
}

@JsonSerializable()
class WorkflowExec {
  static const IconData = MdiIcons.clipboardTextPlayOutline;
  static const CompletedColor = Colors.teal;

  @JsonKey(name: "id")
  int id;
  @JsonKey(name: "label")
  String label;
  @JsonKey(name: "template_id")
  int templateID;
  @JsonKey(name: "tracker_id")
  int trackerID;
  @JsonKey(name: "status")
  WorkflowExecStatus status;
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
      this.status,
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

enum WorkflowExecStatus {
  @JsonValue(1)
  Running,
  @JsonValue(2)
  Finished,
  @JsonValue(3)
  Cancelled,
}

extension WorkflowExecStatusExtension on WorkflowExecStatus {
  String get label {
    return this.toString().substring(this.toString().indexOf('.') + 1);
  }

  IconData get icon {
    switch (this) {
      case WorkflowExecStatus.Running:
        return Icons.play_circle_filled;
      case WorkflowExecStatus.Finished:
        return Icons.done;
      case WorkflowExecStatus.Cancelled:
        return Icons.cancel;
    }
    return Icons.adb;
  }
}

@JsonSerializable()
class ExecStepInfo {
  @JsonKey(name: "decision")
  String decision;
  @JsonKey(name: "started_on")
  DateTime startedOn;
  @JsonKey(name: "completed_on")
  DateTime completedOn;
  @JsonKey(name: "skipped")
  bool skipped;

  ExecStepInfo({this.decision, this.startedOn, this.completedOn, this.skipped});

  factory ExecStepInfo.fromJson(Map<String, dynamic> json) => _$ExecStepInfoFromJson(json);
  Map<String, dynamic> toJSON() => _$ExecStepInfoToJson(this);
}
