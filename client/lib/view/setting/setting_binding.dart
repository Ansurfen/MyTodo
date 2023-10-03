import 'package:get/get.dart';
import 'package:my_todo/view/setting/setting_controller.dart';

class SettingBinding extends Bindings {
  @override
  void dependencies() {
    Get.put<SettingController>(SettingController());
  }
}
