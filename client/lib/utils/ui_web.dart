import 'package:flutter/cupertino.dart';
import 'dart:ui' as ui;

Widget htmlElementView(String viewType) {
  return HtmlElementView(
    viewType: viewType,
  );
}

// ignore: camel_case_types
class platformViewRegistry {
  static registerViewFactory(String viewId, dynamic cb, {bool isVisible = true}) {
    // ignore:undefined_prefixed_name
    ui.platformViewRegistry.registerViewFactory(viewId, cb);
  }
}
