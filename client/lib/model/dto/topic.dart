import 'package:json_annotation/json_annotation.dart';

part 'topic.g.dart';

@JsonSerializable()
class GetTopicDto {
  @JsonKey(name: "id", defaultValue: 0)
  int id;

  @JsonKey(name: "created_at")
  DateTime createAt;

  @JsonKey(name: "deleted_at")
  DateTime deleteAt;

  @JsonKey(name: "name")
  String name;

  @JsonKey(name: "desc")
  String desc;

  // int creator;

  @JsonKey(name: "invite_code")
  String inviteCode;

  GetTopicDto(this.id, this.createAt, this.deleteAt, this.name, this.desc,
      this.inviteCode);

  factory GetTopicDto.fromJson(Map<String, Object?> json) =>
      _$GetTopicDtoFromJson(json);
}
