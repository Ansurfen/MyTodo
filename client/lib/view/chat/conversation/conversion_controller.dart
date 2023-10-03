import 'package:get/get.dart';
import 'package:my_todo/api/chat.dart';
import 'package:my_todo/api/user.dart';
import 'package:my_todo/model/entity/chat.dart';
import 'package:my_todo/model/entity/user.dart';
import 'package:my_todo/utils/dialog.dart';
import 'package:my_todo/utils/guard.dart';

class ConversionController extends GetxController {
  User user = User(0, "", "");
  Rx<List<Chat>> chats = Rx([]);
  int page = 1;

  @override
  void onInit() {
    super.onInit();
    if (Get.arguments != null) {
      user = Get.arguments;
    } else {
      userInfo(int.parse(Get.parameters["id"]!)).then((res) {
        user = res;
        fetch().then((res) {
          chats.value = res.chats;
        });
      });
    }
  }

  Future<GetChatResponse> fetch() {
    return getChat(
        GetChatRequest(from: Guard.user, to: user.id, page: 2, pageSize: 10));
  }

  void sendMessage(Chat msg) {
    msg.from = Guard.user;
    msg.to = user.id;
    addChat(AddChatRequest(msg)).then((value) {}).onError((error, stackTrace) {
      showError(error.toString());
    });
  }
}
