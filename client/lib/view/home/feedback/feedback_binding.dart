import 'package:get/get.dart';
import 'package:my_todo/view/home/feedback/feedback_controller.dart';
import 'package:my_todo/view/home/home_controller.dart';


class FeedbackBinding extends Bindings {
  @override
  void dependencies() {
    Get.put<TodoDrawerController>(FeedbackController());
  }
}
