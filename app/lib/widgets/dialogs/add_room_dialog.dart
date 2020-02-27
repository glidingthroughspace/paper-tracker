import 'package:flutter/material.dart';
import 'package:paper_tracker/widgets/label.dart';

class AddRoomDialog extends StatelessWidget {
  final TextEditingController labelController;
  final void Function() addRoom;

  const AddRoomDialog({Key key, @required this.labelController, @required this.addRoom}) : super(key: key);

  @override
  Widget build(BuildContext context) {
    return AlertDialog(
      content: Column(
        mainAxisSize: MainAxisSize.min,
        children: [
          Label("Add Room"),
          Padding(
            padding: EdgeInsets.only(top: 10.0),
          ),
          TextFormField(
            controller: labelController,
            autofocus: true,
            decoration: InputDecoration(
              labelText: "Room Label",
              enabledBorder: OutlineInputBorder(borderSide: BorderSide(color: Theme.of(context).accentColor)),
              focusedBorder: OutlineInputBorder(borderSide: BorderSide(color: Theme.of(context).accentColor)),
            ),
          ),
        ],
      ),
      actions: [
        FlatButton(
          child: Text("Create"),
          onPressed: addRoom,
        ),
      ],
    );
  }
}
