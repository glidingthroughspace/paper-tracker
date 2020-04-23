import 'package:flutter/material.dart';
import 'package:paper_tracker/model/room.dart';
import 'package:paper_tracker/model/tracker.dart';
import 'package:paper_tracker/model/workflow.dart';
import 'package:paper_tracker/pages/config_page.dart';
import 'package:paper_tracker/pages/tutorial_page.dart';
import 'package:paper_tracker/widgets/lists/room_list.dart';
import 'package:paper_tracker/widgets/lists/tracker_list.dart';
import 'package:paper_tracker/widgets/lists/workflow_exec_list.dart';
import 'package:paper_tracker/widgets/lists/workflow_template_list.dart';

class MainPage extends StatelessWidget {
  static const Route = "/main";

  @override
  Widget build(BuildContext context) {
    return DefaultTabController(
      length: 4,
      child: Scaffold(
        appBar: AppBar(
          title: Text("Paper Tracker"),
          actions: [
            IconButton(
              icon: Icon(Icons.settings),
              onPressed: () => Navigator.of(context).pushNamed(ConfigPage.Route),
            ),
            IconButton(
              icon: Icon(Icons.help),
              onPressed: () => Navigator.of(context).pushNamed(TutorialPage.Route),
            )
          ],
          bottom: TabBar(
            tabs: [
              Tab(
                icon: Icon(WorkflowExec.IconData),
                text: "Workflows",
              ),
              Tab(
                icon: Icon(WorkflowTemplate.IconData),
                text: "Templates",
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
            WorkflowExecList(),
            WorkflowTemplateList(),
            RoomList(),
            TrackerList(),
          ],
        ),
      ),
    );
  }
}
