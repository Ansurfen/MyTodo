import 'package:my_todo/utils/web_sanbox/annotation.dart';

@WebSandInterface()
class TLocation {

  @DartMethod("abc")
  static String abc() {
    return "abc called";
  }

  int test() {
    return 0;
  }
}
