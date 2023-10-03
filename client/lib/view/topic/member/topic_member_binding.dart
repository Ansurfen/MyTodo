import 'package:get/get.dart';
import 'package:my_todo/view/topic/member/topic_member_controller.dart';

class TopicMemberBinding extends Bindings {
  @override
  void dependencies() {
    Get.lazyPut(() => TopicMemberController());
  }
}
