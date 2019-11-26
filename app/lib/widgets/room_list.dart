import 'package:flutter/material.dart';
import 'package:paper_tracker/client/room_client.dart';
import 'package:paper_tracker/model/room.dart';
import 'package:paper_tracker/widgets/card_list.dart';

class RoomList extends StatefulWidget {
  RoomList({Key key}) : super(key: key);

  @override
  _RoomListState createState() => _RoomListState();
}

class _RoomListState extends State<RoomList> with AutomaticKeepAliveClientMixin {
  var roomLabelEditController = TextEditingController();
  var roomClient = RoomClient();
  Future<List<Room>> rooms;

  @override
  void initState() {
    super.initState();
    rooms = roomClient.fetchRooms();
  }

  @override
  Widget build(BuildContext context) {
    return FutureBuilder(
        future: rooms,
        builder: (context, snapshot) {
          if (snapshot.hasData) {
            List<Room> roomList = snapshot.data;
            Map<String, Room> titleObjectMap =
                Map.fromIterable(roomList, key: (room) => room.label, value: (room) => room);

            return Scaffold(
              backgroundColor: Theme.of(context).backgroundColor,
              body: CardList<Room>(
                titleObjectMap: titleObjectMap,
                onTap: (room) {
                  showDialog(context: context, builder: buildAddRoomDialog);
                },
              ),
              floatingActionButton: FloatingActionButton(
                onPressed: onAddRoomButton,
                child: Icon(Icons.add),
              ),
            );
          } else if (snapshot.hasError) {
            return Center(child: Text("${snapshot.error}"));
          }

          // By default, show a loading spinner.
          return Center(child: CircularProgressIndicator());
        });
  }

  void onAddRoomButton() async {
    return showDialog(
      context: context,
      builder: buildAddRoomDialog,
    );
  }

  Widget buildAddRoomDialog(context) {
    return AlertDialog(
      content: Column(mainAxisSize: MainAxisSize.min, children: [
        Text(
          "Add Room",
          style: TextStyle(
            fontSize: 20.0,
            color: Colors.deepOrange,
          ),
        ),
        Padding(
          padding: EdgeInsets.only(top: 10.0),
        ),
        TextFormField(
          controller: roomLabelEditController,
          autofocus: true,
          decoration: InputDecoration(
            labelText: "Room Label",
            border: OutlineInputBorder(),
          ),
        ),
      ]),
      actions: <Widget>[
        FlatButton(
          child: Text("Create"),
          onPressed: () => addRoom(),
        )
      ],
    );
  }

  void addRoom() async {
    var room = Room(label: roomLabelEditController.text);
    await roomClient.addRoom(room);

    rooms = roomClient.fetchRooms();
    Navigator.of(context).pop();
  }

  @override
  bool get wantKeepAlive => true;
}
