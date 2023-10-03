import 'package:get/get.dart';
import 'package:my_todo/view/post/detail/post_detail_controller.dart';

class PostDetailPageBinding extends Bindings {
  @override
  void dependencies() {
    Get.lazyPut(() => PostDetailController());
  }
}
