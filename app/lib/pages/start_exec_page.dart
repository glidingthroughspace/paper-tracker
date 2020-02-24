import 'package:flutter/material.dart';
import 'package:paper_tracker/client/room_client.dart';
import 'package:paper_tracker/client/tracker_client.dart';
import 'package:paper_tracker/client/workflow_exec_client.dart';
import 'package:paper_tracker/client/workflow_template_client.dart';
import 'package:paper_tracker/model/tracker.dart';
import 'package:paper_tracker/model/workflow.dart';
import 'package:paper_tracker/widgets/card_list.dart';
import 'package:paper_tracker/widgets/detail_content.dart';
import 'package:paper_tracker/widgets/dropdown.dart';

class StartExecPage extends StatefulWidget {
  static const String Route = "/workflow/exec/start";

  @override
  _StartExecPageState createState() => _StartExecPageState();
}

class _StartExecPageState extends State<StartExecPage> {
  var roomClient = RoomClient();
  var trackerClient = TrackerClient();
  var templateClient = WorkflowTemplateClient();
  var execClient = WorkflowExecClient();
  var trackerDropdownController = DropdownController();
  var templateDropdownController = DropdownController();
  var workflowStepsListController = WorkflowStepsListController();
  var execLabelEditController = TextEditingController();

  @override
  Widget build(BuildContext context) {
    return DetailContent(
      title: "Start Workflow",
      iconData: WorkflowTemplate.IconData,
      content: buildContent(),
      bottomButtons: [
        IconButton(
          icon: Icon(Icons.play_circle_outline, color: Colors.white),
          onPressed: trackerDropdownController.selectedItem != null && templateDropdownController.selectedItem != null
              ? onStartExec
              : null,
        ),
      ],
    );
  }

  Widget buildContent() {
    var children = [
      TextFormField(
        controller: execLabelEditController,
        autofocus: true,
        decoration: InputDecoration(
          labelText: "Label",
          enabledBorder: OutlineInputBorder(borderSide: BorderSide(color: Theme.of(context).accentColor)),
          focusedBorder: OutlineInputBorder(borderSide: BorderSide(color: Theme.of(context).accentColor)),
        ),
      ),
      Padding(padding: EdgeInsets.only(top: 20.0)),
      Dropdown(
        controller: trackerDropdownController,
        getItems: () async {
          var allTrackers = await trackerClient.getAllTrackers();
          return allTrackers.where((tracker) => tracker.status == TrackerStatus.Idle).toList();
        },
        hintName: "available tracker",
        icon: Tracker.IconData,
        setState: setState,
      ),
      Padding(padding: EdgeInsets.only(top: 5.0)),
      Dropdown(
        controller: templateDropdownController,
        getItems: templateClient.getAllTemplates,
        hintName: "workflow",
        icon: WorkflowTemplate.IconData,
        setState: setState,
      )
    ];

    if (templateDropdownController.selectedItem != null) {
      WorkflowTemplate template = templateDropdownController.selectedItem;
      children.add(
        WorkflowStepsList(
          roomClient: roomClient,
          steps: template.steps,
          controller: workflowStepsListController,
          primaryScroll: false,
        ),
      );
    }

    return Container(
      padding: EdgeInsets.all(15.0),
      child: Column(
        children: children,
      ),
    );
  }

  void onStartExec() async {
    var infos = workflowStepsListController.selectedDecisionMap
        .map((id, decision) => MapEntry(id, ExecStepInfo(decision: decision)));

    var exec = WorkflowExec(
      label: execLabelEditController.text,
      trackerID: trackerDropdownController.selectedItem.id,
      templateID: templateDropdownController.selectedItem.id,
      stepInfos: infos,
    );
    await execClient.startExec(exec);
    await execClient.getAllExecs(refresh: true);
    Navigator.of(context).pop();
  }
}
