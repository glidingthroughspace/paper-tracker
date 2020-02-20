// GENERATED CODE - DO NOT MODIFY BY HAND

part of 'learningStatusResponse.dart';

// **************************************************************************
// JsonSerializableGenerator
// **************************************************************************

LearningStatusResponse _$LearningStatusResponseFromJson(
    Map<String, dynamic> json) {
  return LearningStatusResponse(
    done: json['done'] as bool,
    ssids: (json['ssids'] as List)?.map((e) => e as String)?.toList(),
  );
}

Map<String, dynamic> _$LearningStatusResponseToJson(
        LearningStatusResponse instance) =>
    <String, dynamic>{
      'done': instance.done,
      'ssids': instance.ssids,
    };
