import 'package:get/get.dart';
import 'package:my_todo/view/home/log/log_controller.dart';

class LogBinding extends Bindings {
  @override
  void dependencies() {
    Get.lazyPut(() => LogController());
  }
}
