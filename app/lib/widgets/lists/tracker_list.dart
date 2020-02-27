import 'package:flutter/material.dart';
import 'package:flutter/rendering.dart';
import 'package:paper_tracker/client/room_client.dart';
import 'package:paper_tracker/client/tracker_client.dart';
import 'package:paper_tracker/model/room.dart';
import 'package:paper_tracker/model/tracker.dart';
import 'package:paper_tracker/pages/tracker_page.dart';
import 'package:paper_tracker/widgets/conditional_builder.dart';
import 'package:paper_tracker/widgets/lists/card_list.dart';

class TrackerList extends StatefulWidget {
  TrackerList({Key key}) : super(key: key);

  @override
  _TrackerListState createState() => _TrackerListState();
}

class _TrackerListState extends State<TrackerList> {
  var trackerClient = TrackerClient();
  var roomClient = RoomClient();

  @override
  Widget build(BuildContext context) {
    return FutureBuilder(
        future: trackerClient.getAllTrackers(),
        builder: (context, snapshot) {
          if (snapshot.hasData) {
            List<Tracker> trackerList = snapshot.data;
            var dataList =
                trackerList.map((tracker) => CardListData(tracker.label, buildSubtitle(tracker), tracker)).toList();

            return CardList<Tracker>(
              dataList: dataList,
              onTap: (tracker) => Navigator.of(context).pushNamed(TrackerPage.Route, arguments: tracker.id),
              iconData: Icons.keyboard_arrow_right,
              onRefresh: onRefresh,
            );
          } else if (snapshot.hasError) {
            return Center(child: Text("${snapshot.error}"));
          }

          // By default, show a loading spinner.
          return Center(child: CircularProgressIndicator());
        });
  }

  List<Widget> buildSubtitle(Tracker tracker) {
    return [
      ConditionalBuilder(
        conditional: tracker.isCharging,
        truthy: Text("Charging"),
        falsy: Text("${tracker.batteryPercentage}%"),
      ),
      Padding(padding: EdgeInsets.only(left: 20.0)),
      FutureBuilder(
        future: roomClient.getRoomByID(tracker.lastRoom),
        builder: (context, snapshot) {
          if (snapshot.hasData) {
            Room room = snapshot.data;
            return Text("Room: ${room.label}");
          }
          return Text("Room: Unknown");
        },
      )
    ];
  }

  Future<void> onRefresh() async {
    setState(() {
      trackerClient.getAllTrackers(refresh: true);
    });
  }
}
