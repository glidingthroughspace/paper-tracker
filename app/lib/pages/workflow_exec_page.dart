import 'package:flutter/material.dart';
import 'package:paper_tracker/client/room_client.dart';
import 'package:paper_tracker/client/workflow_exec_client.dart';
import 'package:paper_tracker/client/workflow_template_client.dart';
import 'package:paper_tracker/model/workflow.dart';
import 'package:paper_tracker/widgets/detail_content.dart';
import 'package:paper_tracker/widgets/dialogs/workflow_step_dialog.dart';
import 'package:paper_tracker/widgets/lists/workflow_steps_list.dart';

class WorkflowExecPage extends StatefulWidget {
  static const Route = "/workflow/exec";

  @override
  _WorkflowExecPageState createState() => _WorkflowExecPageState();
}

class _WorkflowExecPageState extends State<WorkflowExecPage> {
  var execClient = WorkflowExecClient();
  var templateClient = WorkflowTemplateClient();
  var roomClient = RoomClient();
  int execID;
  Future<WorkflowExec> futureExec;

  @override
  void didChangeDependencies() {
    super.didChangeDependencies();
    execID = ModalRoute.of(context).settings.arguments;
    futureExec = execClient.getExecByID(execID);
  }

  @override
  Widget build(BuildContext context) {
    return FutureBuilder(
      future: futureExec,
      builder: (context, snapshot) {
        WorkflowExec exec = snapshot.data;

        return DetailContent(
          title: exec != null ? exec.label : "",
          iconData: WorkflowExec.IconData,
          bottomButtons: [],
          content: exec != null ? buildContent(exec) : Container(),
          onRefresh: refreshExec,
        );
      },
    );
  }

  Widget buildContent(WorkflowExec exec) {
    var futureTemplate = templateClient.getTemplateByID(exec.templateID);
    print(exec.currentStepID);

    return FutureBuilder(
      future: futureTemplate,
      builder: (context, snapshot) {
        if (snapshot.hasData) {
          WorkflowTemplate template = snapshot.data;
          return Container(
            padding: EdgeInsets.all(15.0),
            child: WorkflowStepsList(
              steps: template.steps,
              roomClient: roomClient,
              primaryScroll: false,
              stepInfos: exec.stepInfos,
              currentStep: exec.currentStepID,
              onTap: onStepTap,
            ),
          );
        }
        return Container();
      },
    );
  }

  void onStepTap(WFStep step) async {
    var exec = await futureExec;
    var options = Map<String, void Function(WFStep)>();
    if (exec.currentStepID == step.id) {
      options["Skip this step"] = onProgressToStep;
    } else {
      options["Set workflow to this step"] = onProgressToStep;
    }

    showDialog(
      context: context,
      child: OptionsDialog(object: step, options: options),
    );
  }

  void onProgressToStep(WFStep step) async {
    await execClient.progressToStep(execID, step.id);
    refreshExec();
    Navigator.of(context).pop();
  }

  Future<void> refreshExec() async {
    setState(() {
      futureExec = execClient.getExecByID(execID, refresh: true);
    });
  }
}
