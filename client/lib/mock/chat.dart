import 'package:my_todo/mock/provider.dart';

List friends = List.generate(
    13,
    (index) => {
          "name": Mock.username(),
          "dp": "assets/images/flutter.svg",
          "status": "Anything could be here",
          "isAccept": Mock.boolean(),
        });
