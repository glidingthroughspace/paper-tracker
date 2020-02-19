import 'dart:async';

import 'package:flutter/material.dart';
import 'package:paper_tracker/client/room_client.dart';
import 'package:paper_tracker/client/tracker_client.dart';
import 'package:paper_tracker/model/room.dart';
import 'package:paper_tracker/model/tracker.dart';
import 'package:paper_tracker/widgets/card_list.dart';
import 'package:paper_tracker/widgets/conditional_builder.dart';
import 'package:paper_tracker/widgets/countdown_timer.dart';
import 'package:paper_tracker/widgets/detail_content.dart';

class LearningPage extends StatefulWidget {
  static const Route = "/learning";

  @override
  _LearningPageState createState() => _LearningPageState();
}

class _LearningPageState extends State<LearningPage> {
  var trackerClient = TrackerClient();
  var roomClient = RoomClient();
  Future<List<Tracker>> tracker;
  bool countdownDone = false;
  Timer ssidTimer;

  var state = _learningState.Init;
  Room selectedRoom;
  Tracker selectedTracker;
  int learnDuration = 0;
  CheckCardListController checkCardListController;

  @override
  void initState() {
    super.initState();

    tracker = trackerClient.getAllTrackers();
    checkCardListController = CheckCardListController();
  }

  @override
  void dispose() {
    super.dispose();
    ssidTimer?.cancel();
  }

  @override
  Widget build(BuildContext context) {
    LearningPageParams params = ModalRoute.of(context).settings.arguments;

    return DetailContent(
      disableBackNav: state == _learningState.Running,
      title: "Learn Room",
      iconData: Icons.school,
      bottomButtons: [
        IconButton(
          icon: Icon(Icons.check),
          onPressed: state == _learningState.Finished ? onSave : null,
        ),
      ],
      content: Container(
        padding: EdgeInsets.all(15.0),
        child: Column(
          children: [
            buildRoomDropdown(params),
            buildTrackerDropdown(params),
            SizedBox(height: 15.0),
            buildButtonOrCountdown(context),
            SizedBox(height: 15.0),
            ...buildSSIDList()
          ],
        ),
      ),
    );
  }

  ConditionalBuilder buildButtonOrCountdown(BuildContext context) {
    return ConditionalBuilder(
      conditional: state == _learningState.Init,
      truthy: MaterialButton(
        onPressed: onStartLearning,
        child: Text("Start learning"),
        color: Theme.of(context).accentColor,
        minWidth: MediaQuery.of(context).size.width * 0.8,
      ),
      falsy: CountdownTimer(
        duration: Duration(seconds: learnDuration),
        backgroundColor: Theme.of(context).cardColor,
        color: Theme.of(context).accentColor,
        onComplete: () {
          countdownDone = true;
        },
      ),
    );
  }

  Widget buildRoomDropdown(LearningPageParams params) {
    return FutureBuilder(
      future: roomClient.getAllRooms(),
      builder: (context, snapshot) {
        List<Room> roomList = snapshot.hasData ? snapshot.data : [];
        if (snapshot.hasData && selectedRoom == null && params.roomID != null) {
          selectedRoom = roomList.firstWhere((room) => room.id == params.roomID);
        }
        return DropdownButton(
          icon: Icon(Room.IconData),
          items: roomList.map((room) => DropdownMenuItem(value: room, child: Text(room.label))).toList(),
          value: snapshot.hasData ? selectedRoom : null,
          isExpanded: true,
          onChanged: state != _learningState.Init ? null : (value) {
            setState(() {
              selectedRoom = value;
            });
          },
          hint: Text("Please select a room"),
          disabledHint: Text("No rooms found"),
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
              .where((tracker) => tracker.status == TrackerStatus.Idle)
              .map((tracker) => DropdownMenuItem(
                    value: tracker,
                    child: Text(tracker.label),
                  ))
              .toList(),
          value: selectedTracker,
          isExpanded: true,
          onChanged: state != _learningState.Init ? null : (value) {
            setState(() {
              selectedTracker = value;
            });
          },
          hint: Text("Please select a tracker"),
          disabledHint: Text("No tracker available for learning"),
        );
      },
    );
  }

  List<Widget> buildSSIDList() {
    if (state == _learningState.Init) {
      return [];
    }

    return [
      Align(
        alignment: Alignment.centerLeft,
        child: Text(
          "Found SSIDs:",
          style: TextStyle(fontSize: 20.0, fontWeight: FontWeight.bold),
          textAlign: TextAlign.left,
        ),
      ),
      CheckCardList(
        controller: checkCardListController,
      ),
    ];
  }

  void onStartLearning() async {
    var resp = await trackerClient.startLearning(selectedTracker.id);
    setState(() {
      state = _learningState.Running;
      learnDuration = resp.learnTimeSec;
    });
    ssidTimer = Timer.periodic(Duration(seconds: 1), getLearnStatus);
  }

  void getLearnStatus(Timer t) async {
    var resp = await trackerClient.getLearningStatus(selectedTracker.id);
    setState(() {
      checkCardListController.updateFromTitles(resp.ssids);
      if (resp.done && countdownDone) {
        state = _learningState.Finished;
      }
    });
  }

  void onSave() async {
    ssidTimer.cancel();
    await trackerClient.finishLearning(selectedTracker.id, selectedRoom.id, checkCardListController.checked);
    await roomClient.getAllRooms(refresh: true);
    Navigator.of(context).pop();
  }
}

class LearningPageParams {
  int roomID;
  int trackerID;

  LearningPageParams({this.roomID, this.trackerID});
}

enum _learningState { Init, Running, Finished }
