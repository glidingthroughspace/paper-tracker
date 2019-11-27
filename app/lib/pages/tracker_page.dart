import 'package:flutter/material.dart';
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
  Tracker tracker;

  @override
  Widget build(BuildContext context) {
    tracker = ModalRoute.of(context).settings.arguments;

    return DetailContent(
      title: tracker.label,
      iconData: Icons.track_changes,
      bottomButtons: [
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
      ],
      content: Text("Tracker"),
    );
  }

  void setEditing(bool edit) {
    if (edit == false) { // => Save

    }
    setState(() => isEditing = edit);
  }
}