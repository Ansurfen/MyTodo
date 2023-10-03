import 'package:build/build.dart';
import 'package:source_gen/source_gen.dart';

import 'formdata.dart';

Builder formDataBuilder(BuilderOptions options) =>
    LibraryBuilder(FormDataGenerator());
