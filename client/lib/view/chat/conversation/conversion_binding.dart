import 'package:get/get.dart';
import 'package:my_todo/view/chat/conversation/conversion_controller.dart';

class ConversionBinding extends Bindings {
  @override
  void dependencies() {
    Get.lazyPut(() => ConversionController());
  }
}
