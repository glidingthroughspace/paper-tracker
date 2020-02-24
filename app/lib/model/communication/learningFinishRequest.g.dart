// GENERATED CODE - DO NOT MODIFY BY HAND

part of 'learningFinishRequest.dart';

// **************************************************************************
// JsonSerializableGenerator
// **************************************************************************

LearningFinishRequest _$LearningFinishRequestFromJson(
    Map<String, dynamic> json) {
  return LearningFinishRequest(
    roomID: json['room_id'] as int,
    ssids: (json['ssids'] as List)?.map((e) => e as String)?.toList(),
  );
}

Map<String, dynamic> _$LearningFinishRequestToJson(
        LearningFinishRequest instance) =>
    <String, dynamic>{
      'room_id': instance.roomID,
      'ssids': instance.ssids,
    };
