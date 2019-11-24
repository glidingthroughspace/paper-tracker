import 'package:flutter/material.dart';
import 'package:paper_tracker/client/api_client.dart';

class InitPage extends StatefulWidget {
  static const Route = "/";

  @override
  _InitPageState createState() => _InitPageState();
}

class _InitPageState extends State<InitPage> {
  @override
  void initState() {
    super.initState();
    APIClient().isAvailable().then((serverAvailable) {
      if (serverAvailable)
        Navigator.pushReplacementNamed(context, "/main");
      else
        Navigator.pushReplacementNamed(context, "/config");
    });
  }

  @override
  Widget build(BuildContext context) {
    return Scaffold(body: Center(child: CircularProgressIndicator()));
  }
}