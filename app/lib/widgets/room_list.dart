import 'package:flutter/material.dart';
import 'package:paper_tracker/client/room_client.dart';
import 'package:paper_tracker/model/room.dart';
import 'package:paper_tracker/pages/room_page.dart';
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
    fetchTrackers();
  }

  @override
  Widget build(BuildContext context) {
    super.build(context);
    return FutureBuilder(
        future: rooms,
        builder: (context, snapshot) {
          if (snapshot.hasData) {
            List<Room> roomList = snapshot.data;
            Map<String, Room> titleObjectMap =
                Map.fromIterable(roomList, key: (room) => room.label, value: (room) => room);

            return Scaffold(
              body: CardList<Room>(
                titleObjectMap: titleObjectMap,
                onTap: onTapRoom,
                iconData: Icons.keyboard_arrow_right,
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

  Widget buildAddRoomDialog(BuildContext context) {
    return AlertDialog(
      content: Column(
        mainAxisSize: MainAxisSize.min,
        children: [
          Text(
            "Add Room",
            style: TextStyle(
              fontSize: 20.0,
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
              enabledBorder: OutlineInputBorder(borderSide: BorderSide(color: Theme.of(context).accentColor)),
              focusedBorder: OutlineInputBorder(borderSide: BorderSide(color: Theme.of(context).accentColor)),
            ),
          ),
        ],
      ),
      actions: <Widget>[
        FlatButton(
          child: Text("Create"),
          onPressed: () => addRoom(),
        ),
      ],
    );
  }

  void addRoom() async {
    var room = Room(label: roomLabelEditController.text);
    await roomClient.addRoom(room);

    fetchTrackers();
    Navigator.of(context).pop();
  }

  void fetchTrackers() {
    rooms = roomClient.fetchRooms();
  }

  @override
  bool get wantKeepAlive => true;

  void onTapRoom(Room room) async {
    await Navigator.of(context).pushNamed(RoomPage.Route, arguments: room);
    fetchTrackers();
  }
}
