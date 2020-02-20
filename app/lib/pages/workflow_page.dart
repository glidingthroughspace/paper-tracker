import 'package:flutter/material.dart';
import 'package:paper_tracker/client/room_client.dart';
import 'package:paper_tracker/client/workflow_client.dart';
import 'package:paper_tracker/model/communication/createStepRequest.dart';
import 'package:paper_tracker/model/workflow.dart';
import 'package:paper_tracker/widgets/card_list.dart';
import 'package:paper_tracker/widgets/detail_content.dart';
import 'package:paper_tracker/widgets/room_dropdown.dart';

class WorkflowPage extends StatefulWidget {
  static const Route = "/workflow";

  @override
  _WorkflowPageState createState() => _WorkflowPageState();
}

class _WorkflowPageState extends State<WorkflowPage> {
  var workflowClient = WorkflowClient();
  var roomClient = RoomClient();

  int workflowID;
  var stepLabelEditController = TextEditingController();
  var stepDecisionLabelEditController = TextEditingController();
  var roomDropdownController = RoomDropdownController();

  @override
  Widget build(BuildContext context) {
    workflowID = ModalRoute.of(context).settings.arguments;
    var futureWorkflow = workflowClient.getWorkflowByID(workflowID);

    return FutureBuilder(
      future: futureWorkflow,
      builder: (context, snapshot) {
        Workflow workflow = snapshot.data;

        return DetailContent(
          title: workflow != null ? workflow.label : "",
          iconData: Workflow.IconData,
          bottomButtons: [],
          content: workflow != null ? buildContent(workflow) : Container(),
        );
      },
    );
  }

  Widget buildContent(Workflow workflow) {
    return Container(
      padding: EdgeInsets.all(15.0),
      child: WorkflowStepsList(
        steps: workflow.steps,
        roomClient: roomClient,
        onStepAdd: onAddStep,
      ),
    );
  }

  void onAddStep(WFStep prev) {
    showDialog(
      context: context,
      child: buildAddStepDialog(prev),
    );
  }

  Widget buildAddStepDialog(WFStep step) {
    var children = [
      Text(
        "Add Step",
        style: TextStyle(
          fontSize: 20.0,
        ),
      ),
      Padding(padding: EdgeInsets.only(top: 10.0)),
      TextFormField(
        controller: stepLabelEditController,
        decoration: InputDecoration(
          labelText: "Step Label",
          enabledBorder: OutlineInputBorder(borderSide: BorderSide(color: Theme.of(context).accentColor)),
          focusedBorder: OutlineInputBorder(borderSide: BorderSide(color: Theme.of(context).accentColor)),
        ),
      ),
    ];

    if (step.options.length < 2) {
      children.addAll([
        Padding(padding: EdgeInsets.only(top: 10.0)),
        TextFormField(
          controller: stepDecisionLabelEditController,
          decoration: InputDecoration(
            labelText: "Decision Label",
            enabledBorder: OutlineInputBorder(borderSide: BorderSide(color: Theme.of(context).accentColor)),
            focusedBorder: OutlineInputBorder(borderSide: BorderSide(color: Theme.of(context).accentColor)),
          ),
        )
      ]);
    }

    children.addAll([
      Padding(padding: EdgeInsets.only(top: 10.0)),
      RoomDropdown(
        roomClient: roomClient,
        controller: roomDropdownController,
      ),
    ]);

    return AlertDialog(
      content: Column(
        mainAxisSize: MainAxisSize.min,
        children: children,
      ),
      actions: [
        FlatButton(
          child: Text("Add"),
          onPressed: () => addStep(step),
        ),
      ],
    );
  }

  void addStep(WFStep prevStep) async {
    var createStepRequest = CreateStepRequest(
      decisionLabel: stepDecisionLabelEditController.text,
      previousStepID: prevStep.id,
      step: WFStep(
        label: stepLabelEditController.text,
        roomID: roomDropdownController.selectedRoom.id,
      ),
    );
    await workflowClient.addStep(workflowID, createStepRequest);
    await workflowClient.getAllWorkflows(refresh: true);

    setState(() {});
    Navigator.of(context).pop();
  }
}
