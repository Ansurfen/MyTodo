// Copyright 2025 The MyTodo Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.
import 'dart:io';
import 'dart:math';

import 'package:expansion_tile_card/expansion_tile_card.dart';
import 'package:flutter/material.dart';
import 'package:get/get.dart';
import 'package:intl/intl.dart';
import 'package:my_todo/mock/provider.dart';
import 'package:my_todo/model/dto/task.dart';
import 'package:my_todo/model/dto/topic.dart';
import 'package:my_todo/model/entity/task.dart';
import 'package:my_todo/router/provider.dart';
import 'package:my_todo/router/topic.dart';
import 'package:my_todo/theme/color.dart';
import 'package:my_todo/theme/provider.dart';
import 'package:my_todo/utils/clipboard.dart';
import 'package:my_todo/utils/share.dart';

class TaskCardModel {
  int? id;
  String name;
  String topic;
  String desc;
  DateTime startAt;
  List<int> cond;

  TaskCardModel(
    this.name,
    this.topic,
    this.desc,
    this.startAt,
    this.cond, {
    this.id,
  });
}

class TaskCardOld extends StatelessWidget {
  final GetTaskDto model;

  const TaskCardOld({super.key, required this.model});

  @override
  Widget build(BuildContext context) {
    List<Widget> icons = [];
    for (var i = 0; i < model.conds.length; i++) {
      if (model.conds[i] == TaskCondType.qr.index) {
        icons.add(
          Icon(Icons.crop_free, color: Theme.of(context).colorScheme.onPrimary),
        );
      } else if (model.conds[i] == TaskCondType.hand.index) {
        icons.add(
          Icon(Icons.handshake, color: Theme.of(context).colorScheme.onPrimary),
        );
      } else if (model.conds[i] == TaskCondType.locale.index) {
        icons.add(
          Icon(
            Icons.location_on,
            color: Theme.of(context).colorScheme.onPrimary,
          ),
        );
      }
      if (i + 1 != model.conds.length) {
        icons.add(const SizedBox(width: 5));
      }
    }
    return GestureDetector(
      onTap: () {
        RouterProvider.viewTaskDetail(model.id, []);
      },
      child: Card(
        color: ThemeProvider.contrastColor(
          context,
          light: HexColor.fromInt(0xfafafa),
          dark: HexColor.fromInt(0x1c1c1e),
        ),
        shadowColor: Colors.black,
        elevation: 2,
        borderOnForeground: false,
        child: Container(
          height: 120,
          padding: const EdgeInsets.all(10),
          child: Column(
            mainAxisAlignment: MainAxisAlignment.spaceBetween,
            children: [
              Row(
                mainAxisAlignment: MainAxisAlignment.spaceBetween,
                children: [
                  Column(
                    crossAxisAlignment: CrossAxisAlignment.start,
                    children: [
                      Text(model.name, style: const TextStyle(fontSize: 18)),
                      Text(model.topic, style: const TextStyle(fontSize: 12)),
                    ],
                  ),
                  Text(
                    DateFormat("yyyy/MM/dd HH:mm:ss").format(model.departure),
                  ),
                ],
              ),
              Row(children: [...icons]),
            ],
          ),
        ),
      ),
    );
  }
}

class TopicItem extends StatefulWidget {
  final String dp;
  final String name;
  final String time;
  final String msg;
  final int counter;
  final GetTopicDto model;

  const TopicItem({
    super.key,
    required this.dp,
    required this.name,
    required this.time,
    required this.msg,
    required this.counter,
    required this.model,
  });

  @override
  State<TopicItem> createState() => _TopicItemState();
}

