import 'package:flutter/material.dart';
import 'package:intl/intl.dart';
import 'package:paper_tracker/client/workflow_exec_client.dart';
import 'package:paper_tracker/client/workflow_template_client.dart';
import 'package:paper_tracker/model/workflow.dart';
import 'package:paper_tracker/pages/start_exec_page.dart';
import 'package:paper_tracker/pages/workflow_exec_page.dart';
import 'package:paper_tracker/widgets/lists/card_list.dart';

class WorkflowExecList extends StatefulWidget {
  @override
  _WorkflowExecListState createState() => _WorkflowExecListState();
}

class _WorkflowExecListState extends State<WorkflowExecList> {
  var execClient = WorkflowExecClient();
  var templateClient = WorkflowTemplateClient();

  @override
  Widget build(BuildContext context) {
    return FutureBuilder(
      future: execClient.getAllExecs(),
      builder: (context, snapshot) {
        if (snapshot.hasData) {
          List<WorkflowExec> execList = snapshot.data;
          execList.sort((a, b) => a.startedOn.compareTo(b.startedOn));
          execList = execList.reversed.toList();
          var dataList = execList
              .map((exec) => CardListData(exec.label, buildSubtitle(exec), exec,
                  color: exec.status == WorkflowExecStatus.Finished ? WorkflowExec.CompletedColor : null))
              .toList();

          return Scaffold(
            body: CardList<WorkflowExec>(
              dataList: dataList,
              onTap: onTapExec,
              iconData: Icons.keyboard_arrow_right,
              onRefresh: onRefresh,
              subtitleColum: true,
            ),
            floatingActionButton: FloatingActionButton(
              onPressed: onStartExec,
              child: Icon(Icons.add),
              heroTag: "execAddButton",
            ),
          );
        } else if (snapshot.hasError) {
          return Center(child: Text("${snapshot.error}"));
        }

        // By default, show a loading spinner.
        return Center(child: CircularProgressIndicator());
      },
    );
  }

  Future<void> onRefresh() async {
    await execClient.getAllExecs(refresh: true);
    setState(() {});
  }

  void onTapExec(WorkflowExec exec) async {
    await Navigator.of(context).pushNamed(WorkflowExecPage.Route, arguments: exec.id);
  }

  void onStartExec() async {
    await Navigator.of(context).pushNamed(StartExecPage.Route);
  }

  List<Widget> buildSubtitle(WorkflowExec exec) {
    var dateFormatter = DateFormat("dd.MM.yyyy HH:mm");

    var subtitle = <Widget>[
      Text("Started on: ${dateFormatter.format(exec.startedOn.toLocal())}"),
    ];

    if (exec.status == WorkflowExecStatus.Finished) {
      subtitle.add(Text("Completed on: ${dateFormatter.format(exec.completedOn.toLocal())}"));
    } else {
      subtitle.add(
        FutureBuilder(
          future: templateClient.getStepByID(exec.templateID, exec.currentStepID),
          builder: (context, snapshot) {
            if (snapshot.hasData) {
              WFStep step = snapshot.data;
              return Text("Current Step: ${step.label}");
            }
            return Text("Current Step: ");
          },
        ),
      );
    }

    return subtitle;
  }
}
