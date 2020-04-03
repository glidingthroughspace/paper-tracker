import 'dart:async';

import 'package:flutter/cupertino.dart';
import 'package:flutter/foundation.dart';
import 'package:flutter/material.dart';
import 'package:paper_tracker/client/room_client.dart';
import 'package:paper_tracker/client/tracker_client.dart';
import 'package:paper_tracker/model/room.dart';
import 'package:paper_tracker/model/tracker.dart';
import 'package:paper_tracker/widgets/conditional_builder.dart';
import 'package:paper_tracker/widgets/countdown_timer.dart';
import 'package:paper_tracker/widgets/detail_content.dart';
import 'package:paper_tracker/widgets/dialogs/confirm_icon_text_dialog.dart';
import 'package:paper_tracker/widgets/dialogs/wait_tracker_poll_dialog.dart';
import 'package:paper_tracker/widgets/dropdown.dart';
import 'package:paper_tracker/widgets/label.dart';
import 'package:paper_tracker/widgets/lists/check_card_list.dart';

class LearningPage extends StatefulWidget {
  static const Route = "/learning";

  @override
  _LearningPageState createState() => _LearningPageState();
}

class _LearningPageState extends State<LearningPage> {
  var trackerClient = TrackerClient();
  var roomClient = RoomClient();
  bool countdownDone = false;
  Timer ssidTimer;

  var state = _learningState.Init;
  int learnDuration = 0;
  var checkCardListController = CheckCardListController();
  var roomDropdownController = DropdownController();
  var trackerDropdownController = DropdownController();

  @override
  void dispose() {
    ssidTimer?.cancel();
    super.dispose();
  }

  @override
  Widget build(BuildContext context) {
    LearningPageParams params = ModalRoute.of(context).settings.arguments;
    roomDropdownController.defaultID = params.roomID;
    trackerDropdownController.defaultID = params.trackerID;

    return WillPopScope(
      onWillPop: onBackPressed,
      child: DetailContent(
        onBack: () => Navigator.of(context).maybePop(),
        title: "Learn Room",
        iconData: Icons.school,
        bottomButtons: [
          IconButton(
            icon: Icon(Icons.check),
            onPressed: state == _learningState.Finished && (checkCardListController.checked.isNotEmpty || !kReleaseMode)
                ? onSave
                : null,
          ),
        ],
        content: Container(
          padding: EdgeInsets.all(15.0),
          child: Column(
            children: [
              Row(children: [Label("Select a room to learn:"), Spacer()]),
              buildRoomDropdown(),
              Row(children: [Label("Select a tracker to learn with:"), Spacer()]),
              buildTrackerDropdown(),
              SizedBox(height: 15.0),
              buildButtonOrCountdown(context),
              SizedBox(height: 15.0),
              ...buildSSIDList()
            ],
          ),
        ),
      ),
    );
  }

  Widget buildRoomDropdown() {
    return Dropdown(
      getItems: () => roomClient.getAllRooms(refresh: true),
      controller: roomDropdownController,
      hintName: "room",
      icon: Room.IconData,
      itemFixed: state != _learningState.Init,
      setState: setState,
    );
  }

  ConditionalBuilder buildButtonOrCountdown(BuildContext context) {
    return ConditionalBuilder(
      conditional: state == _learningState.Init,
      truthy: MaterialButton(
        onPressed: roomDropdownController.selectedItem != null && trackerDropdownController.selectedItem != null
            ? onWaitForPoll
            : null,
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

  Widget buildTrackerDropdown() {
    return Dropdown(
      getItems: () async {
        var allTrackers = await trackerClient.getAllTrackers(refresh: true);
        return allTrackers.where((tracker) => tracker.status == TrackerStatus.Idle).toList();
      },
      controller: trackerDropdownController,
      hintName: "tracker",
      icon: Tracker.IconData,
      itemFixed: state != _learningState.Init,
      setState: setState,
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

  void onWaitForPoll() async {
    showDialog(
      context: context,
      child: WaitTrackerPollDialog(
        trackerID: trackerDropdownController.selectedItem.id,
        onWaitFinished: onStartLearning,
        trackerClient: trackerClient,
      ),
    );
  }

  void onStartLearning() async {
    var resp = await trackerClient.startLearning(trackerDropdownController.selectedItem.id);
    setState(() {
      state = _learningState.Running;
      learnDuration = resp.learnTimeSec;
    });
    ssidTimer = Timer.periodic(Duration(seconds: 1), getLearnStatus);
  }

  void getLearnStatus(Timer t) async {
    var resp = await trackerClient.getLearningStatus(trackerDropdownController.selectedItem.id);
    setState(() {
      checkCardListController.updateFromTitles(resp.ssids);
      if (resp.done && countdownDone) {
        state = _learningState.Finished;
      }
    });
  }

  void onSave() async {
    ssidTimer.cancel();
    await trackerClient.finishLearning(trackerDropdownController.selectedItem.id,
        roomDropdownController.selectedItem.id, checkCardListController.checked);
    await roomClient.getAllRooms(refresh: true);
    Navigator.of(context).pop();
  }

  Future<bool> onBackPressed() async {
    if (state == _learningState.Init) {
      return true;
    }

    return showDialog(
          context: context,
          builder: (context) => ConfirmIconTextDialog(
            text: "Do you want to cancel learning?",
            icon: Icons.question_answer,
            actions: {
              "No": () => Navigator.of(context).pop(false),
              "Yes": onCancelLearning,
            },
          ),
        ) ??
        false;
  }

  void onCancelLearning() async {
    ssidTimer?.cancel();
    await trackerClient.cancelLearning(trackerDropdownController.selectedItem.id);
    Navigator.of(context).pop(true);
  }
}

class LearningPageParams {
  int roomID;
  int trackerID;

  LearningPageParams({this.roomID, this.trackerID});
}

enum _learningState { Init, Running, Finished }
