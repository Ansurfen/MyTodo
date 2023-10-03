import 'package:get/get.dart';
import 'package:my_todo/view/user/forget/forget_controller.dart';

class ForgetBinding extends Bindings {
  @override
  void dependencies() {
    Get.put<ForgetController>(ForgetController());
  }
}
