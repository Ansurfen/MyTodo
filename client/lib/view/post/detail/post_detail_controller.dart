import 'dart:async';

import 'package:flutter/cupertino.dart';
import 'package:flutter/material.dart';
import 'package:get/get.dart';
import 'package:my_todo/api/post.dart';
import 'package:my_todo/model/entity/image.dart';
import 'package:my_todo/model/entity/post.dart';
import 'package:my_todo/model/vo/post.dart';
import 'package:my_todo/utils/dialog.dart';
import 'package:my_todo/utils/guard.dart';
import 'package:my_todo/mock/post.dart' as mock;

class PostDetailController extends GetxController {
  String username = '';
  bool isMale = false;
  int uid = 0;
  List<MImage> images = [];
  int favorite = 0;
  String content = "";
  late int id;
  String selectedComment = '';
  late GetPostVo data;
  Map<String, PostComment> comments = {};

  PostDetailController({this.username = ''});

  @override
  void onInit() {
    super.onInit();
    id = int.parse(Get.parameters["id"]!);
    fetchAll();
  }

  Future fetchAll() async {
    await fetchPost();
    return fetchComments();
  }

  Future fetchPost() async {
    postDetail(PostDetailRequest(id: id)).then((res) {
      // data = GetPostVo(id, uid, username, isMale, createAt, content, images, favoriteCnt, commentCnt);
    });
  }

  Future fetchComments() async {
    comments.clear();
    if (Guard.isOffline()) {
      for (var e in mock.comments) {
        comments[e.id] = e;
      }
    } else {
      return getPostComment(
              GetPostCommentRequest(pid: id, page: 1, pageSize: 10))
          .then((res) {
        for (var e in res.comments) {
          comments[e.id] = e;
          print(e);
        }
      });
    }
  }

  void handleCommentReply(BuildContext context) {
    showCupertinoModalPopup(
        context: context,
        builder: (BuildContext context) => CupertinoActionSheet(
              message: Column(
                children: [
                  dialogAction(icon: Icons.open_in_new, text: "share".tr),
                  const SizedBox(height: 15),
                  dialogAction(icon: Icons.copy, text: "copy".tr),
                  const SizedBox(height: 15),
                  const Divider(),
                  const SizedBox(height: 15),
                  dialogAction(icon: Icons.warning_amber, text: "report".tr),
                  const SizedBox(height: 15),
                  dialogAction(icon: Icons.delete, text: "delete".tr),
                ],
              ),
            ));
  }
}
