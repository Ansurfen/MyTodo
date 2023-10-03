import 'package:get/get.dart';
import 'package:my_todo/view/task/detail/task_detail_controller.dart';

class TaskInfoBinding extends Bindings {
  @override
  void dependencies() {
    Get.lazyPut<TaskInfoController>(() => TaskInfoController());
  }
}
