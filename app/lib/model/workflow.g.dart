// GENERATED CODE - DO NOT MODIFY BY HAND

part of 'workflow.dart';

// **************************************************************************
// JsonSerializableGenerator
// **************************************************************************

WorkflowTemplate _$WorkflowTemplateFromJson(Map<String, dynamic> json) {
  return WorkflowTemplate(
    id: json['id'] as int,
  )
    ..label = json['label'] as String
    ..steps = (json['steps'] as List)
        ?.map((e) =>
            e == null ? null : WFStep.fromJson(e as Map<String, dynamic>))
        ?.toList();
}

Map<String, dynamic> _$WorkflowTemplateToJson(WorkflowTemplate instance) =>
    <String, dynamic>{
      'id': instance.id,
      'label': instance.label,
      'steps': instance.steps,
    };

WFStep _$WFStepFromJson(Map<String, dynamic> json) {
  return WFStep(
    id: json['id'] as int,
    label: json['label'] as String,
    roomID: json['room_id'] as int,
    options: (json['options'] as Map<String, dynamic>)?.map(
      (k, e) => MapEntry(
          k,
          (e as List)
              ?.map((e) =>
                  e == null ? null : WFStep.fromJson(e as Map<String, dynamic>))
              ?.toList()),
    ),
  );
}

Map<String, dynamic> _$WFStepToJson(WFStep instance) {
  final val = <String, dynamic>{
    'id': instance.id,
    'label': instance.label,
    'room_id': instance.roomID,
  };

  void writeNotNull(String key, dynamic value) {
    if (value != null) {
      val[key] = value;
    }
  }

  writeNotNull('options', instance.options);
  return val;
}

WorkflowExec _$WorkflowExecFromJson(Map<String, dynamic> json) {
  return WorkflowExec(
    id: json['id'] as int,
    label: json['label'] as String,
    templateID: json['template_id'] as int,
    trackerID: json['tracker_id'] as int,
    compeleted: json['completed'] as bool,
    startedOn: json['started_on'] == null
        ? null
        : DateTime.parse(json['started_on'] as String),
    completedOn: json['completed_on'] == null
        ? null
        : DateTime.parse(json['completed_on'] as String),
    currentStepID: json['current_step_id'] as int,
    stepInfos: (json['step_infos'] as Map<String, dynamic>)?.map(
      (k, e) => MapEntry(int.parse(k),
          e == null ? null : ExecStepInfo.fromJson(e as Map<String, dynamic>)),
    ),
  );
}

Map<String, dynamic> _$WorkflowExecToJson(WorkflowExec instance) =>
    <String, dynamic>{
      'id': instance.id,
      'label': instance.label,
      'template_id': instance.templateID,
      'tracker_id': instance.trackerID,
      'completed': instance.compeleted,
      'started_on': instance.startedOn?.toIso8601String(),
      'completed_on': instance.completedOn?.toIso8601String(),
      'current_step_id': instance.currentStepID,
      'step_infos': _stepInfosToJSON(instance.stepInfos),
    };

ExecStepInfo _$ExecStepInfoFromJson(Map<String, dynamic> json) {
  return ExecStepInfo(
    decision: json['decision'] as String,
    startedOn: json['started_on'] == null
        ? null
        : DateTime.parse(json['started_on'] as String),
    completedOn: json['completed_on'] == null
        ? null
        : DateTime.parse(json['completed_on'] as String),
  );
}

Map<String, dynamic> _$ExecStepInfoToJson(ExecStepInfo instance) =>
    <String, dynamic>{
      'decision': instance.decision,
      'started_on': instance.startedOn?.toIso8601String(),
      'completed_on': instance.completedOn?.toIso8601String(),
    };
