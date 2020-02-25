import 'package:flutter/material.dart';
import 'package:paper_tracker/client/workflow_exec_client.dart';
import 'package:paper_tracker/model/workflow.dart';
import 'package:paper_tracker/pages/start_exec_page.dart';
import 'package:tuple/tuple.dart';

import '../card_list.dart';

class WorkflowExecList extends StatefulWidget {
  @override
  _WorkflowExecListState createState() => _WorkflowExecListState();
}

class _WorkflowExecListState extends State<WorkflowExecList> {
  var execClient = WorkflowExecClient();

  @override
  Widget build(BuildContext context) {
    return FutureBuilder(
      future: execClient.getAllExecs(),
      builder: (context, snapshot) {
        if (snapshot.hasData) {
          List<WorkflowExec> execList = snapshot.data;
          List<Tuple2<String, WorkflowExec>> titleObjectList =
              execList.map((exec) => Tuple2(exec.label, exec)).toList();

          return Scaffold(
            body: CardList<WorkflowExec>(
              titleObjectList: titleObjectList,
              onTap: onTapExec,
              iconData: Icons.keyboard_arrow_right,
              onRefresh: onRefresh,
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
    setState(() {
      execClient.getAllExecs(refresh: true);
    });
  }

  void onTapExec(WorkflowExec exec) async {
    //await Navigator.of(context).pushNamed(WorkflowTemplatePage.Route, arguments: workflow.id);
  }

  void onStartExec() async {
    await Navigator.of(context).pushNamed(StartExecPage.Route);
  }
}
