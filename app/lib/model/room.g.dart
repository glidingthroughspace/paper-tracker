// GENERATED CODE - DO NOT MODIFY BY HAND

part of 'room.dart';

// **************************************************************************
// JsonSerializableGenerator
// **************************************************************************

Room _$RoomFromJson(Map<String, dynamic> json) {
  return Room(id: json['ID'] as int, label: json['Label'] as String);
}

Map<String, dynamic> _$RoomToJson(Room instance) =>
    <String, dynamic>{'ID': instance.id, 'Label': instance.label};
