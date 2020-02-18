import 'package:flutter/material.dart';
import 'package:paper_tracker/client/room_client.dart';
import 'package:paper_tracker/client/workflow_client.dart';
import 'package:paper_tracker/model/workflow.dart';
import 'package:paper_tracker/widgets/card_list.dart';
import 'package:paper_tracker/widgets/detail_content.dart';

class WorkflowPage extends StatefulWidget {
  static const Route = "/workflow";

  @override
  _WorkflowPageState createState() => _WorkflowPageState();
}

class _WorkflowPageState extends State<WorkflowPage> {
  var workflowClient = WorkflowClient();
  var roomClient = RoomClient();

  @override
  Widget build(BuildContext context) {
    var workflowID = ModalRoute.of(context).settings.arguments;
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
      ),
    );
  }
}
