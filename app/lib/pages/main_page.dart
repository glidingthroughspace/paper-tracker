import 'package:flutter/material.dart';
import 'package:paper_tracker/widgets/room_list.dart';
import 'package:paper_tracker/widgets/tracker_list.dart';

class MainPage extends StatelessWidget {
  static const Route = "/main";

  @override
  Widget build(BuildContext context) {
    return DefaultTabController(
      length: 3,
      child: Scaffold(
        appBar: AppBar(
          title: Text("Paper Tracker"),
          bottom: TabBar(
            tabs: [
              Text("Workflows"),
              Text("Rooms"),
              Text("Tracker"),
            ],
          ),
        ),
        body: TabBarView(
          children: [
            Text("Workflows"),
            RoomList(),
            TrackerList(),
          ],
        ),
      )
    );
  }
}
