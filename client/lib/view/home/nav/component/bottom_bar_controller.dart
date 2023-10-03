import 'package:flutter/animation.dart';
import 'package:get/get.dart';

class BottomBarController extends GetxController
    with GetTickerProviderStateMixin {
  late AnimationController animationController;

  @override
  void onInit() {
    super.onInit();
    animationController =
        AnimationController(duration: const Duration(seconds: 1), vsync: this);
    animationController.forward();
  }

  @override
  void dispose() {
    animationController.dispose();
    super.dispose();
  }
}
