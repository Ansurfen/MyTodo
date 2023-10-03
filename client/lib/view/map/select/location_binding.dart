import 'package:get/get.dart';
import 'package:my_todo/view/map/select/location_controller.dart';

class LocationBinding extends Bindings {
  @override
  void dependencies() {
    Get.lazyPut(() => LocationController());
  }
}
