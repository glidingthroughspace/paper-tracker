import 'package:flutter/material.dart';
import 'package:fluttertoast/fluttertoast.dart';
import 'package:paper_tracker/client/room_client.dart';
import 'package:paper_tracker/client/tracker_client.dart';
import 'package:paper_tracker/model/room.dart';
import 'package:paper_tracker/model/tracker.dart';
import 'package:paper_tracker/utils.dart';
import 'package:paper_tracker/widgets/conditional_builder.dart';
import 'package:paper_tracker/widgets/detail_content.dart';
import 'package:paper_tracker/widgets/label.dart';

import 'learning_page.dart';

class TrackerPage extends StatefulWidget {
  static const Route = "/tracker";

  @override
  _TrackerPageState createState() => _TrackerPageState();
}

class _TrackerPageState extends State<TrackerPage> {
  var isEditing = false;
  var trackerClient = TrackerClient();
  var roomClient = RoomClient();
  var labelEditController = TextEditingController();
  int trackerID;
  Future<Tracker> futureTracker;

  @override
  void didChangeDependencies() {
    super.didChangeDependencies();
    trackerID = ModalRoute.of(context).settings.arguments;
    futureTracker = trackerClient.getTrackerByID(trackerID);
    futureTracker.then((tracker) => labelEditController.text = tracker.label);
  }

  @override
  Widget build(BuildContext context) {
    return FutureBuilder(
      future: futureTracker,
      builder: (context, snapshot) {
        Widget content;
        Tracker tracker = snapshot.data;
        if (snapshot.hasData) {
          content = buildContent(tracker);
        } else {
          content = CircularProgressIndicator();
        }

        return DetailContent(
          title: tracker != null ? tracker.label : "",
          iconData: Tracker.IconData,
          bottomButtons: buildBottomButtons(tracker),
          content: content,
          onRefresh: refreshTracker,
        );
      },
    );
  }

  Widget buildContent(Tracker tracker) {
    return Container(
      padding: EdgeInsets.all(15.0),
      child: Table(
        defaultVerticalAlignment: TableCellVerticalAlignment.middle,
        columnWidths: {0: FractionColumnWidth(0.3)},
        children: [
          buildLabelRow(),
          getTableSpacing(10.0),
          buildRoomRow(tracker),
          getTableSpacing(10.0),
          buildBatteryRow(tracker),
          getTableSpacing(10.0),
          buildStatusRow(tracker),
          getTableSpacing(10.0),
          buildLearnRow(tracker),
        ],
      ),
    );
  }

  TableRow buildLearnRow(Tracker tracker) {
    return TableRow(
      children: [
        TableCell(child: Label("Learn Room: ")),
        TableCell(
          child: Row(
            mainAxisAlignment: MainAxisAlignment.spaceEvenly,
            children: [
              MaterialButton(
                child: Text("Learn"),
                onPressed: tracker.status == TrackerStatus.Idle ? () => onLearnButton(tracker.id) : null,
                color: Theme.of(context).accentColor,
                disabledColor: Theme.of(context).cardColor,
              ),
              MaterialButton(
                child: Text("Cancel Learning"),
                onPressed: tracker.status == TrackerStatus.Learning || tracker.status == TrackerStatus.LearningFinished
                    ? () => onLearnCancelButton(tracker.id)
                    : null,
                color: Theme.of(context).accentColor,
                disabledColor: Theme.of(context).cardColor,
              )
            ],
          ),
        ),
      ],
    );
  }

  TableRow buildStatusRow(Tracker tracker) {
    return TableRow(
      children: [
        TableCell(child: Label("Status: ")),
        TableCell(
          child: Row(
            children: [
              Icon(tracker.status.icon),
              Padding(padding: EdgeInsets.only(left: 10.0)),
              Label(tracker.status.label),
            ],
          ),
        ),
      ],
    );
  }

  TableRow buildLabelRow() {
    return TableRow(children: [
      TableCell(child: Label("Label: ")),
      TableCell(
        child: TextFormField(
          controller: labelEditController,
          readOnly: !isEditing,
        ),
      ),
    ]);
  }

  TableRow buildBatteryRow(Tracker tracker) {
    return TableRow(children: [
      TableCell(child: Label("Battery: ")),
      TableCell(
        child: Row(
          children: [
            Label(tracker.isCharging ? "Charging - " : ""),
            Label(tracker.batteryPercentage != null ? "${tracker.batteryPercentage}%" : "Unknown"),
          ],
        ),
      ),
    ]);
  }

  TableRow buildRoomRow(Tracker tracker) {
    return TableRow(children: [
      TableCell(child: Label("Last room: ")),
      TableCell(
          child: FutureBuilder(
        future: roomClient.getRoomByID(tracker.lastRoom),
        builder: (context, snapshot) {
          if (snapshot.hasData) {
            Room room = snapshot.data;
            return Label(room.label);
          }
          return Label("Unknown");
        },
      )),
    ]);
  }

  void onLearnButton(int trackerID) {
    Navigator.of(context).pushNamed(LearningPage.Route, arguments: LearningPageParams(trackerID: trackerID));
  }

  List<Widget> buildBottomButtons(Tracker tracker) {
    return [
      ConditionalBuilder(
        conditional: isEditing,
        truthy: IconButton(
          icon: Icon(Icons.save, color: Colors.white),
          onPressed: () => setEditing(tracker, false),
        ),
        falsy: IconButton(
          icon: Icon(Icons.edit, color: Colors.white),
          onPressed: () => setEditing(tracker, true),
        ),
      ),
      IconButton(
        icon: Icon(Icons.delete_forever, color: Colors.white),
        onPressed: () => delete(tracker),
      ),
    ];
  }

  void setEditing(Tracker tracker, bool edit) async {
    if (edit == false && tracker != null) {
      tracker.label = labelEditController.text;
      await trackerClient.updateTracker(tracker);
    }
    setState(() {
      isEditing = edit;
    });
    refreshTracker();
  }

  void delete(Tracker tracker) async {
    if (tracker.status != TrackerStatus.Idle) {
      Fluttertoast.showToast(msg: "Can't delete tracker that is not in idle status");
      return;
    }

    await trackerClient.deleteTracker(tracker.id);
    await trackerClient.getAllTrackers(refresh: true);
    Navigator.of(context).pop();
  }

  void onLearnCancelButton(int trackerID) async {
    await trackerClient.cancelLearning(trackerID);
    refreshTracker();
  }

  Future<void> refreshTracker() async {
    setState(() {
      futureTracker = trackerClient.getTrackerByID(trackerID, refresh: true);
      futureTracker.then((tracker) => labelEditController.text = tracker.label);
    });
  }
}
