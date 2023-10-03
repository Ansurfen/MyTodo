import 'package:dio/dio.dart';
import 'package:flutter/services.dart';
import 'package:get/get.dart';

const Map<Type, String> exceptionFormat = {
  DioException: "exp.dio",
  MissingPluginException: "exp.mp",
  UnsupportedError: "exp.mp"
};

extension ExceptionI18NExtension on Exception {
  String get tr {
    String? key = exceptionFormat[runtimeType];
    if (key != null) {
      return key.tr;
    }
    return "exp.unknown".tr;
  }
}

extension ErrorI18NExtension on Error {
  String get tr {
    String? key = exceptionFormat[runtimeType];
    if (key != null) {
      return key.tr;
    }
    return "exp.unknown".tr;
  }
}