// GENERATED CODE - DO NOT MODIFY BY HAND

part of 'tracker.dart';

// **************************************************************************
// JsonSerializableGenerator
// **************************************************************************

Tracker _$TrackerFromJson(Map<String, dynamic> json) {
  return Tracker(
    id: json['id'] as int,
    label: json['label'] as String,
    lastPoll: json['last_poll'] == null
        ? null
        : DateTime.parse(json['last_poll'] as String),
    lastSleepTime: json['last_sleep_time'] == null
        ? null
        : DateTime.parse(json['last_sleep_time'] as String),
    status: _$enumDecodeNullable(_$TrackerStatusEnumMap, json['status']),
  );
}

Map<String, dynamic> _$TrackerToJson(Tracker instance) => <String, dynamic>{
      'id': instance.id,
      'label': instance.label,
      'last_poll': instance.lastPoll?.toIso8601String(),
      'last_sleep_time': instance.lastSleepTime?.toIso8601String(),
      'status': _$TrackerStatusEnumMap[instance.status],
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

const _$TrackerStatusEnumMap = {
  TrackerStatus.Idle: 1,
  TrackerStatus.Learning: 2,
  TrackerStatus.LearningFinished: 3,
  TrackerStatus.Tracking: 4,
};
