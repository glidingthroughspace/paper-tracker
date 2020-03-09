import 'package:flutter/material.dart';
import 'package:fluttertoast/fluttertoast.dart';
import 'package:paper_tracker/client/room_client.dart';
import 'package:paper_tracker/model/room.dart';
import 'package:paper_tracker/pages/learning_page.dart';
import 'package:paper_tracker/utils.dart';
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
  int roomID;
  Future<Room> futureRoom;

  @override
  void didChangeDependencies() {
    super.didChangeDependencies();
    roomID = ModalRoute.of(context).settings.arguments;
    futureRoom = roomClient.getRoomByID(roomID);
    futureRoom.then((room) => labelEditController.text = room.label);
  }

  @override
  Widget build(BuildContext context) {
    return FutureBuilder(
      future: futureRoom,
      builder: (context, snapshot) {
        Widget content;
        Room room = snapshot.data;
        if (snapshot.hasData) {
          content = buildContent(room);
        } else {
          content = CircularProgressIndicator();
        }

        return DetailContent(
          title: room != null ? room.label : "",
          iconData: Room.IconData,
          bottomButtons: buildBottomButtons(room),
          content: content,
          onRefresh: refreshRoom,
        );
      },
    );
  }

  List<Widget> buildBottomButtons(Room room) {
    return [
      ConditionalBuilder(
        conditional: isEditing,
        truthy: IconButton(
          icon: Icon(Icons.save, color: Colors.white),
          onPressed: () => setEditing(room, false),
        ),
        falsy: IconButton(
          icon: Icon(Icons.edit, color: Colors.white),
          onPressed: () => setEditing(room, true),
        ),
      ),
      IconButton(
        icon: Icon(Icons.delete_forever, color: Colors.white),
        onPressed: () => delete(room),
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
          getTableSpacing(10.0),
          TableRow(children: [
            TableCell(child: Label("Is Learned: ")),
            TableCell(
              child: Row(
                mainAxisAlignment: MainAxisAlignment.spaceAround,
                children: [
                  Icon(room.isLearned ? Icons.check : Icons.close, color: Colors.white),
                  MaterialButton(
                    child: Text(room.isLearned ? "Relearn" : "Learn now"),
                    onPressed: () => Navigator.of(context)
                        .pushNamed(LearningPage.Route, arguments: LearningPageParams(roomID: room.id)),
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

  void setEditing(Room room, bool edit) async {
    if (edit == false && room != null) {
      room.label = labelEditController.text;
      await roomClient.updateRoom(room);
    }
    setState(() {
      isEditing = edit;
    });
    refreshRoom();
  }

  void delete(Room room) async {
    if (room.deleteLocked) {
      Fluttertoast.showToast(msg: "Can't delete room that is in use");
      return;
    }

    await roomClient.deleteRoom(room.id);
    await roomClient.getAllRooms(refresh: true);
    Navigator.of(context).pop();
  }

  Future<void> refreshRoom() async {
    setState(() {
      futureRoom = roomClient.getRoomByID(roomID, refresh: true);
      futureRoom.then((room) => labelEditController.text = room.label);
    });
  }
}
