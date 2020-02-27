import 'package:flutter/material.dart';
import 'package:paper_tracker/client/tracker_client.dart';
import 'package:paper_tracker/model/tracker.dart';
import 'package:paper_tracker/pages/tracker_page.dart';
import 'package:paper_tracker/widgets/lists/card_list.dart';

class TrackerList extends StatefulWidget {
  TrackerList({Key key}) : super(key: key);

  @override
  _TrackerListState createState() => _TrackerListState();
}

class _TrackerListState extends State<TrackerList> {
  var trackerClient = TrackerClient();

  @override
  Widget build(BuildContext context) {
    return FutureBuilder(
        future: trackerClient.getAllTrackers(),
        builder: (context, snapshot) {
          if (snapshot.hasData) {
            List<Tracker> trackerList = snapshot.data;
            var dataList = trackerList
                .map((tracker) => CardListData(tracker.label, tracker.lastRoom.toString(), tracker))
                .toList();

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

  Future<void> onRefresh() async {
    setState(() {
      trackerClient.getAllTrackers(refresh: true);
    });
  }
}
