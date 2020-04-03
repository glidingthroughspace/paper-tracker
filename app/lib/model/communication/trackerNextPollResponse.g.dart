// GENERATED CODE - DO NOT MODIFY BY HAND

part of 'trackerNextPollResponse.dart';

// **************************************************************************
// JsonSerializableGenerator
// **************************************************************************

TrackerNextPollResponse _$TrackerNextPollResponseFromJson(
    Map<String, dynamic> json) {
  return TrackerNextPollResponse(
    nextPollSec: json['next_poll_sec'] as int,
  );
}

Map<String, dynamic> _$TrackerNextPollResponseToJson(
        TrackerNextPollResponse instance) =>
    <String, dynamic>{
      'next_poll_sec': instance.nextPollSec,
    };
