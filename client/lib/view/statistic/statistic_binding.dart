import 'package:get/get.dart';
import 'package:my_todo/view/statistic/statistic_controller.dart';

class StatisticBinding extends Bindings {
  @override
  void dependencies() {
    Get.put<StatisticController>(StatisticController());
  }
}
