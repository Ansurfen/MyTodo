import 'package:dio/dio.dart';
import 'package:json_annotation/json_annotation.dart';
import 'package:my_todo/api/response.dart';
import 'package:my_todo/model/dao/post.dart';
import 'package:my_todo/model/dto/post.dart';
import 'package:my_todo/model/entity/post.dart';
import 'package:my_todo/utils/guard.dart';
import 'package:my_todo/utils/net.dart';
import 'package:my_todo/utils/picker.dart';

class GetPostRequest {
  int page;
  int count;

  GetPostRequest(this.page, this.count);
}

class GetPostResponse extends BaseResponse {
  late final List<GetPostDto> data;

  GetPostResponse(this.data) : super({});

  GetPostResponse.fromResponse(Response res)
      : data = (res.data['data']['posts'] as List)
            .map((e) => GetPostDto.fromJson(e))
            .toList(),
        super(res.data);
}

Future<GetPostResponse> getPost(GetPostRequest req) async {
  if (Guard.isOffline()) {
    await PostDao.findMany();
    // res.map((e) => null)
    return GetPostResponse((await PostDao.findMany())
        .map((e) => GetPostDto(
            e.id ?? 0,
            e.uid,
            "",
            true,
            DateTime.fromMicrosecondsSinceEpoch(e.createAt),
            e.content,
            [],
            0,
            0))
        .toList());
  }
  return GetPostResponse.fromResponse(await HTTP
      .get("/post/get", queryParams: {'page': req.page, 'count': req.count}));
}

// @FormDataSerializable()
class CreatePostRequest {
  // @FormDataKey(name: "uid")
  int user;

  // @FormDataKey(name: "content")
  String content;

  // @FormDataKey(toFormData: )
  List<TFile> images;

  // static _prepare(List<TFile> images) =>
  //     images.map((e) async => MapEntry("files", await e.m)).toList();

  CreatePostRequest(this.user, this.content, this.images);

  Future<FormData> toFormData() async {
    FormData formData = FormData();
    formData.fields.addAll({
      'uid': "$user",
      'content': content,
    }.entries);
    for (TFile img in images) {
      formData.files.add(MapEntry("files", await img.m));
    }
    return formData;
  }
}

class CreatePostResponse extends BaseResponse {
  CreatePostResponse() : super({});

  CreatePostResponse.fromResponse(Response res) : super(res.data);
}

Future<CreatePostResponse> createPost(CreatePostRequest req) async {
  // if (Guard.isOffline()) {
  //   int now = DateTime.now().microsecondsSinceEpoch;
  //   Post c = Post(Guard.user, req.content, now, 0);
  //   await PostDao.create(c);
  //   Guard.eventBus.fire(c);
  //   return CreatePostResponse();
  // }
  if (Guard.isOffline()) {
    int now = DateTime.now().microsecondsSinceEpoch;
    PostDao.create(Post(Guard.user, req.content, now, 0, []));
    return CreatePostResponse();
  }

  return CreatePostResponse.fromResponse(await HTTP.post("/post/add",
      data: await req.toFormData(),
      options: Options(headers: {
        "x-token": Guard.jwt,
      })));
}

class GetPostCommentRequest {
  int pid;
  int page;
  int pageSize;

  GetPostCommentRequest(
      {required this.pid, required this.page, required this.pageSize});

  FormData toFormData() {
    FormData formData = FormData();
    formData.fields.addAll(
        {"pid": "$pid", "page": "$page", "pageSize": "$pageSize"}.entries);
    return formData;
  }
}

class GetPostCommentResponse extends BaseResponse {
  late final List<PostComment> comments;

  GetPostCommentResponse() : super({});

  GetPostCommentResponse.fromResponse(Response res)
      : comments = (res.data["data"]["comments"] as List)
            .map((e) => PostComment.fromJson(e))
            .toList(),
        super(res.data);
}

Future<GetPostCommentResponse> getPostComment(GetPostCommentRequest req) async {
  if (Guard.isOffline()) {
    return GetPostCommentResponse();
  }
  return GetPostCommentResponse.fromResponse(
      await HTTP.post('/post/comment/get', data: req.toFormData()));
}

class PostCommentFavoriteCountRequest {
  String id;

  PostCommentFavoriteCountRequest({required this.id});

  FormData toFormData() {
    FormData formData = FormData();
    formData.fields.addAll({'comment_id': id}.entries);
    return formData;
  }
}

class PostCommentFavoriteCountResponse extends BaseResponse {
  PostCommentFavoriteCountResponse() : super({});

  PostCommentFavoriteCountResponse.fromResponse(Response res) : super(res.data);
}

Future<PostCommentFavoriteCountResponse> postCommentFavoriteCount(
    PostCommentFavoriteCountRequest req) async {
  if (Guard.isOffline()) {
    return PostCommentFavoriteCountResponse();
  }
  return PostCommentFavoriteCountResponse.fromResponse(
      await HTTP.post('/post/comment/favoriteCount', data: req.toFormData()));
}

class PostDetailRequest {
  int id;

  PostDetailRequest({required this.id});
}

@JsonSerializable()
class PostDetailResponse extends BaseResponse {
  @JsonKey(name: "username", defaultValue: '')
  late String username;

  @JsonKey(name: "favorite", defaultValue: 0)
  late int favorite;

  @JsonKey(name: "uid", defaultValue: 0)
  late int uid;

  PostDetailResponse(
      {required this.username, required this.favorite, required this.uid})
      : super({});

  PostDetailResponse.fromResponse(Response res) : super(res.data) {
    username = res.data["data"]["username"];
    uid = res.data["data"]["uid"];
    favorite = res.data["data"]["favorite"];
  }
}

Future<PostDetailResponse> postDetail(PostDetailRequest req) async {
  return PostDetailResponse.fromResponse(
      await HTTP.get('/post/detail/${req.id}'));
}
