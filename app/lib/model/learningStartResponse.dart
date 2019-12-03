import 'package:json_annotation/json_annotation.dart';
part 'learningStartResponse.g.dart';

@JsonSerializable()
class LearningStartResponse {
  @JsonKey(name: "LearnTimeSec")
  int learnTimeSec;

  LearningStartResponse({this.learnTimeSec});

  factory LearningStartResponse.fromJson(Map<String, dynamic> json) => _$LearningStartResponseFromJson(json);
  Map<String, dynamic> toJson() => _$LearningStartResponseToJson(this);
}