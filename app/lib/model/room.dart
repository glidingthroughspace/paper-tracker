import 'package:flutter/material.dart';
import 'package:json_annotation/json_annotation.dart';
part 'room.g.dart';

@JsonSerializable()
class Room {
 static const IconData = Icons.room;

 @JsonKey(name: "ID")
 int id;
 @JsonKey(name: "Label")
 String label;
 bool isLearned;

 Room({this.id, this.label, this.isLearned = false});

 factory Room.fromJson(Map<String, dynamic> json) => _$RoomFromJson(json);
 Map<String, dynamic> toJson() => _$RoomToJson(this);
}