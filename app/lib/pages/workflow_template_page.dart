import 'package:flutter/material.dart';
import 'package:paper_tracker/client/room_client.dart';
import 'package:paper_tracker/client/workflow_template_client.dart';
import 'package:paper_tracker/model/communication/createStepRequest.dart';
import 'package:paper_tracker/model/workflow.dart';
import 'package:paper_tracker/widgets/detail_content.dart';
import 'package:paper_tracker/widgets/dialogs/add_step_dialog.dart';
import 'package:paper_tracker/widgets/dropdown.dart';
import 'package:paper_tracker/widgets/lists/workflow_steps_list.dart';

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
      child: AddStepDialog(
        roomClient: roomClient,
        prevStep: prev,
        labelController: stepLabelEditController,
        decisionController: stepDecisionLabelEditController,
        roomDropdownController: roomDropdownController,
        addStep: addStep,
      ),
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

    refreshTemplate();
    Navigator.of(context).pop();
  }

  Future<void> refreshTemplate() async {
    setState(() {
      futureTemplate = templateClient.getTemplateByID(templateID, refresh: true);
      //futureWorkflow.then((template) => labelEditController.text = template.label);
    });
  }
}
