import 'package:flutter/animation.dart';
import 'package:get/get.dart';
import 'package:my_todo/api/notify.dart';
import 'package:my_todo/model/entity/notify.dart';

class NotificationController extends GetxController
    with GetTickerProviderStateMixin {
  late final AnimationController animationController;
  Rx<List<Notify>> notifications = Rx([]);

  @override
  void onInit() {
    super.onInit();
    animationController = AnimationController(
        duration: const Duration(milliseconds: 600), vsync: this);
    Future.delayed(const Duration(milliseconds: 100), () {
      animationController.forward();
      notifyAll().then((res) => notifications.value = res.notifications);
    });
  }
}
