import 'package:flutter/material.dart';
import 'package:paper_tracker/client/room_client.dart';
import 'package:paper_tracker/model/room.dart';

import '../client/tracker_client.dart';
import '../model/tracker.dart';

class RoomList extends StatefulWidget {
  RoomList({Key key}) : super(key: key);

  @override
  _RoomListState createState() => _RoomListState();
}

class _RoomListState extends State<RoomList> {
  var roomLabelEditController = TextEditingController();
  Future<List<Room>> rooms;

  @override
  void initState() {
    super.initState();
    rooms = RoomClient().fetchRooms();
  }

  @override
  Widget build(BuildContext context) {
    return FutureBuilder(
        future: rooms,
        builder: (context, snapshot) {
          if (snapshot.hasData) {
            List<Widget> listChildren = List<Widget>();
            for (Room room in snapshot.data) {
              listChildren.add(ListTile(
                title: Text(room.label),
                onTap: () {},
              ));
              listChildren.add(Divider());
            }
            return Scaffold(
              body: ListView(
                children: listChildren,
                shrinkWrap: true,
              ),
              floatingActionButton: FloatingActionButton(
                onPressed: onAddRoom,
                child: Icon(Icons.add),
              ),
            );
          } else if (snapshot.hasError) {
            return ListView(children: <Widget>[Text("${snapshot.error}")]);
          }

          // By default, show a loading spinner.
          return Center(child: CircularProgressIndicator());
        });
  }

  void onAddRoom() async {
    return showDialog(
        context: context,
        builder: (context) {
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
                decoration: InputDecoration(
                  labelText: "Room Label",
                  border: OutlineInputBorder(),
                ),
              ),
            ]),
            actions: <Widget>[
              FlatButton(
                child: Text("Create"),
                onPressed: () => {},
              )
            ],
          );
        });
  }
}