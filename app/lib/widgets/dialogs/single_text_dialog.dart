import 'package:flutter/material.dart';
import 'package:paper_tracker/widgets/label.dart';

class SingleTextDialog extends StatelessWidget {
  final TextEditingController labelController;
  final void Function() onButton;
  final String title;
  final String textLabel;
  final String buttonLabel;

  const SingleTextDialog(
      {Key key,
      @required this.labelController,
      @required this.onButton,
      @required this.title,
      @required this.textLabel,
      @required this.buttonLabel})
      : super(key: key);

  @override
  Widget build(BuildContext context) {
    return AlertDialog(
      content: Column(
        mainAxisSize: MainAxisSize.min,
        children: [
          Label(title),
          Padding(
            padding: EdgeInsets.only(top: 20.0),
          ),
          TextFormField(
            controller: labelController,
            autofocus: true,
            decoration: InputDecoration(
              labelText: textLabel,
              enabledBorder: OutlineInputBorder(borderSide: BorderSide(color: Theme.of(context).accentColor)),
              focusedBorder: OutlineInputBorder(borderSide: BorderSide(color: Theme.of(context).accentColor)),
            ),
          ),
        ],
      ),
      actions: [
        FlatButton(
          child: Text(buttonLabel),
          onPressed: onButton,
        ),
      ],
    );
  }
}
