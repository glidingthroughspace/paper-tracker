import 'package:flutter/material.dart';
import 'package:paper_tracker/client/room_client.dart';
import 'package:paper_tracker/model/room.dart';
import 'package:paper_tracker/pages/room_page.dart';
import 'package:paper_tracker/widgets/dialogs/single_text_dialog.dart';
import 'package:paper_tracker/widgets/lists/card_list.dart';

class RoomList extends StatefulWidget {
  RoomList({Key key}) : super(key: key);

  @override
  _RoomListState createState() => _RoomListState();
}

class _RoomListState extends State<RoomList> {
  var roomLabelEditController = TextEditingController();
  var roomClient = RoomClient();

  @override
  Widget build(BuildContext context) {
    return FutureBuilder(
        future: roomClient.getAllRooms(),
        builder: (context, snapshot) {
          if (snapshot.hasData) {
            List<Room> roomList = snapshot.data;
            roomList.sort((a, b) => a.label.compareTo(b.label));
            var dataList = roomList.map((room) => CardListData(room.label, buildSubtitle(room), room)).toList();

            return Scaffold(
              body: CardList<Room>(
                dataList: dataList,
                onTap: onTapRoom,
                iconData: Icons.keyboard_arrow_right,
                onRefresh: onRefresh,
              ),
              floatingActionButton: FloatingActionButton(
                onPressed: onAddRoomButton,
                child: Icon(Icons.add),
                heroTag: "roomAddButton",
              ),
            );
          } else if (snapshot.hasError) {
            return Center(child: Text("${snapshot.error}"));
          }

          // By default, show a loading spinner.
          return Center(child: CircularProgressIndicator());
        });
  }

  Future<void> onRefresh() async {
    await roomClient.getAllRooms(refresh: true);
    setState(() {});
  }

  void onAddRoomButton() async {
    return showDialog(
      context: context,
      child: SingleTextDialog(
        title: "Add Room",
        textLabel: "Room Label",
        buttonLabel: "Create",
        labelController: roomLabelEditController,
        onButton: addRoom,
      ),
    );
  }

  void addRoom() async {
    var room = Room(label: roomLabelEditController.text);
    await roomClient.addRoom(room);

    onRefresh();
    Navigator.of(context).pop();
  }

  void onTapRoom(Room room) async {
    await Navigator.of(context).pushNamed(RoomPage.Route, arguments: room.id);
  }

  List<Widget> buildSubtitle(Room room) {
    return [
      Text("Learned:"),
      Icon(room.isLearned ? Icons.check : Icons.close, color: Colors.grey, size: 20.0),
    ];
  }
}
