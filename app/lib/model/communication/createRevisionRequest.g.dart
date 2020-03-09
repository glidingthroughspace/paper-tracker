// GENERATED CODE - DO NOT MODIFY BY HAND

part of 'createRevisionRequest.dart';

// **************************************************************************
// JsonSerializableGenerator
// **************************************************************************

CreateRevisionRequest _$CreateRevisionRequestFromJson(
    Map<String, dynamic> json) {
  return CreateRevisionRequest(
    revisionLabel: json['revision_label'] as String,
  );
}

Map<String, dynamic> _$CreateRevisionRequestToJson(
        CreateRevisionRequest instance) =>
    <String, dynamic>{
      'revision_label': instance.revisionLabel,
    };
