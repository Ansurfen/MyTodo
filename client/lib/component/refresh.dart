import 'dart:async';

import 'package:easy_refresh/easy_refresh.dart';
import 'package:flutter/material.dart';
import 'package:get/get.dart';

class TodoRefreshFooter extends ClassicFooter {
  BuildContext context;

  TodoRefreshFooter(this.context);

  @override
  double? get infiniteOffset => null;

  @override
  String? get dragText => "refresh.drag".tr;

  @override
  String? get armedText => "refresh.armed".tr;

  @override
  String? get readyText => "refresh.ready".tr;

  @override
  String? get processingText => "refresh.processing".tr;

  @override
  String? get processedText => "refresh.processed".tr;

  @override
  String? get noMoreText => "refresh.no_more".tr;

  @override
  String? get failedText => "refresh.failed".tr;

  @override
  String? get messageText => "refresh.message".tr;

  @override
  IconThemeData? get iconTheme =>
      IconThemeData(color: Theme.of(context).primaryColor);
}

class TodoRefreshHeader extends ClassicHeader {
  BuildContext context;

  TodoRefreshHeader(this.context);

  @override
  String? get dragText => "refresh.drag".tr;

  @override
  String? get armedText => "refresh.armed".tr;

  @override
  String? get readyText => "refresh.ready".tr;

  @override
  String? get processingText => "refresh.processing".tr;

  @override
  String? get processedText => "refresh.processed".tr;

  @override
  String? get noMoreText => "refresh.no_more".tr;

  @override
  String? get failedText => "refresh.failed".tr;

  @override
  String? get messageText => "refresh.message".tr;

  @override
  IconThemeData? get iconTheme =>
      IconThemeData(color: Theme.of(context).primaryColor);
}

Widget refreshContainer(
    {required BuildContext context,
    required Widget child,
    FutureOr Function()? onRefresh,
    FutureOr Function()? onLoad,
    EasyRefreshController? controller}) {
  return EasyRefresh(
    controller: controller,
    header: TodoRefreshHeader(context),
    footer: TodoRefreshFooter(context),
    onRefresh: onRefresh,
    onLoad: onLoad,
    child: child,
  );
}
