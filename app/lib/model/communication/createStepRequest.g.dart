// GENERATED CODE - DO NOT MODIFY BY HAND

part of 'createStepRequest.dart';

// **************************************************************************
// JsonSerializableGenerator
// **************************************************************************

CreateStepRequest _$CreateStepRequestFromJson(Map<String, dynamic> json) {
  return CreateStepRequest(
    previousStepID: json['prev_step_id'] as int,
    decisionLabel: json['decision_label'] as String,
    step: json['step'] == null
        ? null
        : WFStep.fromJson(json['step'] as Map<String, dynamic>),
  );
}

Map<String, dynamic> _$CreateStepRequestToJson(CreateStepRequest instance) =>
    <String, dynamic>{
      'prev_step_id': instance.previousStepID,
      'decision_label': instance.decisionLabel,
      'step': _stepToJSON(instance.step),
    };
