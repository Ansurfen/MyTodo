import 'package:my_todo/model/dto/post.dart';
import 'package:my_todo/utils/guard.dart';

class GetPostVo extends GetPostDto {
  late List<String> imageUri;

  GetPostVo(super.id, super.uid, super.username, super.isMale, super.createAt,
      super.content, super.images, super.favoriteCnt, super.commentCnt)
      : imageUri = [];

  GetPostVo.fromDto(GetPostDto v)
      : super(v.id, v.uid, v.username, v.isMale, v.createAt, v.content,
            v.images, v.favoriteCnt, v.commentCnt) {
    imageUri =
        v.images.map((e) => "${Guard.server}/post/image/${e.path}").toList();
  }
}
