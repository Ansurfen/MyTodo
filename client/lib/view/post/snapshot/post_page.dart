import 'dart:async';

import 'package:flutter/material.dart';
import 'package:get/get.dart';
import 'package:my_todo/api/post.dart';
import 'package:my_todo/component/container/empty_container.dart';
import 'package:my_todo/hook/post.dart';
import 'package:my_todo/model/entity/post.dart';
import 'package:my_todo/theme/provider.dart';
import 'package:my_todo/view/home/nav/component/app_bar.dart';
import 'package:my_todo/theme/color.dart';
import 'package:my_todo/utils/dialog.dart';
import 'package:my_todo/utils/guard.dart';
import 'package:my_todo/component/refresh.dart';
import 'package:my_todo/view/post/snapshot/post_card.dart';
import 'package:my_todo/model/vo/post.dart';

class PostPage extends StatefulWidget {
  const PostPage({super.key});

  @override
  State<StatefulWidget> createState() => _PostPageState();
}

class _PostPageState extends State<PostPage>
    with AutomaticKeepAliveClientMixin, TickerProviderStateMixin {
  List<Widget> views = [];
  late StreamSubscription<Post> _uploadPost;
  late TabController tabController;
  @override
  void initState() {
    super.initState();
    tabController = TabController(length: 3, vsync: this, initialIndex: 1);
    Future.delayed(Duration.zero, () {
      getPost(GetPostRequest(1, 10)).then((res) {
        for (var post in res.data) {
          views.add(PostCard(model: GetPostVo.fromDto(post)));
          views.add(_postCardSpace());
        }
        setState(() {});
      }).catchError((err) {
        showError(err);
      });
    });
    _uploadPost = PostHook.subscribeSnapshot(onData: (post) {
      setState(() {
        views.add(PostCard(
            model: GetPostVo(0, 0, Guard.userName(), true, DateTime.timestamp(),
                post.content, [], 0, 0)));
        views.add(_postCardSpace());
      });
    });
  }

  Widget _postCardSpace() {
    return Container(
      height: 10,
      color: ThemeProvider.contrastColor(context,
          light: Colors.grey.withOpacity(0.2),
          dark: HexColor.fromInt(0x1c1c1e)),
    );
  }

  @override
  void dispose() {
    _uploadPost.cancel();
    super.dispose();
  }

  @override
  Widget build(BuildContext context) {
    super.build(context);
    Size size = MediaQuery.sizeOf(context);
    return Scaffold(
        appBar: AppBar(
            actions: [
              settingWidget(),
              const SizedBox(
                width: 10,
              ),
              multiWidget(context)
            ],
            centerTitle: true,
            backgroundColor: Theme.of(context).colorScheme.primary,
            title: TabBar(
              controller: tabController,
              isScrollable: true,
              labelColor: Theme.of(context).colorScheme.onPrimary,
              unselectedLabelColor: Theme.of(context).colorScheme.onTertiary,
              indicatorSize: TabBarIndicatorSize.label,
              indicator: UnderlineTabIndicator(
                  borderRadius: const BorderRadius.all(Radius.circular(10)),
                  borderSide: BorderSide(
                    width: 1,
                    color: Theme.of(context).colorScheme.onPrimary,
                  )),
              tabs: [
                Tab(
                  text: "post.me".tr,
                ),
                Tab(
                  text: "post.find".tr,
                ),
                Tab(
                  text: "post.friend".tr,
                )
              ],
            )),
        backgroundColor: Theme.of(context).colorScheme.primary,
        body: TabBarView(
          controller: tabController,
          children: [_me(), _find(), _friend()],
        ));
  }

  Widget _me() {
    return refreshContainer(
        context: context,
        onLoad: () {},
        onRefresh: () {
          views.clear();
          getPost(GetPostRequest(1, 10)).then((res) {
            for (var post in res.data) {
              views.add(PostCard(model: GetPostVo.fromDto(post)));
              views.add(_postCardSpace());
            }
            setState(() {});
          }).catchError((err) {
            showError(err);
          });
        },
        child: EmptyContainer(
            icon: Icons.rss_feed,
            desc: "not post, clicks + button to create on bottom bar",
            what: "what is post?",
            render: views.isNotEmpty,
            alignment: Alignment.center,
            padding:
                EdgeInsets.only(top: MediaQuery.sizeOf(context).height * 0.35),
            onTap: () {
              showTipDialog(context, content: "what_is_post".tr);
            },
            child: ListView.builder(
              itemCount: views.length,
              itemBuilder: (BuildContext context, int index) {
                return views[index];
              },
            )));
  }

  Widget _find() {
    return refreshContainer(
        context: context,
        onLoad: () {},
        onRefresh: () {},
        child: EmptyContainer(
            icon: Icons.rss_feed,
            desc: "not post, clicks + button to create on bottom bar",
            what: "what is post?",
            render: views.isNotEmpty,
            alignment: Alignment.center,
            padding:
                EdgeInsets.only(top: MediaQuery.sizeOf(context).height * 0.35),
            onTap: () {
              showTipDialog(context, content: "what_is_post".tr);
            },
            child: ListView.builder(
              itemCount: views.length,
              itemBuilder: (BuildContext context, int index) {
                return views[index];
              },
            )));
  }

  Widget _friend() {
    return refreshContainer(
        context: context,
        onLoad: () {},
        onRefresh: () {},
        child: EmptyContainer(
            icon: Icons.rss_feed,
            desc: "not post, clicks + button to create on bottom bar",
            what: "what is post?",
            render: views.isNotEmpty,
            alignment: Alignment.center,
            padding:
                EdgeInsets.only(top: MediaQuery.sizeOf(context).height * 0.35),
            onTap: () {
              showTipDialog(context, content: "what_is_post".tr);
            },
            child: ListView.builder(
              itemCount: views.length,
              itemBuilder: (BuildContext context, int index) {
                return views[index];
              },
            )));
  }

  @override
  bool get wantKeepAlive => true;
}
