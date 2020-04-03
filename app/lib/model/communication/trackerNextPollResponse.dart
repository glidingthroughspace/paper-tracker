import 'package:json_annotation/json_annotation.dart';

part 'trackerNextPollResponse.g.dart';

@JsonSerializable()
class TrackerNextPollResponse {
  @JsonKey(name: "next_poll_sec")
  int nextPollSec;

  TrackerNextPollResponse({this.nextPollSec});

  factory TrackerNextPollResponse.fromJson(Map<String, dynamic> json) => _$TrackerNextPollResponseFromJson(json);
  Map<String, dynamic> toJson() => _$TrackerNextPollResponseToJson(this);
}
