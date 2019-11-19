import 'package:json_annotation/json_annotation.dart';
part 'tracker.g.dart';

@JsonSerializable()
class Tracker {
  @JsonKey(name: "ID")
  int id;
  @JsonKey(name: "Label")
  String label;
  @JsonKey(name: "LastPoll")
  String lastPoll;
  @JsonKey(name: "LastSleepTime")
  String lastSleepTime;

  Tracker({this.id, this.label, this.lastPoll, this.lastSleepTime});

  factory Tracker.fromJson(Map<String, dynamic> json) => _$TrackerFromJson(json);
  Map<String, dynamic> toJson() => _$TrackerToJson(this);
}