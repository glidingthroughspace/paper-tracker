import 'package:flutter/material.dart';
import 'package:paper_tracker/client/api_client.dart';
import 'package:paper_tracker/pages/config_page.dart';
import 'package:paper_tracker/pages/main_page.dart';

void main() => runApp(PaperTrackerApp());

class PaperTrackerApp extends StatefulWidget {
  @override
  _PaperTrackerAppState createState() => _PaperTrackerAppState();
}

class _PaperTrackerAppState extends State<PaperTrackerApp> {
  Future<bool> serverAvailable;

  @override
  void initState() {
    super.initState();
    serverAvailable = APIClient().isAvailable();
  }

  @override
  Widget build(BuildContext context) {
    return FutureBuilder(
      future: serverAvailable,
      builder: (context, snapshot) {
        Widget mainWidget;
        if (snapshot.hasData) {
          if (snapshot.data == true)
            mainWidget = MainPage();
          else
            mainWidget = ConfigPage();
        } else {
          mainWidget = Scaffold(
            body: Center(child: CircularProgressIndicator()),
          );
        }

        return MaterialApp(
          title: "Paper Tracker",
          theme: ThemeData(
            primarySwatch: Colors.blue,
          ),
          home: mainWidget,
        );
      },
    );
  }
}
