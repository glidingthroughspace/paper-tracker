import 'package:flutter/material.dart';

import '../client/tracker_client.dart';
import '../model/tracker.dart';

class TrackerList extends StatefulWidget {
  TrackerList({Key key}) : super(key: key);

  @override
  _TrackerListState createState() => _TrackerListState();
}

class _TrackerListState extends State<TrackerList> {
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
          List<Widget> listChildren = List<Widget>();
          for (Tracker tracker in snapshot.data) {
            listChildren.add(ListTile(
              title: Text("ID: ${tracker.id.toString()}; Label: ${tracker.label}")
            ));
          }
          return ListView(children: listChildren);
        } else if (snapshot.hasError) {
          return ListView(
            children: <Widget>[Text("${snapshot.error}")]
          );
        }

        // By default, show a loading spinner.
        return CircularProgressIndicator();
      }
    );
  }
}
