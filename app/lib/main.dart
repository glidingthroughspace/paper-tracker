import 'package:flutter/material.dart';
import 'package:paper_tracker/pages/config_page.dart';
import 'package:paper_tracker/pages/init_page.dart';
import 'package:paper_tracker/pages/main_page.dart';

void main() => runApp(PaperTrackerApp());

class PaperTrackerApp extends StatelessWidget {
  @override
  Widget build(BuildContext context) {
    return MaterialApp(
      title: "Paper Tracker",
      theme: ThemeData(
        primarySwatch: Colors.blue,
      ),
      initialRoute: "/",
      routes: {
        InitPage.Route: (context) => InitPage(),
        MainPage.Route: (context) => MainPage(),
        ConfigPage.Route: (context) => ConfigPage(),
      },
    );
  }
}
