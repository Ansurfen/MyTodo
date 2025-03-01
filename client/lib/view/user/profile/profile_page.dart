// Copyright 2025 The MyTodo Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.
import 'package:flutter/material.dart';
import 'package:get/get.dart';
import 'package:my_todo/component/icon.dart';
import 'package:my_todo/component/image.dart';
import 'package:my_todo/component/scaffold.dart';
import 'package:my_todo/utils/image.dart';
import 'package:my_todo/view/home/nav/component/app_bar.dart';
import 'package:my_todo/view/user/profile/profile_controller.dart';

class ProfilePage extends StatefulWidget {
  const ProfilePage({super.key});

  @override
  State<ProfilePage> createState() => _ProfilePageState();
}

class _ProfilePageState extends State<ProfilePage> {
  ProfileController controller = Get.find<ProfileController>();

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      appBar: todoAppBar(
        context,
        leading: todoLeadingIconButton(
          context,
          onPressed: Get.back,
          icon: Icons.arrow_back_ios,
        ),
        elevation: 5,
        actions: [
          notificationWidget(context),
          const SizedBox(width: 30),
          settingWidget(),
          const SizedBox(width: 20),
          multiWidget(context),
          const SizedBox(width: 10),
        ],
      ),
      backgroundColor: Theme.of(context).colorScheme.primary,
      body: SingleChildScrollView(
        padding: const EdgeInsets.symmetric(horizontal: 10),
        child: SizedBox(
          width: MediaQuery.of(context).size.width,
          child: Column(
            mainAxisAlignment: MainAxisAlignment.center,
            crossAxisAlignment: CrossAxisAlignment.center,
            children: [
              const SizedBox(height: 60),
              CircleAvatar(
                backgroundImage: TodoImage.userProfile(controller.id),
                radius: 50,
              ),
              const SizedBox(height: 10),
              Obx(
                () => Text(
                  controller.user.value.name,
                  style: const TextStyle(
                    fontWeight: FontWeight.bold,
                    fontSize: 22,
                  ),
                ),
              ),
              const SizedBox(height: 3),
              Text(controller.user.value.description, style: TextStyle()),
              const SizedBox(height: 20),
              Row(
                mainAxisSize: MainAxisSize.min,
                children: [
                  TextButton(
                    style: ButtonStyle(
                      backgroundColor: WidgetStateProperty.all(
                        Theme.of(context).primaryColor,
                      ),
                    ),
                    onPressed: () {
                      // RouterProvider.viewChatConversation(
                      //     controller.user.value);
                    },
                    child: const Icon(Icons.message, color: Colors.white),
                  ),
                  const SizedBox(width: 10),
                  TextButton(
                    style: ButtonStyle(
                      backgroundColor: WidgetStateProperty.all(
                        Theme.of(context).primaryColor,
                      ),
                    ),
                    onPressed: () {},
                    child: const Text(
                      "follow",
                      style: TextStyle(color: Colors.white),
                    ),
                  ),
                ],
              ),
              const SizedBox(height: 40),
              Padding(
                padding: const EdgeInsets.symmetric(horizontal: 50),
                child: Row(
                  mainAxisAlignment: MainAxisAlignment.spaceBetween,
                  children: [
                    _buildCategory(
                      "post_count".tr,
                      controller.user.value.postCount,
                    ),
                    _buildCategory(
                      "follower_count".tr,
                      controller.user.value.followerCount,
                    ),
                    _buildCategory(
                      "topic_count".tr,
                      controller.user.value.topicCount,
                    ),
                  ],
                ),
              ),
              const SizedBox(height: 20),
              GridView.builder(
                shrinkWrap: true,
                physics: const NeverScrollableScrollPhysics(),
                primary: false,
                padding: const EdgeInsets.all(5),
                itemCount: 15,
                gridDelegate: const SliverGridDelegateWithFixedCrossAxisCount(
                  crossAxisCount: 3,
                  childAspectRatio: 200 / 200,
                ),
                itemBuilder: (BuildContext context, int index) {
                  return Padding(
                    padding: const EdgeInsets.all(5.0),
                    child: ImageView.svg("assets/images/flutter.svg"),
                  );
                },
              ),
            ],
          ),
        ),
      ),
    );
  }

  Widget _buildCategory(String title, int count) {
    return Column(
      children: [
        Text(
          count.toString(),
          style: const TextStyle(fontWeight: FontWeight.bold, fontSize: 22),
        ),
        const SizedBox(height: 4),
        Text(title, style: const TextStyle()),
      ],
    );
  }
}
