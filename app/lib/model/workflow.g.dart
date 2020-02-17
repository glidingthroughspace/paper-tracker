// GENERATED CODE - DO NOT MODIFY BY HAND

part of 'workflow.dart';

// **************************************************************************
// JsonSerializableGenerator
// **************************************************************************

Workflow _$WorkflowFromJson(Map<String, dynamic> json) {
  return Workflow(id: json['id'] as int)
    ..label = json['label'] as String
    ..isTemplate = json['is_template'] as bool
    ..steps = (json['steps'] as List)
        ?.map((e) =>
            e == null ? null : WFStep.fromJson(e as Map<String, dynamic>))
        ?.toList();
}

Map<String, dynamic> _$WorkflowToJson(Workflow instance) => <String, dynamic>{
      'id': instance.id,
      'label': instance.label,
      'is_template': instance.isTemplate,
      'steps': instance.steps
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
                ?.map((e) => e == null
                    ? null
                    : WFStep.fromJson(e as Map<String, dynamic>))
                ?.toList()),
      ));
}

Map<String, dynamic> _$WFStepToJson(WFStep instance) => <String, dynamic>{
      'id': instance.id,
      'label': instance.label,
      'room_id': instance.roomID,
      'options': instance.options
    };
