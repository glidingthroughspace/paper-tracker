import 'package:flutter/material.dart';
import 'package:paper_tracker/client/room_client.dart';
import 'package:paper_tracker/model/room.dart';

class RoomDropdownController {
  Room selectedRoom;
  int defaultID;
}

class RoomDropdown extends StatefulWidget {
  final RoomClient roomClient;
  final RoomDropdownController controller;

  const RoomDropdown({Key key, @required this.roomClient, @required this.controller}) : super(key: key);

  @override
  _RoomDropdownState createState() => _RoomDropdownState();
}

class _RoomDropdownState extends State<RoomDropdown> {
  @override
  Widget build(BuildContext context) {
    return FutureBuilder(
      future: widget.roomClient.getAllRooms(),
      builder: (context, snapshot) {
        List<Room> roomList = snapshot.hasData ? snapshot.data : [];
        if (snapshot.hasData && widget.controller.selectedRoom == null && widget.controller.defaultID != null) {
          widget.controller.selectedRoom = roomList.firstWhere((room) => room.id == widget.controller.defaultID);
        }

        return DropdownButton(
          icon: Icon(Room.IconData),
          items: roomList.map((room) => DropdownMenuItem(value: room, child: Text(room.label))).toList(),
          value: snapshot.hasData ? widget.controller.selectedRoom : null,
          isExpanded: true,
          onChanged: (value) {
            setState(() {
              widget.controller.selectedRoom = value;
            });
          },
          hint: Text("Please select a room"),
          disabledHint: Text("No rooms found"),
        );
      },
    );
  }
}
