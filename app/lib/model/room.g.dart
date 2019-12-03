// GENERATED CODE - DO NOT MODIFY BY HAND

part of 'room.dart';

// **************************************************************************
// JsonSerializableGenerator
// **************************************************************************

Room _$RoomFromJson(Map<String, dynamic> json) {
  return Room(
      id: json['id'] as int,
      label: json['label'] as String,
      isLearned: json['is_learned'] as bool);
}

Map<String, dynamic> _$RoomToJson(Room instance) => <String, dynamic>{
      'id': instance.id,
      'label': instance.label,
      'is_learned': instance.isLearned
    };
