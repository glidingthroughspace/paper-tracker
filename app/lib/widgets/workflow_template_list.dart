import 'package:flutter/material.dart';
import 'package:paper_tracker/client/workflow_template_client.dart';
import 'package:paper_tracker/model/workflow.dart';
import 'package:paper_tracker/pages/workflow_page.dart';
import 'package:tuple/tuple.dart';

import 'card_list.dart';
import 'label.dart';

class WorkflowTemplateList extends StatefulWidget {
  @override
  _WorkflowTemplateListState createState() => _WorkflowTemplateListState();
}

class _WorkflowTemplateListState extends State<WorkflowTemplateList> with AutomaticKeepAliveClientMixin {
  var templateClient = WorkflowTemplateClient();
  var templateLabelEditController = TextEditingController();

  @override
  Widget build(BuildContext context) {
    super.build(context);
    return FutureBuilder(
      future: templateClient.getAllTemplates(),
      builder: (context, snapshot) {
        if (snapshot.hasData) {
          List<WorkflowTemplate> templateList = snapshot.data;
          List<Tuple2<String, WorkflowTemplate>> titleObjectList =
              templateList.map((workflow) => Tuple2(workflow.label, workflow)).toList();

          return Scaffold(
            body: CardList<WorkflowTemplate>(
              titleObjectList: titleObjectList,
              onTap: onTapWorkflow,
              iconData: Icons.keyboard_arrow_right,
              onRefresh: onRefresh,
            ),
            floatingActionButton: FloatingActionButton(
              onPressed: onAddTemplateButton,
              child: Icon(Icons.add),
              heroTag: "templateAddButton",
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
      templateClient.getAllTemplates(refresh: true);
    });
  }

  void onAddTemplateButton() async {
    return showDialog(
      context: context,
      builder: buildAddTemplateDialog,
    );
  }

  Widget buildAddTemplateDialog(BuildContext context) {
    return AlertDialog(
      content: Column(
        mainAxisSize: MainAxisSize.min,
        children: [
          Label("Add Workflow Template"),
          Padding(
            padding: EdgeInsets.only(top: 10.0),
          ),
          TextFormField(
            controller: templateLabelEditController,
            autofocus: true,
            decoration: InputDecoration(
              labelText: "Workflow Template Label",
              enabledBorder: OutlineInputBorder(borderSide: BorderSide(color: Theme.of(context).accentColor)),
              focusedBorder: OutlineInputBorder(borderSide: BorderSide(color: Theme.of(context).accentColor)),
            ),
          ),
        ],
      ),
      actions: [
        FlatButton(
          child: Text("Create"),
          onPressed: () => addTemplate(),
        ),
      ],
    );
  }

  void addTemplate() async {
    var template = WorkflowTemplate(label: templateLabelEditController.text);
    await templateClient.createTemplate(template);
    await templateClient.getAllTemplates(refresh: true);

    setState(() {});
    Navigator.of(context).pop();
  }

  @override
  bool get wantKeepAlive => true;

  void onTapWorkflow(WorkflowTemplate workflow) async {
    await Navigator.of(context).pushNamed(WorkflowTemplatePage.Route, arguments: workflow.id);
  }
}
