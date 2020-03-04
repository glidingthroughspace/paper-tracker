import 'package:flutter/material.dart';

class ConfirmIconTextDialog extends StatelessWidget {
  final String text;
  final IconData icon;
  final Map<String, void Function()> actions;

  const ConfirmIconTextDialog({Key key, this.text = "Confirm this", this.icon = Icons.adb, @required this.actions})
      : super(key: key);

  @override
  Widget build(BuildContext context) {
    var buttons = actions
        .map((title, onPressed) => MapEntry(FlatButton(child: Text(title), onPressed: onPressed), null))
        .keys
        .toList();

    return AlertDialog(
      content: Row(children: [
        Icon(icon),
        Padding(padding: EdgeInsets.only(left: 10.0)),
        Text(text),
      ]),
      actions: buttons,
    );
  }
}
