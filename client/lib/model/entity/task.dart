// Copyright 2025 The MyTodo Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.
import 'package:json_annotation/json_annotation.dart';

part 'task.g.dart';

@JsonSerializable()
class Task {
  @JsonKey(name: "id", defaultValue: 0, required: false)
  int? id;

  @JsonKey(name: "user", defaultValue: 0, required: false)
  int? user;

  @JsonKey(name: "name", defaultValue: "", required: true)
  String name;

  @JsonKey(name: "desc", defaultValue: "", required: true)
  String desc;

  @JsonKey(name: "departure", defaultValue: 0, required: true)
  int startAt;

  @JsonKey(name: "arrival", defaultValue: 0, required: true)
  int endAt;

  Task(this.name, this.desc, this.startAt, this.endAt, {this.id, this.user});

  factory Task.fromJson(Map<String, dynamic> json) => _$TaskFromJson(json);

  Map<String, dynamic> toJson() => _$TaskToJson(this);

  @override
  String toString() {
    return """{"id": $id, "name": "$name", "desc": "$desc", "startAt": $startAt, "endAt": $endAt}""";
  }
}

@JsonSerializable()
class TaskCondition {
  @JsonKey(name: "type")
  int type;

  @JsonKey(name: "param")
  String param;

  TaskCondition(this.type, this.param);

  factory TaskCondition.fromJson(Map<String, Object?> json) =>
      _$TaskConditionFromJson(json);

  Map<String, Object?> toJson() => _$TaskConditionToJson(this);

  static conditionsFromJson(List conds) =>
      conds.map((e) => TaskCondition.fromJson(e)).toList();
}

enum TaskCondType { hand, timer, locale, file, image, content, qr }
