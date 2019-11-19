// GENERATED CODE - DO NOT MODIFY BY HAND

part of 'tracker.dart';

// **************************************************************************
// JsonSerializableGenerator
// **************************************************************************

Tracker _$TrackerFromJson(Map<String, dynamic> json) {
  return Tracker(
      id: json['ID'] as int,
      label: json['Label'] as String,
      lastPoll: json['LastPoll'] as String,
      lastSleepTime: json['LastSleepTime'] as String);
}

Map<String, dynamic> _$TrackerToJson(Tracker instance) => <String, dynamic>{
      'ID': instance.id,
      'Label': instance.label,
      'LastPoll': instance.lastPoll,
      'LastSleepTime': instance.lastSleepTime
    };
