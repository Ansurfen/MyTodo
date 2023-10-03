import 'package:flutter/animation.dart';

class TodoAnimateStyle {
  static fadeOutOpacity(Animation<double> parent) {
    return Tween<double>(begin: 0.0, end: 1.0).animate(CurvedAnimation(
        parent: parent,
        curve: const Interval((1 / 6) * 1, 1.0, curve: Curves.fastOutSlowIn)));
  }
}
