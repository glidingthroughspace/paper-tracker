import 'package:flutter/material.dart';
import 'package:paper_tracker/client/api_client.dart';
import 'package:paper_tracker/config.dart';

class ConfigPage extends StatefulWidget {
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
      return showDialog(
          context: context,
          builder: (context) {
            return AlertDialog(
              content: Text("Server available"),
            );
          });
    } else {
      return showDialog(
          context: context,
          builder: (context) {
            return AlertDialog(
              content: Text("Server not available"),
            );
          });
    }
  }

  @override
  void initState() {
    config.getServerURL().then((url) => {
          setState(() {
            urlEditController.text = url;
          })
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
        padding: const EdgeInsets.all(30.0),
        child: Column(children: [
          Padding(
            padding: EdgeInsets.only(top: 140),
          ),
          Text(
            "Paper Tracker Config",
            style: TextStyle(color: Colors.deepOrange, fontSize: 30.0),
          ),
          Padding(
            padding: EdgeInsets.only(top: 50.0),
          ),
          TextFormField(
            controller: urlEditController,
            decoration: InputDecoration(
              labelText: "Server URL",
              border: OutlineInputBorder(),
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
