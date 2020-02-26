import 'package:flutter/material.dart';

class ConfirmIconTextDialog extends StatelessWidget {
  final String text;
  final IconData icon;

  const ConfirmIconTextDialog({Key key, this.text = "Confirm this", this.icon = Icons.adb}) : super(key: key);

  @override
  Widget build(BuildContext context) {
    return AlertDialog(
      content: Row(children: [
        Icon(icon),
        Padding(padding: EdgeInsets.only(left: 10.0)),
        Text(text),
      ]),
      actions: <Widget>[
        FlatButton(
          child: Text("OK"),
          onPressed: () => Navigator.of(context).pop(),
        )
      ],
    );
  }
}
