import 'package:flutter/material.dart';
import 'package:paper_tracker/client/api_client.dart';
import 'package:paper_tracker/pages/main_page.dart';
import 'package:paper_tracker/pages/server_config_page.dart';

class InitPage extends StatefulWidget {
  static const Route = "/";

  @override
  _InitPageState createState() => _InitPageState();
}

class _InitPageState extends State<InitPage> {
  @override
  void initState() {
    super.initState();
    APIClient().isAvailable().timeout(Duration(seconds: 2)).then((serverAvailable) {
      if (serverAvailable)
        Navigator.pushReplacementNamed(context, MainPage.Route);
      else
        Navigator.pushReplacementNamed(context, ServerConfigPage.Route);
    }).catchError((error) {
      Navigator.pushReplacementNamed(context, ServerConfigPage.Route);
    });
  }

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      body: Center(
        child: CircularProgressIndicator(),
      ),
    );
  }
}
