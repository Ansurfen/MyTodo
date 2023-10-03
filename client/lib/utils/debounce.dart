Future<T> Function() doOnce<T>(Future<T> Function() callback) {
  bool once = false;
  return () {
    if (!once) {
      once = true;
      return callback();
    }
    return Future.value();
  };
}