class _TopicItemState extends State<TopicItem> {
  @override
  Widget build(BuildContext context) {
    return Padding(
      padding: const EdgeInsets.symmetric(horizontal: 8.0),
      child: ListTile(
        contentPadding: const EdgeInsets.all(0),
        leading: const CircleAvatar(
          // backgroundImage: ,
          radius: 25,
        ),
        title: Text(
          widget.name,
          maxLines: 1,
          style: const TextStyle(fontWeight: FontWeight.bold),
        ),
        subtitle: Text(
          widget.msg,
          overflow: TextOverflow.ellipsis,
          maxLines: 2,
        ),
        trailing: Column(
          crossAxisAlignment: CrossAxisAlignment.end,
          children: [
            const SizedBox(height: 10),
            Text(
              widget.time,
              style: const TextStyle(fontWeight: FontWeight.w300, fontSize: 11),
            ),
            const SizedBox(height: 5),
            widget.counter == 0
                ? const SizedBox()
                : Container(
                  padding: const EdgeInsets.all(1),
                  decoration: BoxDecoration(
                    borderRadius: BorderRadius.circular(6),
                  ),
                  constraints: const BoxConstraints(
                    minWidth: 11,
                    minHeight: 11,
                  ),
                  child: Padding(
                    padding: const EdgeInsets.only(top: 1, left: 5, right: 5),
                    child: Text(
                      "${widget.counter}",
                      style: const TextStyle(color: Colors.white, fontSize: 10),
                      textAlign: TextAlign.center,
                    ),
                  ),
                ),
          ],
        ),
        onTap: () {
          RouterProvider.to(
            TopicRouter.detail,
            query: "/?id=${widget.model.id}",
            arguments: widget.model,
          );
        },
      ),
    );
  }
}

class TaskCard extends StatefulWidget {
  const TaskCard({
    super.key,
    required this.title,
    required this.msg,
    required this.model,
  });

  final GetTopicDto model;
  final String title;
  final String msg;

  @override
  State<TaskCard> createState() => _TaskCardState();
}

class _TaskCardState extends State<TaskCard>
    with AutomaticKeepAliveClientMixin {
  List<Color> colors = [
    const Color(0xff8D7AEE),
    const Color(0xffF468B7),
    const Color(0xffFEC85C),
    const Color(0xff5FD0D3),
    const Color(0xffBFACAA),
  ];
  Random r = Random();

  @override
  Widget build(BuildContext context) {
    super.build(context);
    final ButtonStyle flatButtonStyle = TextButton.styleFrom(
      shape: const RoundedRectangleBorder(
        borderRadius: BorderRadius.all(Radius.circular(4.0)),
      ),
    );
    bool isLight = Theme.of(context).brightness == Brightness.light;
    var key = widget.key as ValueKey<ExpansionTileCardState>;

    var conds = ConditionItem.randomList(
      Mock.number(min: 1, max: ConditionType.values.length),
    );
    return ExpansionTileCard(
      key: key,
      elevation: 0,
      baseColor: isLight ? Colors.grey.shade50 : HexColor.fromInt(0x1c1c1e),
      expandedColor: isLight ? Colors.grey.shade50 : HexColor.fromInt(0x1c1c1e),
      leading: CircleAvatar(
        backgroundColor: colors[r.nextInt(colors.length)],
        child: Text(
          widget.title[0],
          style: const TextStyle(
            color: Colors.white,
            fontWeight: FontWeight.bold,
            fontSize: 16,
          ),
        ),
      ),
      title: Text(
        widget.title,
        style: TextStyle(
          fontWeight: FontWeight.bold,
          fontSize: 16,
          color: Theme.of(context).colorScheme.onPrimary,
        ),
      ),
      subtitle: Text(
        widget.msg,
        style: TextStyle(
          color: isLight ? Colors.black26 : Colors.grey,
          fontWeight: FontWeight.bold,
          fontSize: 12,
        ),
      ),
      children: [
        const Divider(thickness: 1.0, height: 1.0),
        Align(
          alignment: Alignment.centerLeft,
          child: Padding(
            padding: const EdgeInsets.symmetric(
              horizontal: 16.0,
              vertical: 8.0,
            ),
            child: ListView.builder(
              shrinkWrap: true,
              physics: const NeverScrollableScrollPhysics(),
              itemCount: conds.length,
              itemBuilder: (context, idx) {
                return _buildCondition(context, conds[idx]);
              },
            ),
          ),
        ),
        ButtonBar(
          alignment: MainAxisAlignment.spaceAround,
          buttonHeight: 52.0,
          buttonMinWidth: 90.0,
          children: [
            // TODO
            IconButton(
              onPressed: () {
                RouterProvider.viewTaskDetail(1, conds);
              },
              icon: Icon(Icons.article, color: Theme.of(context).primaryColor),
            ),
            IconButton(
              onPressed: () {
                RouterProvider.viewTopicMember(widget.model.id);
              },
              icon: Icon(Icons.group, color: Theme.of(context).primaryColor),
            ),
            IconButton(
              onPressed: () async {
                TodoShare.shareUri(
                  context,
                  Uri.parse(widget.model.inviteCode),
                ).then(
                  (value) => Get.snackbar(
                    "Clipboard",
                    "Topic's invite code is copied on clipboard.",
                    backgroundColor: Theme.of(context).colorScheme.primary,
                  ),
                );
                await TodoClipboard.set(widget.model.inviteCode);
              },
              icon: Icon(Icons.share, color: Theme.of(context).primaryColor),
            ),
          ],
        ),
      ],
    );
  }

  Widget _buildCondition(BuildContext context, ConditionItem item) {
    return InkWell(
      onTap: () {},
      child: ListTile(
        contentPadding: EdgeInsets.all(0),
        title: Column(
          crossAxisAlignment: CrossAxisAlignment.start,
          children: [
            Text(
              item.type.toString(),
              style: const TextStyle(
                color: Colors.black,
                fontSize: 18,
                fontWeight: FontWeight.bold,
              ),
            ),
            Text(
              item.subtitle,
              style: const TextStyle(color: Colors.black54, fontSize: 16),
            ),
          ],
        ),
        trailing:
            item.finish
                ? Icon(Icons.check, color: Colors.greenAccent)
                : Icon(Icons.close, color: Colors.redAccent),
        leading: Icon(item.icon()),
      ),
    );
  }

  @override
  bool get wantKeepAlive => true;
}

