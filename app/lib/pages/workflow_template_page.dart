import 'package:flutter/material.dart';
import 'package:paper_tracker/client/room_client.dart';
import 'package:paper_tracker/client/workflow_template_client.dart';
import 'package:paper_tracker/model/communication/createStepRequest.dart';
import 'package:paper_tracker/model/workflow.dart';
import 'package:paper_tracker/widgets/detail_content.dart';
import 'package:paper_tracker/widgets/dialogs/add_step_dialog.dart';
import 'package:paper_tracker/widgets/dialogs/edit_step_dialog.dart';
import 'package:paper_tracker/widgets/dialogs/workflow_step_dialog.dart';
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
  var stepDecisionLabelEditController = [TextEditingController()];
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
        WorkflowTemplate template = snapshot.data;

        return DetailContent(
          title: template != null ? template.label : "",
          iconData: WorkflowTemplate.IconData,
          bottomButtons: [],
          content: template != null ? buildContent(template) : Container(),
          onRefresh: refreshTemplate,
        );
      },
    );
  }

  Widget buildContent(WorkflowTemplate template) {
    return Container(
      padding: EdgeInsets.all(15.0),
      child: WorkflowStepsList(
        steps: template.steps,
        roomClient: roomClient,
        onStepAdd: onAddStep,
        primaryScroll: false,
        onTap: onStepTap,
      ),
    );
  }

  void onAddStep(WFStep prev) {
    if (stepDecisionLabelEditController.length == 0) {
      stepDecisionLabelEditController = [TextEditingController()];
    }

    stepLabelEditController.text = "";
    stepDecisionLabelEditController[0].text = "";
    showDialog(
      context: context,
      child: AddStepDialog(
        roomClient: roomClient,
        prevStep: prev,
        labelController: stepLabelEditController,
        decisionController: stepDecisionLabelEditController[0],
        roomDropdownController: roomDropdownController,
        addStep: addStep,
      ),
    );
  }

  void addStep(WFStep prevStep) async {
    if (prevStep != null) {
      var createStepRequest = CreateStepRequest(
        decisionLabel: stepDecisionLabelEditController[0].text,
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

  void onStepTap(WFStep step) {
    showDialog(
      context: context,
      child: WorkflowStepDialog(
        step: step,
        onEdit: onEditStep,
        onDelete: onDeleteStep,
      ),
    );
  }

  void onEditStep(WFStep step) {
    Navigator.of(context).pop();

    stepLabelEditController.text = step.label;
    roomDropdownController.defaultID = step.roomID;
    stepDecisionLabelEditController = [];
    step.options.forEach((decision, step) {
      stepDecisionLabelEditController.add(TextEditingController());
      stepDecisionLabelEditController.last.text = decision;
    });

    showDialog(
      context: context,
      child: EditStepDialog(
        step: step,
        editStep: editStep,
        labelController: stepLabelEditController,
        roomDropdownController: roomDropdownController,
        decisionController: stepDecisionLabelEditController,
        roomClient: roomClient,
      ),
    );
  }

  void editStep(WFStep step) async {
    var options = Map.fromIterables(
      stepDecisionLabelEditController.map((controller) => controller.text),
      step.options.map((decision, steps) => MapEntry([WFStep(id: steps.first.id)], null)).keys,
    );

    var editedStep = WFStep(
      id: step.id,
      label: stepLabelEditController.text,
      roomID: roomDropdownController.selectedItem.id,
      options: options,
    );

    await templateClient.updateStep(templateID, editedStep);
    Navigator.of(context).pop();
    refreshTemplate();
  }

  void onDeleteStep(WFStep step) {}
}
