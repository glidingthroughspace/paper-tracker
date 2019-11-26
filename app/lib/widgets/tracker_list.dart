import 'package:flutter/material.dart';
import 'package:paper_tracker/widgets/card_list.dart';

import '../client/tracker_client.dart';
import '../model/tracker.dart';

class TrackerList extends StatefulWidget {
  TrackerList({Key key}) : super(key: key);

  @override
  _TrackerListState createState() => _TrackerListState();
}

class _TrackerListState extends State<TrackerList> with AutomaticKeepAliveClientMixin {
  Future<List<Tracker>> trackers;

  @override
  void initState() {
    super.initState();
    trackers = TrackerClient().fetchTrackers();
  }

  @override
  Widget build(BuildContext context) {
    return FutureBuilder(
        future: trackers,
        builder: (context, snapshot) {
          if (snapshot.hasData) {
            List<Tracker> trackerList = snapshot.data;
            Map<String, Tracker> titleObjectMap =
                Map.fromIterable(trackerList, key: (tracker) => tracker.label, value: (tracker) => tracker);
            return CardList<Tracker>(
              titleObjectMap: titleObjectMap,
              onTap: (tracker) {
                print(tracker);
              },
            );
          } else if (snapshot.hasError) {
            return Center(child: Text("${snapshot.error}"));
          }

          // By default, show a loading spinner.
          return Center(child: CircularProgressIndicator());
        });
  }

  @override
  bool get wantKeepAlive => true;
}
