// GENERATED CODE - DO NOT MODIFY BY HAND

part of 'workflow.dart';

// **************************************************************************
// JsonSerializableGenerator
// **************************************************************************

WorkflowTemplate _$WorkflowTemplateFromJson(Map<String, dynamic> json) {
  return WorkflowTemplate(
    id: json['id'] as int,
    label: json['label'] as String,
    firstRevisionID: json['first_revision_id'] as int,
    steps: (json['steps'] as List)
        ?.map((e) =>
            e == null ? null : WFStep.fromJson(e as Map<String, dynamic>))
        ?.toList(),
    stepEditingLocked: json['step_editing_locked'] as bool,
  );
}

Map<String, dynamic> _$WorkflowTemplateToJson(WorkflowTemplate instance) =>
    <String, dynamic>{
      'id': instance.id,
      'label': instance.label,
      'first_revision_id': instance.firstRevisionID,
      'steps': _stepsToJson(instance.steps),
      'step_editing_locked': instance.stepEditingLocked,
    };

WFStep _$WFStepFromJson(Map<String, dynamic> json) {
  return WFStep(
    id: json['id'] as int,
    label: json['label'] as String,
    roomIDs: (json['room_ids'] as List)?.map((e) => e as int)?.toList(),
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
    'room_ids': instance.roomIDs,
  };

  void writeNotNull(String key, dynamic value) {
    if (value != null) {
      val[key] = value;
    }
  }

  writeNotNull('options', _optionsToJson(instance.options));
  return val;
}

WorkflowExec _$WorkflowExecFromJson(Map<String, dynamic> json) {
  return WorkflowExec(
    id: json['id'] as int,
    label: json['label'] as String,
    templateID: json['template_id'] as int,
    trackerID: json['tracker_id'] as int,
    status: _$enumDecodeNullable(_$WorkflowExecStatusEnumMap, json['status']),
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
      'status': _$WorkflowExecStatusEnumMap[instance.status],
      'started_on': instance.startedOn?.toIso8601String(),
      'completed_on': instance.completedOn?.toIso8601String(),
      'current_step_id': instance.currentStepID,
      'step_infos': _stepInfosToJSON(instance.stepInfos),
    };

T _$enumDecode<T>(
  Map<T, dynamic> enumValues,
  dynamic source, {
  T unknownValue,
}) {
  if (source == null) {
    throw ArgumentError('A value must be provided. Supported values: '
        '${enumValues.values.join(', ')}');
  }

  final value = enumValues.entries
      .singleWhere((e) => e.value == source, orElse: () => null)
      ?.key;

  if (value == null && unknownValue == null) {
    throw ArgumentError('`$source` is not one of the supported values: '
        '${enumValues.values.join(', ')}');
  }
  return value ?? unknownValue;
}

T _$enumDecodeNullable<T>(
  Map<T, dynamic> enumValues,
  dynamic source, {
  T unknownValue,
}) {
  if (source == null) {
    return null;
  }
  return _$enumDecode<T>(enumValues, source, unknownValue: unknownValue);
}

const _$WorkflowExecStatusEnumMap = {
  WorkflowExecStatus.Running: 1,
  WorkflowExecStatus.Finished: 2,
  WorkflowExecStatus.Cancelled: 3,
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
    skipped: json['skipped'] as bool,
  );
}

Map<String, dynamic> _$ExecStepInfoToJson(ExecStepInfo instance) =>
    <String, dynamic>{
      'decision': instance.decision,
      'started_on': instance.startedOn?.toIso8601String(),
      'completed_on': instance.completedOn?.toIso8601String(),
      'skipped': instance.skipped,
    };
