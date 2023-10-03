import 'package:get/get.dart';
import 'package:my_todo/view/task/detail/task_detail_binding.dart';
import 'package:my_todo/view/task/detail/task_detail_page.dart';

class TaskRouter {
  static List<GetPage> pages = [detail];

  static final detail = GetPage(
      name: '/detail',
      page: () => const TaskInfoPage(),
      binding: TaskInfoBinding());
}
