import 'dart:convert';

import 'package:http/http.dart' as http;
import 'package:flutter/material.dart';

import 'model/tracker.dart';

class TrackerList extends StatefulWidget {
  TrackerList({Key key}) : super(key: key);

  @override
  _TrackerListState createState() => _TrackerListState();
}

class _TrackerListState extends State<TrackerList> {
  Future<List<Tracker>> trackers;

  Future<List<Tracker>> fetchTrackers() async {
    final response = await http.get('http://192.168.0.164:8080/tracker');
    if (response.statusCode == 200) {
      final rawList = json.decode(response.body) as List;
      return rawList.map((i) => Tracker.fromJson(i)).toList();
    } else {
      throw Exception("Failed to load trackers");
    }
  }

  @override
  void initState() {
    super.initState();
    trackers = fetchTrackers();
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
