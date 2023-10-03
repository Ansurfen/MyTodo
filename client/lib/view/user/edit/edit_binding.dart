import 'package:get/get.dart';
import 'package:my_todo/view/user/edit/edit_controller.dart';

class EditBinding extends Bindings {
  @override
  void dependencies() {
    Get.put<EditController>(EditController());
  }
}
