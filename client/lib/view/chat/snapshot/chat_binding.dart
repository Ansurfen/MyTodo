import 'package:get/get.dart';
import 'package:my_todo/view/chat/snapshot/chat_controller.dart';

class CharBinding extends Bindings {
  @override
  void dependencies() {
    Get.lazyPut(() => ChatController());
  }
}
