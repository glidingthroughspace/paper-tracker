import 'package:flutter/material.dart';
import 'package:paper_tracker/widgets/tracker_list.dart';

class MainPage extends StatelessWidget {
  @override
  Widget build(BuildContext context) {
    return Scaffold(
      appBar: AppBar(
        title: Text("Paper Tracker"),
      ),
      body: TrackerList(),
    );
  }
}
