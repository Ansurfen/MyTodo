import 'package:get/get.dart';
import 'package:my_todo/view/notification/notification_controller.dart';

class NotificationBinding extends Bindings {
  @override
  void dependencies() {
    Get.lazyPut(() => NotificationController());
  }
}
