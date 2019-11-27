import 'package:flutter/material.dart';
import 'package:paper_tracker/model/room.dart';
import 'package:paper_tracker/widgets/conditional_builder.dart';
import 'package:paper_tracker/widgets/detail_content.dart';

class RoomPage extends StatefulWidget {
  static const String Route = "/page";

  @override
  _RoomPageState createState() => _RoomPageState();
}

class _RoomPageState extends State<RoomPage> {
  var isEditing = false;
  Room room;

  @override
  Widget build(BuildContext context) {
    room = ModalRoute.of(context).settings.arguments;

    return DetailContent(
      title: room.label,
      iconData: Icons.room,
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
      content: Text("Content"),
    );
  }

  void setEditing(bool edit) {
    if (edit == false) { // => Saving

    }
    setState(() => isEditing = edit);
  }
}
