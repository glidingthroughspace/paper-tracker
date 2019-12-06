import 'package:flutter/material.dart';
import 'package:paper_tracker/client/room_client.dart';
import 'package:paper_tracker/model/room.dart';
import 'package:paper_tracker/pages/learning_page.dart';
import 'package:paper_tracker/widgets/conditional_builder.dart';
import 'package:paper_tracker/widgets/detail_content.dart';
import 'package:paper_tracker/widgets/label.dart';

class RoomPage extends StatefulWidget {
  static const String Route = "/page";

  @override
  _RoomPageState createState() => _RoomPageState();
}

class _RoomPageState extends State<RoomPage> {
  var isEditing = false;
  var labelEditController = TextEditingController();
  var roomClient = RoomClient();

  @override
  Widget build(BuildContext context) {
    var roomID = ModalRoute.of(context).settings.arguments;
    var futureRoom = roomClient.getRoomByID(roomID);

    return FutureBuilder(
      future: futureRoom,
      builder: (context, snapshot) {
        var title = "";
        Widget content;
        if (snapshot.hasData) {
          Room room = snapshot.data;
          labelEditController.text = room.label;
          title = room.label;
          content = buildContent(room);
        } else {
          content = CircularProgressIndicator();
        }

        return DetailContent(
          title: title,
          iconData: Room.IconData,
          bottomButtons: buildBottomButtons(),
          content: content,
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

  Widget buildContent(Room room) {
    return Container(
      padding: EdgeInsets.all(15.0),
      child: Table(
        defaultVerticalAlignment: TableCellVerticalAlignment.middle,
        columnWidths: {0: FractionColumnWidth(0.3)},
        children: [
          TableRow(children: [
            TableCell(child: Label("Label: ")),
            TextFormField(
              controller: labelEditController,
              readOnly: !isEditing,
            ),
          ]),
          TableRow(children: [
            TableCell(child: Label("Is Learned: ")),
            TableCell(
              child: Row(
                mainAxisAlignment: MainAxisAlignment.spaceAround,
                children: [
                  Icon(room.isLearned ? Icons.check : Icons.close, color: Colors.white),
                  MaterialButton(
                    child: Text(room.isLearned ? "Relearn" : "Learn now"),
                    onPressed: () => Navigator.of(context).pushNamed(LearningPage.Route, arguments: LearningPageParams(roomID: room.id)),
                    color: Theme.of(context).accentColor,
                  ),
                ],
              ),
            ),
          ]),
        ],
      ),
    );
  }

  void setEditing(bool edit) {
    if (edit == false) {
      // => Saving
    }
    setState(() => isEditing = edit);
  }
}
