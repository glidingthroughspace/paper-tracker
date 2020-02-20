import 'package:flutter/material.dart';
import 'package:json_annotation/json_annotation.dart';
import 'package:paper_tracker/widgets/dropdown.dart';

part 'tracker.g.dart';

@JsonSerializable()
class Tracker implements DropdownCapable {
  static const IconData = Icons.track_changes;

  @JsonKey(name: "id")
  int id;
  @JsonKey(name: "label")
  String label;
  @JsonKey(name: "last_poll")
  DateTime lastPoll;
  @JsonKey(name: "last_sleep_time")
  DateTime lastSleepTime;
  @JsonKey(name: "status")
  TrackerStatus status;

  Tracker({this.id, this.label, this.lastPoll, this.lastSleepTime, this.status});

  factory Tracker.fromJson(Map<String, dynamic> json) => _$TrackerFromJson(json);

  Map<String, dynamic> toJson() => _$TrackerToJson(this);
}

enum TrackerStatus {
  @JsonValue(1)
  Idle,
  @JsonValue(2)
  Learning,
  @JsonValue(3)
  LearningFinished,
  @JsonValue(4)
  Tracking
}
