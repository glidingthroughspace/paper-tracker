import 'package:json_annotation/json_annotation.dart';
part 'learningStatusResponse.g.dart';

@JsonSerializable()
class LearningStatusResponse {
  @JsonKey(name: "done")
  bool done;
  @JsonKey(name: "ssids")
  List<String > ssids;

  LearningStatusResponse({this.done, this.ssids});

  factory LearningStatusResponse.fromJson(Map<String, dynamic> json) => _$LearningStatusResponseFromJson(json);
  Map<String, dynamic> toJson() => _$LearningStatusResponseToJson(this);
}
