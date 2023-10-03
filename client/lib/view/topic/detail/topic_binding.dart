import 'package:get/get.dart';
import 'package:my_todo/view/topic/detail/topic_controller.dart';

class TopicBinding extends Bindings {
  @override
  void dependencies() {
    Get.lazyPut(() => TopicController());
  }
}
