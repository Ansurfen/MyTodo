import 'package:get/get.dart';
import 'package:my_todo/view/user/sign/sign_controller.dart';

class SignBinding extends Bindings {
  @override
  void dependencies() {
    Get.lazyPut(() => SignController());
  }
}
