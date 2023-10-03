import 'package:get/get.dart';
import 'package:my_todo/view/home/about/about_controller.dart';
import 'package:my_todo/view/home/home_controller.dart';

class AboutBinding extends Bindings {
  @override
  void dependencies() {
    Get.put<TodoDrawerController>(AboutController());
  }
}
