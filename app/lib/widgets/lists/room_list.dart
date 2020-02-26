import 'package:flutter/material.dart';
import 'package:paper_tracker/client/room_client.dart';
import 'package:paper_tracker/model/room.dart';
import 'package:paper_tracker/pages/room_page.dart';
import 'package:paper_tracker/widgets/card_list.dart';
import 'package:paper_tracker/widgets/dialogs/add_room_dialog.dart';
import 'package:tuple/tuple.dart';

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
            List<Tuple2<String, Room>> titleObjectList = roomList.map((room) => Tuple2(room.label, room)).toList();

            return Scaffold(
              body: CardList<Room>(
                titleObjectList: titleObjectList,
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
    setState(() {
      roomClient.getAllRooms(refresh: true);
    });
  }

  void onAddRoomButton() async {
    return showDialog(
      context: context,
      child: AddRoomDialog(labelController: roomLabelEditController, addRoom: addRoom),
    );
  }

  void addRoom() async {
    var room = Room(label: roomLabelEditController.text);
    await roomClient.addRoom(room);
    await roomClient.getAllRooms(refresh: true);

    setState(() {});
    Navigator.of(context).pop();
  }

  void onTapRoom(Room room) async {
    await Navigator.of(context).pushNamed(RoomPage.Route, arguments: room.id);
  }
}
