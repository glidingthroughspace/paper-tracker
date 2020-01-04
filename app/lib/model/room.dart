import 'package:flutter/material.dart';
import 'package:json_annotation/json_annotation.dart';
part 'room.g.dart';

@JsonSerializable()
class Room {
 static const IconData = Icons.room;

 @JsonKey(name: "id")
 int id;
 @JsonKey(name: "label")
 String label;
 @JsonKey(name: "is_learned")
 bool isLearned;

 Room({this.id, this.label, this.isLearned = false});

 factory Room.fromJson(Map<String, dynamic> json) => _$RoomFromJson(json);
 Map<String, dynamic> toJson() => _$RoomToJson(this);
}