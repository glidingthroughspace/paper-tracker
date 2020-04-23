import 'package:flutter/material.dart';
import 'package:paper_tracker/client/api_client.dart';
import 'package:paper_tracker/config.dart';
import 'package:paper_tracker/pages/main_page.dart';
import 'package:paper_tracker/pages/tutorial_page.dart';
import 'package:paper_tracker/widgets/dialogs/confirm_icon_text_dialog.dart';
import 'package:paper_tracker/widgets/dialogs/waiting_text_dialog.dart';

class ServerConfigPage extends StatefulWidget {
  static const Route = "/server/config";

  @override
  _ServerConfigPageState createState() => _ServerConfigPageState();
}

class _ServerConfigPageState extends State<ServerConfigPage> {
  final urlEditController = TextEditingController();
  final config = Config();

  void onSubmit() async {
    showDialog(context: context, child: WaitingTextDialog(text: "Connecting to server..."));

    await config.setServerURL(urlEditController.text);
    var serverAvailable = await APIClient().isAvailable();

    Navigator.of(context).pop();

    if (serverAvailable) {
      Navigator.of(context).pushReplacementNamed(MainPage.Route);
    } else {
      return showDialog(
        context: context,
        child: ConfirmIconTextDialog(
          text: "Server not available",
          icon: Icons.warning,
          actions: {"OK": () => Navigator.of(context).pop()},
        ),
      );
    }
  }

  @override
  void initState() {
    config.getServerURL().then((url) {
      if (url == null) {
        Navigator.of(context).pushNamed(TutorialPage.Route);
      }
      urlEditController.text = url;
    });
    super.initState();
  }

  @override
  void dispose() {
    urlEditController.dispose();
    super.dispose();
  }

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      persistentFooterButtons: [
        IconButton(
          icon: Icon(Icons.help),
          onPressed: () => Navigator.of(context).pushNamed(TutorialPage.Route),
        )
      ],
      body: Container(
        padding: EdgeInsets.all(30.0),
        child: Column(children: [
          Padding(
            padding: EdgeInsets.only(top: 140),
          ),
          Text(
            "Paper Tracker Config",
            style: TextStyle(color: Theme.of(context).accentColor, fontSize: 30.0),
          ),
          Padding(
            padding: EdgeInsets.only(top: 50.0),
          ),
          TextFormField(
            controller: urlEditController,
            decoration: InputDecoration(
              labelText: "Server URL",
              enabledBorder: OutlineInputBorder(borderSide: BorderSide(color: Theme.of(context).accentColor)),
              focusedBorder: OutlineInputBorder(borderSide: BorderSide(color: Theme.of(context).accentColor)),
              suffixIcon: IconButton(
                icon: Icon(Icons.clear),
                onPressed: () => urlEditController.clear(),
              ),
            ),
            validator: (val) {
              if (val.length == 0) {
                return "URL cannot be empty";
              } else {
                return null;
              }
            },
            keyboardType: TextInputType.url,
          ),
          Padding(
            padding: EdgeInsets.only(top: 25.0),
          ),
          MaterialButton(
            child: Text("Submit"),
            onPressed: onSubmit,
            color: Theme.of(context).accentColor,
          ),
        ]),
      ),
    );
  }
}
