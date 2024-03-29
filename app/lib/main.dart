import 'package:flutter/material.dart';
import 'package:flutter/services.dart';
import 'package:paper_tracker/pages/config_page.dart';
import 'package:paper_tracker/pages/init_page.dart';
import 'package:paper_tracker/pages/learning_page.dart';
import 'package:paper_tracker/pages/main_page.dart';
import 'package:paper_tracker/pages/room_page.dart';
import 'package:paper_tracker/pages/server_config_page.dart';
import 'package:paper_tracker/pages/start_exec_page.dart';
import 'package:paper_tracker/pages/tracker_page.dart';
import 'package:paper_tracker/pages/tutorial_page.dart';
import 'package:paper_tracker/pages/workflow_exec_page.dart';
import 'package:paper_tracker/pages/workflow_template_page.dart';

void main() {
  WidgetsFlutterBinding.ensureInitialized();
  SystemChrome.setPreferredOrientations([DeviceOrientation.portraitUp]).then((_) => runApp(PaperTrackerApp()));
}

class PaperTrackerApp extends StatelessWidget {
  @override
  Widget build(BuildContext context) {
    var darkBackground = Color.fromRGBO(56, 66, 86, 1.0);
    var greyBackground = Color.fromRGBO(84, 93, 110, .9);
    var purple = Color.fromRGBO(148, 0, 238, 1.0);

    var theme = ThemeData.dark().copyWith(
      backgroundColor: darkBackground,
      scaffoldBackgroundColor: darkBackground,
      cardColor: greyBackground,
      accentColor: purple,
      primaryColor: darkBackground,
      indicatorColor: purple,
      textSelectionHandleColor: purple,
      floatingActionButtonTheme: ThemeData.dark().floatingActionButtonTheme.copyWith(
            backgroundColor: purple,
          ),
      iconTheme: ThemeData.dark().iconTheme.copyWith(
            color: purple,
          ),
      dialogBackgroundColor: darkBackground,
    );

    return MaterialApp(
      title: "Paper Tracker",
      theme: theme,
      initialRoute: "/",
      routes: {
        InitPage.Route: (context) => InitPage(),
        MainPage.Route: (context) => MainPage(),
        ServerConfigPage.Route: (context) => ServerConfigPage(),
        RoomPage.Route: (context) => RoomPage(),
        TrackerPage.Route: (context) => TrackerPage(),
        LearningPage.Route: (context) => LearningPage(),
        WorkflowTemplatePage.Route: (context) => WorkflowTemplatePage(),
        WorkflowExecPage.Route: (context) => WorkflowExecPage(),
        StartExecPage.Route: (context) => StartExecPage(),
        TutorialPage.Route: (context) => TutorialPage(),
        ConfigPage.Route: (context) => ConfigPage(),
      },
    );
  }
}
