import 'package:flutter/material.dart';
import 'package:json_annotation/json_annotation.dart';
import 'package:paper_tracker/widgets/dropdown.dart';

part 'room.g.dart';

@JsonSerializable()
class Room implements DropdownCapable {
  static const IconData = Icons.room;

  @JsonKey(name: "id")
  int id;
  @JsonKey(name: "label")
  String label;
  @JsonKey(name: "is_learned")
  bool isLearned;
  @JsonKey(name: "delete_locked")
  bool deleteLocked;

  Room({this.id, this.label, this.isLearned = false});

  factory Room.fromJson(Map<String, dynamic> json) => _$RoomFromJson(json);
  Map<String, dynamic> toJson() => _$RoomToJson(this);
}
