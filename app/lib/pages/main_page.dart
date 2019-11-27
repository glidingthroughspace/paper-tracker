import 'package:flutter/material.dart';
import 'package:paper_tracker/model/room.dart';
import 'package:paper_tracker/model/tracker.dart';
import 'package:paper_tracker/model/workflow.dart';
import 'package:paper_tracker/widgets/room_list.dart';
import 'package:paper_tracker/widgets/tracker_list.dart';

class MainPage extends StatelessWidget {
  static const Route = "/main";

  @override
  Widget build(BuildContext context) {
    return DefaultTabController(
        length: 3,
        child: Scaffold(
          backgroundColor: Theme.of(context).backgroundColor,
          appBar: AppBar(
            backgroundColor: Theme.of(context).backgroundColor,
            title: Text("Paper Tracker"),
            bottom: TabBar(
              tabs: [
                Tab(
                  icon: Icon(Workflow.IconData),
                  text: "Workflows",
                ),
                Tab(
                  icon: Icon(Room.IconData),
                  text: "Rooms",
                ),
                Tab(
                  icon: Icon(Tracker.IconData),
                  text: "Tracker",
                ),
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
        ));
  }
}
