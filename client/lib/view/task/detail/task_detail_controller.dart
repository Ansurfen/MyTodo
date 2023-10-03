import 'package:flutter/material.dart';
import 'package:get/get.dart';
import 'package:my_todo/api/task.dart';
import 'package:my_todo/model/dto/task.dart';
import 'package:my_todo/model/entity/task.dart';
import 'package:my_todo/router/provider.dart';
import 'package:my_todo/utils/dialog.dart';
import 'package:my_todo/utils/net.dart';
import 'package:my_todo/utils/picker.dart';
import 'package:web_socket_channel/web_socket_channel.dart';

class TaskInfoController extends GetxController {
  late final int id;
  InfoTaskDto? _task;
  Rx<List<TaskForm>> forms = Rx([]);
  late TaskForm selectedTask;
  late List<ValueChanged<dynamic>> onTaps;
  TextEditingController textAreaController = TextEditingController();
  Rx<String> qrCode = "".obs;
  WebSocketChannel? qrChannel;
  List<TFile> images = [];

  @override
  void onInit() {
    super.onInit();
    var data = Get.parameters;
    id = int.parse(data['id']!);
    onTaps = [
      (v) {
        commitTask(CommitTaskRequest(id, TaskCondType.hand.index, ""))
            .then((value) => null);
      },
      (v) {},
      (v) {
        RouterProvider.viewMapLocate()?.then((res) {
          commitTask(CommitTaskRequest(id, TaskCondType.locale.index, res))
              .then((res) {
            for (var form in forms.value) {
              if (form.type == TaskCondType.locale) {
                form.param.value = res.param;
                break;
              }
            }
          });
        });
      },
      (v) {
        commitTask(CommitTaskRequest(id, TaskCondType.file.index, "",
                files: (v as List).map((e) => e as TFile).toList()))
            .then((res) {});
      },
      (v) {
        commitTask(CommitTaskRequest(id, TaskCondType.image.index, "",
                images: (v as List).map((e) => e as TFile).toList()))
            .then((res) {})
            .onError((error, stackTrace) {
          showError(error.toString());
        });
      },
      (v) {
        commitTask(CommitTaskRequest(id, TaskCondType.content.index, v))
            .then((res) {})
            .onError((error, stackTrace) {
          showError(error.toString());
        });
      },
      (v) {}
    ];
    Future.delayed(Duration.zero, () {
      infoTask(InfoTaskRequest(id)).then((res) {
        _task = res.task;
        for (var cond in res.task.conds) {
          var opt = TaskForm.option(id)[cond.type];
          if (cond.type == TaskCondType.locale.index) {
            var res = cond.gotParams[0].split(",");
            if (res.length == 3) {
              opt.param.value = res[2];
              opt.isCompleted = true;
            }
          } else {
            if (cond.gotParams.isNotEmpty) {
              opt.param.value = cond.gotParams[0];
              opt.isCompleted = true;
            }
          }
          opt.wantCond = cond.wantParams;
          forms.value.add(opt);
        }
        if (forms.value.isNotEmpty) {
          forms.value[0].selected = true;
          selectedTask = forms.value[0];
        }
      }).onError((error, stackTrace) {
        showError(error.toString());
      });
    });
  }

  @override
  void dispose() {
    if (qrChannel != null) {
      qrChannel!.sink.close();
    }
    super.dispose();
  }

  void qrListen(String key) {
    qrChannel = WS.listen("/event/qr", onInit: key, callback: (v) {
      qrCode.value = v;
    });
  }

  InfoTaskDto get task {
    return _task ?? InfoTaskDto("", "", DateTime.now(), DateTime.now(), []);
  }

  String get condDesc {
    for (var form in forms.value) {
      if (form.selected) {
        return form.desc;
      }
    }
    return "???";
  }

  String get commitText {
    for (var form in forms.value) {
      if (form.selected) {
        return form.isCompleted ? "重新编辑" : "提交任务";
      }
    }
    return "???";
  }

  int get completedNumber {
    int count = 0;
    for (var form in forms.value) {
      if (form.isCompleted) {
        count++;
      }
    }
    return count;
  }

  ValueChanged<dynamic> get commitOnTap {
    for (var form in forms.value) {
      if (form.selected) {
        return onTaps[form.type.index];
      }
    }
    return (v) {};
  }

  void commit() {
    switch (selectedTask.type) {
      case TaskCondType.hand:
      case TaskCondType.timer:
      case TaskCondType.locale:
      case TaskCondType.file:
      case TaskCondType.content:
      case TaskCondType.image:
        break;
      case TaskCondType.qr:
    }
  }
}

class TaskForm {
  bool isCompleted = false;
  bool selected = false;
  bool committed = false;
  late String text;
  late String desc;
  TaskCondType type;
  List<String>? wantCond;
  Rx<String> param = "".obs;
  ValueChanged<dynamic>? onTap;

  TaskForm(this.type, this.text, this.desc, {this.onTap});

  static List<TaskForm> option(int id) {
    var list = [
      TaskForm(TaskCondType.hand, "手动签到", "手动完成签到", onTap: (v) {
        commitTask(CommitTaskRequest(id, TaskCondType.hand.index, ""))
            .then((value) => null);
      })
        ..selected = true,
      TaskForm(TaskCondType.timer, "定时签到", ""),
      TaskForm(TaskCondType.locale, "位置签到", "根据指定位置签到", onTap: (v) {
        RouterProvider.viewMapLocate()?.then((res) {
          commitTask(CommitTaskRequest(id, TaskCondType.locale.index, res))
              .then((res) {
            // for (var form in forms.value) {
            //   if (form.type == TaskCondType.locale) {
            //     form.param.value = res.param;
            //     break;
            //   }
            // }
          });
        });
      }),
      TaskForm(TaskCondType.file, "文件上传", "选择文件上传"),
      TaskForm(TaskCondType.image, "图片上传", "选择图片上传"),
      TaskForm(TaskCondType.content, "文字内容", "选择文字内容"),
      TaskForm(TaskCondType.qr, "扫码签到", "扫描发布者二维码即可完成签到"),
    ];
    return list;
  }
}
