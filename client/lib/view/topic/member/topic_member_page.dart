import 'package:flutter/cupertino.dart';
import 'package:flutter/material.dart';
import 'package:get/get.dart';
import 'package:my_todo/component/icon.dart';
import 'package:my_todo/component/image.dart';
import 'package:my_todo/component/refresh.dart';
import 'package:my_todo/component/scaffold.dart';
import 'package:my_todo/mock/chat.dart';
import 'package:my_todo/mock/provider.dart';
import 'package:my_todo/model/entity/topic.dart';
import 'package:my_todo/router/provider.dart';
import 'package:my_todo/utils/dialog.dart';
import 'package:my_todo/view/home/nav/component/app_bar.dart';
import 'package:my_todo/view/topic/member/topic_member_controller.dart';

class TopicMemberPage extends StatefulWidget {
  const TopicMemberPage({super.key});

  @override
  State<TopicMemberPage> createState() => _TopicMemberPageState();
}

class _TopicMemberPageState extends State<TopicMemberPage> {
  TopicMemberController controller = Get.find<TopicMemberController>();

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      backgroundColor: Theme.of(context).colorScheme.primary,
      appBar: todoAppBar(
        context,
        leading: todoLeadingIconButton(context,
            onPressed: Get.back, icon: Icons.arrow_back_ios),
        title: Text("member".tr),
        elevation: 5,
        actions: [
          notificationWidget(context),
          const SizedBox(
            width: 30,
          ),
          settingWidget(),
          const SizedBox(
            width: 20,
          ),
          const IconButton(
              onPressed: RouterProvider.viewTopicInvite,
              icon: Icon(Icons.share)),
          const SizedBox(
            width: 10,
          )
        ],
      ),
      body: refreshContainer(
          context: context,
          onRefresh: () {},
          onLoad: () {},
          child: Obx(() => ListView.separated(
                padding: const EdgeInsets.all(10),
                separatorBuilder: (BuildContext context, int index) {
                  return Align(
                    alignment: Alignment.centerRight,
                    child: SizedBox(
                      height: 0.5,
                      width: MediaQuery.of(context).size.width / 1.3,
                      child: const Divider(),
                    ),
                  );
                },
                itemCount: controller.members.value.length,
                itemBuilder: (BuildContext context, int index) {
                  Map friend = friends[index];
                  TopicMember member = controller.members.value[index];
                  int perm = Mock.number(max: 1);
                  return Padding(
                    padding: const EdgeInsets.symmetric(horizontal: 8.0),
                    child: ListTile(
                      leading: CircleAvatar(
                        backgroundImage: TodoImage.userProfile(member.id),
                        radius: 25,
                      ),
                      contentPadding: const EdgeInsets.all(0),
                      title: Text(member.name),
                      subtitle: Text(friend['status']),
                      trailing: IconButton(
                        onPressed: () {
                          handleMember(perm);
                        },
                        icon: const Icon(Icons.more_vert),
                      ),
                      onTap: () {
                        RouterProvider.viewUserProfile(member.id);
                      },
                    ),
                  );
                },
              ))),
    );
  }

  void handleMember(int perm) {
    showCupertinoModalPopup(
        context: context,
        builder: (BuildContext context) => CupertinoActionSheet(
              message: Column(
                children: [
                  dialogAction(
                      icon: perm == 0 ? Icons.power : Icons.power_off,
                      text: perm == 0 ? "grant" : "un_grant",
                      onTap: () {}),
                  const SizedBox(height: 15),
                  dialogAction(icon: Icons.notifications, text: "notification"),
                  const SizedBox(height: 15),
                  const Divider(),
                  dialogAction(
                      icon: Icons.delete,
                      text: "delete".tr,
                      onTap: () {
                        Navigator.of(context).pop();
                        showAlert(context,
                            title: "移除成员",
                            content: "确认移除该成员",
                            onConfirm: () {

                            });
                      })
                ],
              ),
            ));
  }
}
