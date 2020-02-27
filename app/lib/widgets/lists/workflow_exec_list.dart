import 'package:flutter/material.dart';
import 'package:paper_tracker/client/workflow_exec_client.dart';
import 'package:paper_tracker/model/workflow.dart';
import 'package:paper_tracker/pages/start_exec_page.dart';
import 'package:paper_tracker/widgets/lists/card_list.dart';

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
          var dataList = execList.map((exec) => CardListData(exec.label, exec.compeleted.toString(), exec)).toList();

          return Scaffold(
            body: CardList<WorkflowExec>(
              dataList: dataList,
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
