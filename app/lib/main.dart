import 'package:flutter/material.dart';
import 'package:paper_tracker/pages/config_page.dart';
import 'package:paper_tracker/pages/init_page.dart';
import 'package:paper_tracker/pages/learning_page.dart';
import 'package:paper_tracker/pages/main_page.dart';
import 'package:paper_tracker/pages/room_page.dart';
import 'package:paper_tracker/pages/tracker_page.dart';

void main() => runApp(PaperTrackerApp());

class PaperTrackerApp extends StatelessWidget {
  @override
  Widget build(BuildContext context) {
    return MaterialApp(
      title: "Paper Tracker",
      theme: ThemeData(
        backgroundColor: Color.fromRGBO(56, 66, 86, 1.0),
        cardColor: Color.fromRGBO(84, 93, 110, .9),
        accentColor: Color.fromRGBO(148, 0, 238, 1.0),
        canvasColor: Color.fromRGBO(56, 66, 86, 1.0),
      ),
      initialRoute: "/",
      routes: {
        InitPage.Route: (context) => InitPage(),
        MainPage.Route: (context) => MainPage(),
        ConfigPage.Route: (context) => ConfigPage(),
        RoomPage.Route: (context) => RoomPage(),
        TrackerPage.Route: (context) => TrackerPage(),
        LearningPage.Route: (context) => LearningPage(),
      },
    );
  }
}
