class BaseResponse {
  final int? code;
  final String? msg;

  BaseResponse(Map<String, dynamic> json)
      : code = json['code'],
        msg = json['msg'];

  @override
  String toString() {
    return "code: $code, msg: $msg";
  }
}
