import 'package:flutter/material.dart';
import 'package:paper_tracker/client/room_client.dart';
import 'package:paper_tracker/model/room.dart';
import 'package:paper_tracker/model/workflow.dart';
import 'package:paper_tracker/widgets/dropdown.dart';

class EditStepDialog extends StatelessWidget {
  final WFStep step;
  final TextEditingController labelController;
  final List<TextEditingController> decisionController;
  final DropdownController roomDropdownController;
  final RoomClient roomClient;
  final void Function(WFStep) editStep;

  const EditStepDialog(
      {Key key,
      @required this.step,
      @required this.labelController,
      @required this.decisionController,
      @required this.roomClient,
      @required this.editStep,
      @required this.roomDropdownController})
      : super(key: key);

  @override
  Widget build(BuildContext context) {
    var children = [
      Text(
        "Edit Step",
        style: TextStyle(
          fontSize: 20.0,
        ),
      ),
      Padding(padding: EdgeInsets.only(top: 10.0)),
      TextFormField(
        controller: labelController,
        decoration: InputDecoration(
          labelText: "Step Label",
          enabledBorder: OutlineInputBorder(borderSide: BorderSide(color: Theme.of(context).accentColor)),
          focusedBorder: OutlineInputBorder(borderSide: BorderSide(color: Theme.of(context).accentColor)),
        ),
      ),
    ];

    for (int it = 0; it < decisionController.length; it++) {
      children.addAll([
        Padding(padding: EdgeInsets.only(top: 10.0)),
        TextFormField(
          controller: decisionController[it],
          decoration: InputDecoration(
            labelText: "Decision Label ${it + 1}",
            enabledBorder: OutlineInputBorder(borderSide: BorderSide(color: Theme.of(context).accentColor)),
            focusedBorder: OutlineInputBorder(borderSide: BorderSide(color: Theme.of(context).accentColor)),
          ),
        )
      ]);
    }

    children.addAll([
      Padding(padding: EdgeInsets.only(top: 10.0)),
      Dropdown(
        getItems: () async {
          var rooms = await roomClient.getAllRooms(refresh: true);
          return rooms.where((room) => room.isLearned).toList();
        },
        controller: roomDropdownController,
        hintName: "learned room",
        icon: Room.IconData,
      ),
    ]);

    return AlertDialog(
      content: Column(
        mainAxisSize: MainAxisSize.min,
        children: children,
      ),
      actions: [
        FlatButton(
          child: Text("Edit"),
          onPressed: () => editStep(step),
        ),
      ],
    );
  }
}
