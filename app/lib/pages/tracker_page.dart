import 'package:flutter/material.dart';
import 'package:paper_tracker/client/tracker_client.dart';
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
          TableRow(children: [
            TableCell(child: Label("Label: ")),
            TableCell(
              child: TextFormField(
                controller: labelEditController,
                readOnly: !isEditing,
              ),
            ),
          ]),
          getTableSpacing(10.0),
          TableRow(
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
          ),
          getTableSpacing(10.0),
          TableRow(
            children: [
              TableCell(child: Label("Learn Room: ")),
              TableCell(
                child: Row(
                  mainAxisAlignment: MainAxisAlignment.spaceAround,
                  children: [
                    MaterialButton(
                      child: Text("Learn"),
                      onPressed: tracker.status == TrackerStatus.Idle ? () => onLearnButton(tracker.id) : null,
                      color: Theme.of(context).accentColor,
                    ),
                  ],
                ),
              ),
            ],
          ),
        ],
      ),
    );
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
      await trackerClient.getAllTrackers(refresh: true);
    }
    setState(() => isEditing = edit);
  }

  void delete(Tracker tracker) async {
    await trackerClient.deleteTracker(tracker.id);
    await trackerClient.getAllTrackers(refresh: true);
    Navigator.of(context).pop();
  }
}
