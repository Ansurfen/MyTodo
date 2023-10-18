import 'package:flutter/material.dart';
import 'package:flutter_easyloading/flutter_easyloading.dart';
import 'package:get/get.dart';
import 'package:my_todo/i18n/index.dart';
import 'package:my_todo/router/provider.dart';
import 'package:my_todo/utils/db.dart';
import 'package:my_todo/utils/guard.dart';
import 'package:my_todo/utils/notification.dart';
import 'package:my_todo/theme/provider.dart';
import 'package:oktoast/oktoast.dart';

Future main() async {
  runApp(await myTodo());
}

Future<Widget> myTodo() async {
  WidgetsFlutterBinding.ensureInitialized();
  await Guard.init();
  await Future.wait([
    DBProvider.init(),
    NotifyProvider.init(),
    Future(() => ThemeProvider.init()),
  ]);

  return OKToast(
      child: GetMaterialApp(
    title: "My Todo",
    debugShowCheckedModeBanner: false,
    builder: EasyLoading.init(),
    translations: I18N(),
    locale: Guard.initLanguage(),
    fallbackLocale: const Locale('en', 'US'),
    theme: TodoThemeData.lightTheme(),
    darkTheme: TodoThemeData.darkTheme(),
    themeMode: ThemeMode.light,
    initialRoute: RouterProvider.initialRoute(),
    getPages: RouterProvider.pages,
    defaultTransition: Transition.fade,
    unknownRoute: RouterProvider.notFoundPage,
  ));
}
