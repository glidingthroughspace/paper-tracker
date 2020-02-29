import 'package:flutter/material.dart';
import 'package:paper_tracker/model/workflow.dart';

class WorkflowStepDialog extends StatelessWidget {
  final WFStep step;
  final void Function(WFStep) onEdit;
  final void Function(WFStep) onDelete;

  const WorkflowStepDialog({Key key, this.step, this.onEdit, this.onDelete}) : super(key: key);

  @override
  Widget build(BuildContext context) {
    List<Widget> buttons = [];
    if (onEdit != null) {
      buttons.add(MaterialButton(
        child: Text("Edit"),
        shape: Border.all(color: Colors.grey),
        onPressed: () => onEdit(step),
      ));
    }
    if (onDelete != null) {
      buttons.add(MaterialButton(
        child: Text("Delete"),
        shape: Border.all(color: Colors.grey),
        onPressed: () => onDelete(step),
      ));
    }

    return AlertDialog(
      content: Column(
        mainAxisSize: MainAxisSize.min,
        crossAxisAlignment: CrossAxisAlignment.stretch,
        children: buttons,
      ),
    );
  }
}
