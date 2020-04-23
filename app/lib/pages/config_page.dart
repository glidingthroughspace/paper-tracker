import 'package:flutter/material.dart';
import 'package:fluttertoast/fluttertoast.dart';
import 'package:paper_tracker/client/api_client.dart';
import 'package:paper_tracker/client/config_client.dart';
import 'package:paper_tracker/utils.dart';
import 'package:url_launcher/url_launcher.dart';

class ConfigPage extends StatefulWidget {
  static const Route = "/config";

  @override
  _ConfigPageState createState() => _ConfigPageState();
}

class _ConfigPageState extends State<ConfigPage> {
  var configClient = ConfigClient();
  Future<Map<String, dynamic>> configFuture;
  Map<String, TextEditingController> textController;

  @override
  void initState() {
    fetchConfig();
    super.initState();
  }

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      appBar: AppBar(
        title: Text("Settings"),
        centerTitle: true,
        leading: IconButton(
          icon: Icon(Icons.arrow_back),
          onPressed: () => Navigator.of(context).pop(),
        ),
      ),
      bottomNavigationBar: buildBottomNavigation(context),
      body: SingleChildScrollView(
        child: buildBody(context),
      ),
    );
  }

  Widget buildBody(BuildContext context) {
    return FutureBuilder(
      future: configFuture,
      builder: (context, snapshot) {
        if (snapshot.hasData) {
          Map<String, dynamic> configMap = snapshot.data;
          return buildSettingsList(context, configMap);
        }
        return CircularProgressIndicator();
      },
    );
  }

  Widget buildSettingsList(BuildContext context, Map<String, dynamic> configMap) {
    var tableRows = configMap
        .map((key, val) => MapEntry(
            TableRow(children: [
              TableCell(child: Text("$key: ")),
              TextFormField(controller: textController[key]),
            ]),
            null))
        .keys
        .toList();

    tableRows.add(getTableSpacing(30.0));
    tableRows.add(TableRow(children: [
      TableCell(child: Text("Data Export: ")),
      TableCell(child: FlatButton(child: Text("Download"), onPressed: downloadExport)),
    ]));

    return Container(
      padding: EdgeInsets.all(15.0),
      child: Table(
        defaultVerticalAlignment: TableCellVerticalAlignment.middle,
        columnWidths: {0: FractionColumnWidth(0.4)},
        children: tableRows,
      ),
    );
  }

  Container buildBottomNavigation(BuildContext context) {
    return Container(
      height: 55.0,
      child: BottomAppBar(
        color: Theme.of(context).cardColor,
        child: Container(
          padding: EdgeInsets.symmetric(horizontal: 10.0),
          child: Row(
            mainAxisAlignment: MainAxisAlignment.spaceAround,
            children: [
              IconButton(
                icon: Icon(Icons.settings_backup_restore, color: Colors.white),
                onPressed: () => setState(() {
                  fetchConfig();
                }),
              ),
              IconButton(
                icon: Icon(Icons.save, color: Colors.white),
                onPressed: () => saveConfig(),
              )
            ],
          ),
        ),
      ),
    );
  }

  void saveConfig() async {
    var oldConfig = await configFuture;
    var newConfig = Map<String, dynamic>();
    for (var entry in oldConfig.entries) {
      var key = entry.key;
      var oldVal = entry.value;

      var input = textController[key].text;
      if (oldVal is int) {
        newConfig[key] = int.parse(input);
      } else if (oldVal is bool) {
        newConfig[key] = input.toLowerCase() == "true" ? true : false;
      } else if (oldVal is List) {
        // Only supports string lists for now
        input = input.replaceAll("[", "");
        input = input.replaceAll("]", "");
        newConfig[key] = input.split(",").map((val) => val.trim()).toList();
      } else {
        Fluttertoast.showToast(msg: "Unsupported config type: ${oldVal.runtimeType}");
        return;
      }
    }
    try {
      configClient.setConfig(newConfig);
      setState(() {
        fetchConfig();
      });
    } catch (ex) {
      Fluttertoast.showToast(msg: "Failed to set config: $ex");
    }
  }

  void fetchConfig() {
    configFuture = configClient.getConfig().then((configMap) {
      textController = configMap.map((key, val) => MapEntry(key, TextEditingController(text: val.toString())));
      return configMap;
    });
  }

  void downloadExport() async {
    var url = await APIClient().buildURI("/export.xlsx", null);
    if (await canLaunch(url.toString())) {
      await launch(url.toString());
    } else {
      Fluttertoast.showToast(msg: "Failed to open url: ${url.toString()}");
    }
  }
}
