import 'package:flutter/cupertino.dart';
import 'package:get/get.dart';

class ForgetController extends GetxController {
  var forgetPasswordFormKey = GlobalKey<FormState>();

  late String email;

  void sendCaptcha() {
    if ((forgetPasswordFormKey.currentState as FormState).validate()) {
      (forgetPasswordFormKey.currentState as FormState).save();
      // ScaffoldMessenger.of(context).showSnackBar(
      //   const SnackBar(content: Text('Submitting data..')),
      // );
      print(email);
    }
  }
}
