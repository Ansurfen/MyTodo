import 'package:emoji_picker_flutter/emoji_picker_flutter.dart';
import 'package:flutter/material.dart';
import 'package:get/get.dart';
import 'package:my_todo/component/container/empty_container.dart';
import 'package:my_todo/component/icon.dart';
import 'package:my_todo/component/image.dart';
import 'package:my_todo/component/input.dart';
import 'package:my_todo/component/refresh.dart';
import 'package:my_todo/component/scaffold.dart';
import 'package:my_todo/model/entity/chat.dart';
import 'package:my_todo/router/provider.dart';
import 'package:my_todo/theme/provider.dart';
import 'package:my_todo/utils/guard.dart';
import 'package:my_todo/view/chat/conversation/chat_container.dart';
import 'package:my_todo/view/chat/conversation/conversion_controller.dart';
import 'package:my_todo/theme/color.dart';

class Conversation extends StatefulWidget {
  const Conversation({super.key});

  @override
  State<Conversation> createState() => _ConversationState();
}

class _ConversationState extends State<Conversation> {
  List conversation = [];
  List<Chat> chats = [];
  TodoInputController todoInputController =
      TodoInputController(TextEditingController(), TextEditingController());
  ConversionController controller = Get.find<ConversionController>();
  Rx<bool> show = false.obs;
  String replyID = '';
  @override
  void dispose() {
    todoInputController.dispose();
    super.dispose();
  }

  @override
  Widget build(BuildContext context) {
    final key = GlobalKey<EmojiPickerState>();
    todoInputController.defaultConfig(context);
    return Scaffold(
      backgroundColor: Theme.of(context).colorScheme.primary,
      appBar: todoAppBar(
        context,
        elevation: 3,
        leading: todoLeadingIconButton(context, onPressed: () {
          Get.back();
        }, icon: Icons.arrow_back_ios),
        title: userProfile(),
      ),
      body: LayoutBuilder(builder: (context, constraints) {
        return Column(
          mainAxisSize: MainAxisSize.min,
          children: [
            Flexible(
              child: refreshContainer(
                  context: context,
                  onLoad: () {},
                  child: EmptyContainer(
                      icon: Icons.chat,
                      desc: 'try to send the first message',
                      what: '',
                      render: controller.chats.value.isNotEmpty,
                      child: Obx(() => ListView.separated(
                            padding: const EdgeInsets.symmetric(horizontal: 10),
                            itemCount: controller.chats.value.length,
                            reverse: true,
                            itemBuilder: (BuildContext context, int index) {
                              Chat chat = controller.chats.value[index];
                              return GestureDetector(
                                  onLongPressMoveUpdate: (details) async {
                                    final overlay = Overlay.of(context)
                                        .context
                                        .findRenderObject() as RenderBox;
                                    final menuItem = await showMenu<int>(
                                        context: context,
                                        items: [
                                          PopupMenuItem(
                                              value: 1, child: Text('copy'.tr)),
                                          PopupMenuItem(
                                              value: 2,
                                              child: Text('reply'.tr)),
                                        ],
                                        position: RelativeRect.fromSize(
                                            details.globalPosition &
                                                const Size(48.0, 48.0),
                                            overlay.size));
                                    // Check if menu item clicked
                                    if (context.mounted) {
                                      switch (menuItem) {
                                        case 1:
                                          ScaffoldMessenger.of(context)
                                              .showSnackBar(const SnackBar(
                                            content: Text('Copy clicked'),
                                            behavior: SnackBarBehavior.floating,
                                          ));
                                          break;
                                        case 2:
                                          show.value = true;
                                          replyID = chats[index].id;
                                          break;
                                        default:
                                      }
                                    }
                                  },
                                  child: ChatContainer(data: chat));
                            },
                            separatorBuilder:
                                (BuildContext context, int index) =>
                                    const SizedBox(height: 15),
                          )))),
            ),
            Container(
                decoration: BoxDecoration(
                  border: Border(
                    top: BorderSide(
                        color: ThemeProvider.contrastColor(context,
                            light: HexColor.fromInt(0xceced2),
                            dark: Colors.grey.withOpacity(0.8)),
                        width: 1),
                  ),
                ),
                child: Padding(
                    padding: const EdgeInsets.symmetric(vertical: 2),
                    child: Obx(() => TodoInput(
                        showChild: show.value,
                        controller: todoInputController,
                        child: Container(
                            decoration: BoxDecoration(
                                color: ThemeProvider.contrastColor(context,
                                    light: Colors.grey.withOpacity(0.2),
                                    dark: Colors.black.withOpacity(0.2)),
                                borderRadius: const BorderRadius.all(
                                    Radius.circular(8.0))),
                            child: Row(
                              mainAxisAlignment: MainAxisAlignment.spaceBetween,
                              children: [
                                const Padding(
                                  padding: EdgeInsets.only(left: 5),
                                  child: Text("回复  "),
                                ),
                                IconButton(
                                    onPressed: () {
                                      show.value = false;
                                    },
                                    icon: const Icon(
                                      Icons.close,
                                      size: 18,
                                    ))
                              ],
                            )),
                        onTap: (v) {
                          Chat msg = Chat(
                              from: Guard.user,
                              to: controller.user.id,
                              content: [v]);
                          if (show.value) {
                            msg.reply = replyID;
                            replyID = "0";
                          }
                          controller.sendMessage(msg);
                          if (show.value) {
                            show.value = false;
                          }
                          setState(() {});
                        })))),
            TodoInputView(
                controller: todoInputController,
                state: key,
                maxWidth: constraints.maxWidth),
          ],
        );
      }),
    );
  }

  Widget userProfile() {
    return InkWell(
      child: Row(
        children: [
          Padding(
            padding: const EdgeInsets.only(left: 0.0, right: 10.0),
            child: CircleAvatar(
              backgroundImage: TodoImage.userProfile(controller.user.id),
            ),
          ),
          Expanded(
            child: Column(
              crossAxisAlignment: CrossAxisAlignment.start,
              children: [
                Text(
                  controller.user.name,
                  style: const TextStyle(
                    fontWeight: FontWeight.bold,
                    fontSize: 14,
                  ),
                ),
                const SizedBox(height: 5),
                const Text(
                  "Online",
                  style: TextStyle(
                    fontWeight: FontWeight.w400,
                    fontSize: 11,
                  ),
                ),
              ],
            ),
          ),
        ],
      ),
      onTap: () {
        RouterProvider.viewUserProfile(controller.user.id);
      },
    );
  }
}
