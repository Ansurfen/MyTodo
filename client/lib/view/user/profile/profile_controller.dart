import 'package:get/get.dart';
import 'package:my_todo/api/user.dart';
import 'package:my_todo/model/entity/user.dart';

class ProfileController extends GetxController {
  late int id;
  late Rx<User> user;

  @override
  void onInit() {
    super.onInit();
    id = int.parse(Get.parameters["id"]!);
    user = User(id, "", "").obs;
    Future.delayed(Duration.zero, () {
      userInfo(id).then((u) => user.value = u);
    });
  }

  void follow() {

  }
}
