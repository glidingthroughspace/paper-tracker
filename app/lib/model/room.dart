import 'package:json_annotation/json_annotation.dart';
part 'room.g.dart';

@JsonSerializable()
class Room {
 @JsonKey(name: "ID")
 int id;
 @JsonKey(name: "Label")
 String label;

 Room({this.id, this.label});

 factory Room.fromJson(Map<String, dynamic> json) => _$RoomFromJson(json);
 Map<String, dynamic> toJson() => _$RoomToJson(this);
}