import 'dart:convert';

import 'package:dio/dio.dart';
import 'package:json_annotation/json_annotation.dart';
import 'package:my_todo/api/response.dart';
import 'package:my_todo/model/entity/chat.dart';
import 'package:my_todo/utils/guard.dart';
import 'package:my_todo/utils/json.dart';
import 'package:my_todo/utils/net.dart';

part 'chat.g.dart';

class AddChatRequest {
  Chat chat;

  AddChatRequest(this.chat);

  FormData toFormData() {
    FormData formData = FormData();
    formData.fields.addAll({
      'from': '${chat.from}',
      'to': '${chat.to}',
      'content': chat.content.isNotEmpty ? chat.content[0] : " ",
      "reply": "${chat.reply}"
    }.entries);
    return formData;
  }
}

Future addChat(AddChatRequest req) async {
  return await HTTP.post('/chat/add',
      data: req.toFormData(),
      options: Options(headers: {'x-token': Guard.jwt}));
}

Future chatFriend() async {
  return await HTTP.get('/chat/friend',
      options: Options(headers: {'x-token': Guard.jwt}));
}

@JsonSerializable()
class GetChatRequest {
  @JsonKey(name: "from")
  int from;

  @JsonKey(name: "to")
  int to;

  @JsonKey(name: "page")
  int page;

  @JsonKey(name: "pageSize")
  int pageSize;

  GetChatRequest(
      {required this.from,
      required this.to,
      required this.page,
      required this.pageSize});

  FormData toFormData() {
    FormData formData = FormData();
    formData.fields.addAll({
      'from': "$from",
      'to': "$to",
      'page': "$page",
      'pageSize': "$pageSize"
    }.entries);
    return formData;
  }

  JsonObject toJson() => _$GetChatRequestToJson(this);
}

class GetChatResponse extends BaseResponse {
  List<Chat> chats = [];

  GetChatResponse() : super({});

  GetChatResponse.fromResponse(Response res) : super(res.data) {
    if (res.data["data"]["chats"] != null) {
      chats = (res.data["data"]["chats"] as List)
          .map((e) => Chat.fromJson(e))
          .toList();
    }
  }
}

Future<GetChatResponse> getChat(GetChatRequest req) async {
  return GetChatResponse.fromResponse(await HTTP.post('/chat/get',
      data: jsonEncode(req),
      options: Options(headers: {'x-token': Guard.jwt})));
}
