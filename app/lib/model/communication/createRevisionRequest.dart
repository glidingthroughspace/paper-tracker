import 'package:json_annotation/json_annotation.dart';

part 'createRevisionRequest.g.dart';

@JsonSerializable()
class CreateRevisionRequest {
  @JsonKey(name: "revision_label")
  String revisionLabel;

  CreateRevisionRequest({this.revisionLabel});

  factory CreateRevisionRequest.fromJson(Map<String, dynamic> json) => _$CreateRevisionRequestFromJson(json);
  Map<String, dynamic> toJson() => _$CreateRevisionRequestToJson(this);
}
