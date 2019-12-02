import 'package:flutter/material.dart';
import 'package:paper_tracker/client/room_client.dart';
import 'package:paper_tracker/client/tracker_client.dart';
import 'package:paper_tracker/model/room.dart';
import 'package:paper_tracker/model/tracker.dart';
import 'package:paper_tracker/widgets/detail_content.dart';

class LearningPage extends StatefulWidget {
  static const Route = "/learning";

  @override
  _LearningPageState createState() => _LearningPageState();
}

class _LearningPageState extends State<LearningPage> {
  var trackerClient = TrackerClient();
  var roomClient = RoomClient();
  Future<List<Room>> rooms;
  Future<List<Tracker>> tracker;

  Room selectedRoom;
  Tracker selectedTracker;

  @override
  void initState() {
    super.initState();

    rooms = roomClient.fetchRooms();
    tracker = trackerClient.fetchTrackers();
  }

  @override
  Widget build(BuildContext context) {
    LearningPageParams params = ModalRoute.of(context).settings.arguments;

    return DetailContent(
      title: "Learn Room",
      iconData: Icons.school,
      content: Container(
        padding: EdgeInsets.all(15.0),
        child: Column(
          children: [
            buildRoomDropdown(params),
            buildTrackerDropdown(params),
          ],
        ),
      ),
    );
  }

  Widget buildRoomDropdown(LearningPageParams params) {
    return FutureBuilder(
      future: rooms,
      builder: (context, snapshot) {
        List<Room> roomList = snapshot.hasData ? snapshot.data : [];
        if (snapshot.hasData && selectedRoom == null && params.roomID != null) {
          selectedRoom = roomList.firstWhere((room) => room.id == params.roomID);
        }
        return DropdownButton(
          icon: Icon(Room.IconData),
          items: roomList.map((room) => DropdownMenuItem(value: room, child: Text(room.label))).toList(),
          value: selectedRoom,
          isExpanded: true,
          style: TextStyle(color: Colors.white),
          onChanged: (value) {
            setState(() {
              selectedRoom = value;
            });
          },
          hint: Text("Please select a room", style: TextStyle(color: Colors.white)),
        );
      },
    );
  }

  Widget buildTrackerDropdown(LearningPageParams params) {
    return FutureBuilder(
      future: tracker,
      builder: (context, snapshot) {
        List<Tracker> trackerList = snapshot.hasData ? snapshot.data : [];
        if (snapshot.hasData && selectedTracker == null && params.trackerID != null) {
          selectedTracker = trackerList.firstWhere((tracker) => tracker.id == params.trackerID);
        }
        return DropdownButton(
          icon: Icon(Tracker.IconData),
          items: trackerList
              .map((tracker) => DropdownMenuItem(
                    value: tracker,
                    child: Text(tracker.label),
                  ))
              .toList(),
          value: selectedTracker,
          style: TextStyle(color: Colors.white),
          isExpanded: true,
          onChanged: (value) {
            setState(() {
              selectedTracker = value;
            });
          },
          hint: Text("Please select a tracker", style: TextStyle(color: Colors.white)),
        );
      },
    );
  }
}

class LearningPageParams {
  int roomID;
  int trackerID;

  LearningPageParams({this.roomID, this.trackerID});
}
