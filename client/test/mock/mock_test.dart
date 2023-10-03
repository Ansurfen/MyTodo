import 'package:flutter_test/flutter_test.dart';
import 'package:my_todo/mock/provider.dart';

void main() {
  print(Mock.dateTime());
  for (int i = 0; i < 10; i++) {
    test("random number test", () {
      expect(Mock.number(min: -10, max: 10), inInclusiveRange(-10, 10));
    });

    test("random location test", () {
      var coordinates = Mock.location();
      expect(coordinates.latitude, inInclusiveRange(-90, 90));
      expect(coordinates.longitude, inInclusiveRange(-180, 180));
    });
  }
}
