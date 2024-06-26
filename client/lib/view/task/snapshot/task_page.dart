import 'dart:async';
import 'package:flutter/material.dart';
import 'package:get/get.dart';
import 'package:my_todo/component/animate/fade_out_slow_in_container.dart';
import 'package:my_todo/component/container/empty_container.dart';
import 'package:my_todo/component/scaffold.dart';
import 'package:my_todo/component/title/title_view.dart';
import 'package:my_todo/mock/statistic.dart';
import 'package:my_todo/theme/animate.dart';
import 'package:my_todo/view/home/nav/component/app_bar.dart';
import 'package:my_todo/view/task/snapshot/task_card.dart';
import 'package:my_todo/view/task/snapshot/task_controller.dart';
import 'package:my_todo/view/task/snapshot/task_skeleton.dart';
import 'package:my_todo/view/task/component/statistic_table.dart';
import 'package:my_todo/utils/dialog.dart';
import 'package:my_todo/component/refresh.dart';
import 'package:easy_refresh/easy_refresh.dart';

class TaskPage extends StatefulWidget {
  const TaskPage({super.key});

  @override
  State<StatefulWidget> createState() => _TaskPageState();
}

class _TaskPageState extends State<TaskPage>
    with AutomaticKeepAliveClientMixin {
  TaskController controller = Get.find<TaskController>();
  EasyRefreshController easyRefreshController = EasyRefreshController();

  @override
  Widget build(BuildContext context) {
    super.build(context);
    return todoScaffold(
      context,
      appBar: AppBar(
        title: Padding(
          padding: const EdgeInsets.only(top: 5, left: 40),
          child: Text(
            "todo".tr,
            style: const TextStyle(fontSize: 20),
          ),
        ),
        actions: [
          notificationWidget(context),
          settingWidget(),
          multiWidget(context),
        ],
      ),
      body: FutureBuilder<bool>(
        future: controller.getData,
        builder: (context, snapshot) {
          if (snapshot.hasData) {
            Future.delayed(const Duration(milliseconds: 50), () {
              controller.animationController.forward();
            });
            return _taskList(context);
          }
          return const TaskSkeletonPage();
        },
      ),
    );
  }

  Widget _taskList(BuildContext context) {
    Size size = MediaQuery.sizeOf(context);
    var opacity =
        TodoAnimateStyle.fadeOutOpacity(controller.animationController);
    return refreshContainer(
        context: context,
        controller: easyRefreshController,
        onRefresh: () async {
          await Future.delayed(const Duration(seconds: 0), () {
            controller.refreshTask();
          });
        },
        onLoad: () async {
          await Future.delayed(const Duration(seconds: 0), () {
            controller.loadTask();
          });
        },
        child: Stack(
          children: [
            SingleChildScrollView(
                child: Column(children: [
              StatisticTable(
                data: statisticTableData,
                animation: opacity,
                animationController: controller.animationController,
              ),
              FadeAnimatedBuilder(
                animation: controller.animationController,
                opacity: opacity,
                child: TitleView(
                  onTap: () {
                    controller.showMask.value = true;
                  },
                  iconSize: 25,
                  iconColor: Theme.of(context).primaryColor,
                  icon: Icons.filter_alt,
                  titleTxt: 'task'.tr,
                  subTxt: 'filter'.tr,
                ),
              ),
              FadeAnimatedBuilder(
                  animation: controller.animationController,
                  opacity: opacity,
                  child: Obx(() => EmptyContainer(
                        height: MediaQuery.sizeOf(context).height / 2.5,
                        icon: Icons.rss_feed,
                        desc: "no_task".tr,
                        what: "what_is_task".tr,
                        render: controller.tasks.value.isNotEmpty,
                        alignment: Alignment.topCenter,
                        padding: EdgeInsets.only(top: size.height * 0.2),
                        onTap: () {
                          showTipDialog(context, content: "what_is_task".tr);
                        },
                        child: Padding(
                            padding: const EdgeInsets.symmetric(horizontal: 10),
                            child: ListView.builder(
                              physics: const NeverScrollableScrollPhysics(),
                              shrinkWrap: true,
                              itemCount: controller.tasks.value.length,
                              itemBuilder: (ctx, idx) {
                                var task = controller.tasks.value[
                                    controller.tasks.value.keys.elementAt(idx)];
                                if (task != null) {
                                  return TaskCard(model: task);
                                }
                                return null;
                              },
                            )),
                      ))),
            ])),
            Obx(() => controller.showMask.value
                ? Stack(
                    children: [
                      GestureDetector(
                        onTap: () {
                          controller.showMask.value = false;
                        },
                        child: Container(
                          height: MediaQuery.of(context).size.height,
                          color: Colors.black.withOpacity(0.4),
                        ),
                      ),
                      Container(
                        height: MediaQuery.of(context).size.height / 3,
                        decoration: const BoxDecoration(
                            color: Colors.white,
                            borderRadius: BorderRadius.only(
                                bottomLeft: Radius.circular(10),
                                bottomRight: Radius.circular(10))),
                        child: Column(
                          children: [Container()],
                        ),
                      ),
                    ],
                  )
                : Container()),
          ],
        ));
  }

  @override
  bool get wantKeepAlive => true;
}

List<String> taskTypes = ["已完成", "进行中", "未开始"];
// TODO：time picker, topic, task type
