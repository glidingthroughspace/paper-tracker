import 'package:flutter/material.dart';
import 'package:paper_tracker/client/api_client.dart';
import 'package:paper_tracker/config.dart';
import 'package:paper_tracker/pages/main_page.dart';

class ConfigPage extends StatefulWidget {
  static const Route = "/config";

  @override
  _ConfigPageState createState() => _ConfigPageState();
}

class _ConfigPageState extends State<ConfigPage> {
  final urlEditController = TextEditingController();
  final config = Config();

  void onSubmit() async {
    await config.setServerURL(urlEditController.text);
    var serverAvailable = await APIClient().isAvailable();
    if (serverAvailable) {
      Navigator.of(context).pushReplacementNamed(MainPage.Route);
    } else {
      return showDialog(
          context: context,
          builder: (context) {
            return AlertDialog(
              content: Row(children: [
                Icon(Icons.warning),
                Padding(padding: EdgeInsets.only(left: 10.0)),
                Text("Server not available"),
              ]),
              actions: <Widget>[
                FlatButton(
                  child: Text("OK"),
                  onPressed: () => Navigator.of(context).pop(),
                )
              ],
            );
          });
    }
  }

  @override
  void initState() {
    config.getServerURL().then((url) {
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
          )
        ]),
      ),
    );
  }
}