class ConditionItem {
  String subtitle;
  bool finish;
  ConditionType type;

  ConditionItem({
    required this.finish,
    required this.subtitle,
    required this.type,
  });

  IconData icon() {
    switch (type) {
      case ConditionType.click:
        return Icons.ads_click;
      case ConditionType.file:
        return Icons.drive_folder_upload;
      case ConditionType.qr:
        return Icons.qr_code;
      case ConditionType.locale:
        return Icons.location_on;
      case ConditionType.text:
        return Icons.abc;
    }
  }

  static ConditionItem random() {
    return ConditionItem(
      finish: true,
      type:
          ConditionType.values[Mock.number(
            max: ConditionType.values.length - 1,
          )],
      subtitle: "",
    );
  }

  static List<ConditionItem> randomList(int count) {
    // 1. 打乱 ConditionType.values
    List<ConditionType> shuffledTypes = List.of(ConditionType.values)
      ..shuffle();

    // 2. 选取 count 个唯一的 ConditionType
    List<ConditionType> selectedTypes = shuffledTypes.take(count).toList();

    // 3. 生成 ConditionItem 列表
    return selectedTypes.map((type) {
      return ConditionItem(
        finish: Mock.boolean(),
        type: type,
        subtitle: "Random subtitle",
      );
    }).toList();
  }
}

enum ConditionType {
  click,
  qr,
  locale,
  text,
  file;

  @override
  String toString() {
    switch (this) {
      case ConditionType.click:
        return "condition_click".tr;
      case ConditionType.qr:
        return "condition_qr".tr;
      case ConditionType.locale:
        return "condition_locale".tr;
      case ConditionType.file:
        return "condition_file".tr;
      case ConditionType.text:
        return "condition_text".tr;
    }
  }

  static ConditionType fromString(String value) {
    switch (value) {
      case 'locale':
        return ConditionType.locale;
      case 'click':
        return ConditionType.click;
      case 'qr':
        return ConditionType.qr;
      case 'text':
        return ConditionType.text;
      case 'file':
        return ConditionType.file;
      default:
        throw ArgumentError('Invalid enum value: $value');
    }
  }
}
