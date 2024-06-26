import 'package:get/get.dart';
import 'package:my_todo/api/topic.dart';
import 'package:my_todo/model/entity/topic.dart';

class TopicMemberController extends GetxController {
  late int id;
  Rx<List<TopicMember>> members = Rx([]);

  @override
  void onInit() {
    super.onInit();
    id = int.parse(Get.parameters["id"]!);
    getSubscribedMember(GetSubscribedMemberRequest(id: id)).then((res) {
      members.value = res.members;
      members.refresh();
    });
  }
}
