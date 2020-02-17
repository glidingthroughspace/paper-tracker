import 'package:flutter/material.dart';
import 'package:paper_tracker/client/workflow_client.dart';
import 'package:paper_tracker/model/workflow.dart';
import 'package:tuple/tuple.dart';

import 'card_list.dart';

class WorkflowList extends StatefulWidget {
  @override
  _WorkflowListState createState() => _WorkflowListState();
}

class _WorkflowListState extends State<WorkflowList> with AutomaticKeepAliveClientMixin {
  var workflowClient = WorkflowClient();

  @override
  Widget build(BuildContext context) {
    super.build(context);
    return FutureBuilder(
      future: workflowClient.getAllWorkflows(),
      builder: (context, snapshot) {
        if (snapshot.hasData) {
          List<Workflow> workflowList = snapshot.data;
          List<Tuple2<String, Workflow>> titleObjectList =
              workflowList.map((workflow) => Tuple2(workflow.label, workflow)).toList();

          return Scaffold(
            body: CardList<Workflow>(
              titleObjectList: titleObjectList,
              onTap: null,
              iconData: Icons.keyboard_arrow_right,
            ),
            floatingActionButton: FloatingActionButton(
              onPressed: null,
              child: Icon(Icons.add),
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

  @override
  bool get wantKeepAlive => true;
}
