import 'package:get/get.dart';
import 'package:my_todo/view/qr/qr_controller.dart';

class QRBinding extends Bindings {
  @override
  void dependencies() {
    Get.lazyPut(() => QRController());
  }
}
