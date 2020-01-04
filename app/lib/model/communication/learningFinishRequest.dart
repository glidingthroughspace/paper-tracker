import 'package:json_annotation/json_annotation.dart';
part 'learningFinishRequest.g.dart';

@JsonSerializable()
class LearningFinishRequest {
  @JsonKey(name: "room_id")
  int roomID;
  @JsonKey(name: "ssids")
  List<String> ssids;

  LearningFinishRequest({this.roomID, this.ssids});

  factory LearningFinishRequest.fromJson(Map<String, dynamic> json) => _$LearningFinishRequestFromJson(json);
  Map<String, dynamic> toJson() => _$LearningFinishRequestToJson(this);
}