import 'package:get/get.dart';
import 'package:my_todo/view/user/license/license_controller.dart';

class LicenseBinding extends Bindings {
  @override
  void dependencies() {
    Get.put<LicenseController>(LicenseController());
  }
}
