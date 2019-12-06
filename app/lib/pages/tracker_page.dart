import 'package:flutter/material.dart';
import 'package:paper_tracker/client/tracker_client.dart';
import 'package:paper_tracker/model/tracker.dart';
import 'package:paper_tracker/widgets/conditional_builder.dart';
import 'package:paper_tracker/widgets/detail_content.dart';

class TrackerPage extends StatefulWidget {
  static const Route = "/tracker";

  @override
  _TrackerPageState createState() => _TrackerPageState();
}

class _TrackerPageState extends State<TrackerPage> {
  var isEditing = false;
  var trackerClient = TrackerClient();

  @override
  Widget build(BuildContext context) {
    var trackerID = ModalRoute
        .of(context)
        .settings
        .arguments;
    var futureTracker = trackerClient.getTrackerByID(trackerID);

    return FutureBuilder(
      future: futureTracker,
      builder: (context, snapshot) {
        Tracker tracker = snapshot.data;

        return DetailContent(
          title: tracker != null ? tracker.label : "",
          iconData: Tracker.IconData,
          bottomButtons: buildBottomButtons(),
          content: Text("Tracker"),
        );
      },
    );
  }

  List<Widget> buildBottomButtons() {
    return [
      ConditionalBuilder(
        conditional: isEditing,
        truthy: IconButton(
          icon: Icon(Icons.save, color: Colors.white),
          onPressed: () => setEditing(false),
        ),
        falsy: IconButton(
          icon: Icon(Icons.edit, color: Colors.white),
          onPressed: () => setEditing(true),
        ),
      ),
      IconButton(
        icon: Icon(Icons.delete_forever, color: Colors.white),
        onPressed: () {},
      ),
    ];
  }

  void setEditing(bool edit) {
    if (edit == false) {
      // => Save

    }
    setState(() => isEditing = edit);
  }
}
