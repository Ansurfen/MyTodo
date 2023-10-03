import 'package:get/get.dart';

int dateTimeString2Int(String str) {
  return DateTime.parse(str).microsecondsSinceEpoch;
}

DateTime string2DateTime(String str) {
  return DateTime.parse(str);
}

String formatTimeDifference(DateTime dateTime) {
  final now = DateTime.now();
  final difference = now.difference(dateTime);

  if (difference.inDays > 0) {
    return '${difference.inDays} ${'day_ago'.tr}';
  } else if (difference.inHours > 0) {
    return '${difference.inHours} ${'hour_ago'.tr}';
  } else if (difference.inMinutes > 0) {
    return '${difference.inMinutes} ${'min_ago'.tr}';
  } else {
    return '${difference.inSeconds} ${'sec_ago'.tr}';
  }
}
