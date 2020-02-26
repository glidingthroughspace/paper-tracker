import 'package:flutter/material.dart';
import 'package:paper_tracker/client/room_client.dart';
import 'package:paper_tracker/client/workflow_template_client.dart';
import 'package:paper_tracker/model/communication/createStepRequest.dart';
import 'package:paper_tracker/model/room.dart';
import 'package:paper_tracker/model/workflow.dart';
import 'package:paper_tracker/widgets/card_list.dart';
import 'package:paper_tracker/widgets/detail_content.dart';
import 'package:paper_tracker/widgets/dropdown.dart';

class WorkflowTemplatePage extends StatefulWidget {
  static const Route = "/workflow/template";

  @override
  _WorkflowTemplatePageState createState() => _WorkflowTemplatePageState();
}

class _WorkflowTemplatePageState extends State<WorkflowTemplatePage> {
  var templateClient = WorkflowTemplateClient();
  var roomClient = RoomClient();

  var stepLabelEditController = TextEditingController();
  var stepDecisionLabelEditController = TextEditingController();
  var roomDropdownController = DropdownController();
  int templateID;
  Future<WorkflowTemplate> futureTemplate;

  @override
  void didChangeDependencies() {
    super.didChangeDependencies();
    templateID = ModalRoute.of(context).settings.arguments;
    futureTemplate = templateClient.getTemplateByID(templateID);
    //futureWorkflow.then((template) => labelEditController.text = template.label);
  }

  @override
  Widget build(BuildContext context) {
    return FutureBuilder(
      future: futureTemplate,
      builder: (context, snapshot) {
        WorkflowTemplate workflow = snapshot.data;

        return DetailContent(
          title: workflow != null ? workflow.label : "",
          iconData: WorkflowTemplate.IconData,
          bottomButtons: [],
          content: workflow != null ? buildContent(workflow) : Container(),
          onRefresh: refreshTemplate,
        );
      },
    );
  }

  Widget buildContent(WorkflowTemplate workflow) {
    return Container(
      padding: EdgeInsets.all(15.0),
      child: WorkflowStepsList(
        steps: workflow.steps,
        roomClient: roomClient,
        onStepAdd: onAddStep,
        primaryScroll: false,
      ),
    );
  }

  void onAddStep(WFStep prev) {
    stepLabelEditController.text = "";
    stepDecisionLabelEditController.text = "";
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

    if (step != null && step.options.length < 2) {
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
      Dropdown(
        getItems: () async {
          var rooms = await roomClient.getAllRooms(refresh: true);
          return rooms.where((room) => room.isLearned).toList();
        },
        controller: roomDropdownController,
        hintName: "learned room",
        icon: Room.IconData,
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
    if (prevStep != null) {
      var createStepRequest = CreateStepRequest(
        decisionLabel: stepDecisionLabelEditController.text,
        previousStepID: prevStep.id,
        step: WFStep(
          label: stepLabelEditController.text,
          roomID: roomDropdownController.selectedItem.id,
        ),
      );
      await templateClient.addStep(templateID, createStepRequest);
    } else {
      var step = WFStep(
        label: stepLabelEditController.text,
        roomID: roomDropdownController.selectedItem.id,
      );
      await templateClient.addStartStep(templateID, step);
    }
    await templateClient.getAllTemplates(refresh: true);

    setState(() {});
    Navigator.of(context).pop();
  }

  Future<void> refreshTemplate() async {
    setState(() {
      futureTemplate = templateClient.getTemplateByID(templateID, refresh: true);
      //futureWorkflow.then((template) => labelEditController.text = template.label);
    });
  }
}
