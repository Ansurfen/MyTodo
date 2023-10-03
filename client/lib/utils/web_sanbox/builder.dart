import 'dart:async';

import 'package:build/build.dart';
import 'package:my_todo/utils/web_sanbox/gen.dart';
import 'package:source_gen/source_gen.dart';

class JSBuilder extends Builder {
  JSBuilder();

  @override
  Future build(BuildStep buildStep) {
    // TODO: implement build
    throw UnimplementedError();
  }

  @override
  final buildExtensions = const {
    ".dart": [".g.js"]
  };
}

Builder dartBindingBuilder(BuilderOptions options) =>
    LibraryBuilder(TestGenerator());
