import 'package:flutter/material.dart';
import 'package:json_annotation/json_annotation.dart';
import 'package:material_design_icons_flutter/material_design_icons_flutter.dart';
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
  @JsonKey(name: "last_room")
  int lastRoom;
  @JsonKey(name: "status")
  TrackerStatus status;
  @JsonKey(name: "battery_percentage")
  int batteryPercentage;
  @JsonKey(name: "is_charging", defaultValue: false)
  bool isCharging;

  Tracker({this.id, this.label, this.lastPoll, this.lastRoom, this.status, this.batteryPercentage, this.isCharging});

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

extension TrackerStatusExtension on TrackerStatus {
  String get label {
    if (this == TrackerStatus.LearningFinished) return "Finishing learning";

    return this.toString().substring(this.toString().indexOf('.') + 1);
  }

  IconData get icon {
    switch (this) {
      case TrackerStatus.Idle:
        return MdiIcons.progressClock;
      case TrackerStatus.Learning:
        return Icons.school;
      case TrackerStatus.LearningFinished:
        return Icons.school;
      case TrackerStatus.Tracking:
        return Icons.track_changes;
    }
    return Icons.adb;
  }
}
