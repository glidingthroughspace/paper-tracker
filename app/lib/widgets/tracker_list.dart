import 'package:flutter/material.dart';
import 'package:paper_tracker/pages/tracker_page.dart';
import 'package:paper_tracker/widgets/card_list.dart';
import 'package:tuple/tuple.dart';

import '../client/tracker_client.dart';
import '../model/tracker.dart';

class TrackerList extends StatefulWidget {
  TrackerList({Key key}) : super(key: key);

  @override
  _TrackerListState createState() => _TrackerListState();
}

class _TrackerListState extends State<TrackerList>
    with AutomaticKeepAliveClientMixin {
  var trackerClient = TrackerClient();

  @override
  void initState() {
    super.initState();
  }

  @override
  Widget build(BuildContext context) {
    super.build(context);
    return FutureBuilder(
        future: trackerClient.getAllTrackers(),
        builder: (context, snapshot) {
          if (snapshot.hasData) {
            List<Tracker> trackerList = snapshot.data;
            List<Tuple2<String, Tracker>> titleObjectList = trackerList
                .map((tracker) => Tuple2(tracker.label, tracker))
                .toList();
            return CardList<Tracker>(
              titleObjectList: titleObjectList,
              onTap: (tracker) => Navigator.of(context)
                  .pushNamed(TrackerPage.Route, arguments: tracker.id),
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

  @override
  bool get wantKeepAlive => true;
}
