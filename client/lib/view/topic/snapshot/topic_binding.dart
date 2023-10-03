import 'package:get/get.dart';
import 'package:my_todo/view/topic/snapshot/topic_controller.dart';

class SubscribeBinding extends Bindings {
  @override
  void dependencies() {
    Get.lazyPut(() => TopicSnapshotController());
  }
}
