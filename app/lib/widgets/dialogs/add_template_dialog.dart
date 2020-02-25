import 'package:flutter/material.dart';
import 'package:paper_tracker/widgets/label.dart';

class AddTemplateDialog extends StatelessWidget {
  final TextEditingController labelController;
  final void Function() addTemplate;

  const AddTemplateDialog({Key key, @required this.labelController, @required this.addTemplate}) : super(key: key);

  @override
  Widget build(BuildContext context) {
    return AlertDialog(
      content: Column(
        mainAxisSize: MainAxisSize.min,
        children: [
          Label("Add Workflow Template"),
          Padding(
            padding: EdgeInsets.only(top: 10.0),
          ),
          TextFormField(
            controller: labelController,
            autofocus: true,
            decoration: InputDecoration(
              labelText: "Workflow Template Label",
              enabledBorder: OutlineInputBorder(borderSide: BorderSide(color: Theme.of(context).accentColor)),
              focusedBorder: OutlineInputBorder(borderSide: BorderSide(color: Theme.of(context).accentColor)),
            ),
          ),
        ],
      ),
      actions: [
        FlatButton(
          child: Text("Create"),
          onPressed: addTemplate,
        ),
      ],
    );
  }
}
