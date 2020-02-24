import 'package:flutter/material.dart';
import 'package:paper_tracker/client/workflow_template_client.dart';
import 'package:paper_tracker/model/workflow.dart';
import 'package:paper_tracker/pages/workflow_page.dart';
import 'package:tuple/tuple.dart';

import 'card_list.dart';

class WorkflowList extends StatefulWidget {
  @override
  _WorkflowListState createState() => _WorkflowListState();
}

class _WorkflowListState extends State<WorkflowList> with AutomaticKeepAliveClientMixin {
  var workflowClient = WorkflowTemplateClient();

  @override
  Widget build(BuildContext context) {
    super.build(context);
    return FutureBuilder(
      future: workflowClient.getAllTemplates(),
      builder: (context, snapshot) {
        if (snapshot.hasData) {
          List<WorkflowTemplate> workflowList = snapshot.data;
          List<Tuple2<String, WorkflowTemplate>> titleObjectList =
              workflowList.map((workflow) => Tuple2(workflow.label, workflow)).toList();

          return Scaffold(
            body: CardList<WorkflowTemplate>(
              titleObjectList: titleObjectList,
              onTap: onTapWorkflow,
              iconData: Icons.keyboard_arrow_right,
              onRefresh: onRefresh,
            ),
            floatingActionButton: FloatingActionButton(
              onPressed: null,
              child: Icon(Icons.add),
              heroTag: "workflowAddButton",
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
      workflowClient.getAllTemplates(refresh: true);
    });
  }

  @override
  bool get wantKeepAlive => true;

  void onTapWorkflow(WorkflowTemplate workflow) async {
    await Navigator.of(context).pushNamed(WorkflowTemplatePage.Route, arguments: workflow.id);
  }
}
