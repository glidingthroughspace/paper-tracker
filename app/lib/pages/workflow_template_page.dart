import 'package:flutter/material.dart';
import 'package:flutter/widgets.dart';
import 'package:fluttertoast/fluttertoast.dart';
import 'package:paper_tracker/client/room_client.dart';
import 'package:paper_tracker/client/workflow_template_client.dart';
import 'package:paper_tracker/model/communication/createRevisionRequest.dart';
import 'package:paper_tracker/model/communication/createStepRequest.dart';
import 'package:paper_tracker/model/workflow.dart';
import 'package:paper_tracker/widgets/detail_content.dart';
import 'package:paper_tracker/widgets/dialogs/add_step_dialog.dart';
import 'package:paper_tracker/widgets/dialogs/edit_step_dialog.dart';
import 'package:paper_tracker/widgets/dialogs/single_text_dialog.dart';
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
  var revisionLabelEditingController = TextEditingController();
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
          bottomButtons: buildBottomButtons(template),
          content: template != null ? buildContent(template) : Container(),
          onRefresh: refreshTemplate,
        );
      },
    );
  }

  Widget buildContent(WorkflowTemplate template) {
    var children = <Widget>[
      WorkflowStepsList(
        steps: template.steps,
        roomClient: roomClient,
        onStepAdd: template.stepEditingLocked ? null : onAddStep,
        primaryScroll: false,
        onTap: template.stepEditingLocked ? null : onStepTap,
      )
    ];

    if (template.stepEditingLocked) {
      children.add(Row(
        mainAxisSize: MainAxisSize.max,
        mainAxisAlignment: MainAxisAlignment.center,
        children: [
          Icon(Icons.lock, size: 18.0),
          Padding(padding: EdgeInsets.only(left: 5)),
          Text("Editing is locked"),
        ],
      ));
    }

    return Container(
      padding: EdgeInsets.all(15.0),
      child: Column(
        crossAxisAlignment: CrossAxisAlignment.start,
        children: children,
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

    roomClient.getAllRooms(refresh: true); // In case now a room cannot be deleted
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
      child: OptionsDialog(object: step, options: {
        "Edit Step": onEditStep,
        "Delete Step": onDeleteStep,
        "Move Step Up": onMoveStepUp,
        "Move Step Down": onMoveStepDown,
      }),
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
    refreshTemplate();
    Navigator.of(context).pop();
  }

  void onDeleteStep(WFStep step) async {
    await templateClient.deleteStep(templateID, step.id);
    refreshTemplate();
    Navigator.of(context).pop();
  }

  void onMoveStepUp(WFStep step) async {}

  void onMoveStepDown(WFStep step) async {}

  List<Widget> buildBottomButtons(WorkflowTemplate template) {
    return [
      IconButton(
        icon: Icon(Icons.delete_forever, color: Colors.white),
        onPressed: () => delete(template),
      ),
      IconButton(
        icon: Icon(Icons.content_copy, color: Colors.white),
        onPressed: () => newRevision(template),
      )
    ];
  }

  void delete(WorkflowTemplate template) async {
    if (template.stepEditingLocked) {
      Fluttertoast.showToast(msg: "Can't delete template that is in use");
      return;
    }

    await templateClient.deleteTemplate(template.id);
    await templateClient.getAllTemplates(refresh: true);
    roomClient.getAllRooms(refresh: true);
    Navigator.of(context).pop();
  }

  void newRevision(WorkflowTemplate template) async {
    revisionLabelEditingController.text = template.label;
    showDialog(
      context: context,
      child: SingleTextDialog(
        title: "Copy to New Revision",
        textLabel: "Template Label",
        buttonLabel: "Copy",
        labelController: revisionLabelEditingController,
        onButton: () => onCreateRevision(template),
      ),
    );
  }

  void onCreateRevision(WorkflowTemplate template) async {
    var revisionID = await templateClient.createRevision(
        template.id, CreateRevisionRequest(revisionLabel: revisionLabelEditingController.text));
    await templateClient.getAllTemplates(refresh: true);
    Navigator.of(context).pop();
    Navigator.of(context).pushNamed(WorkflowTemplatePage.Route, arguments: revisionID);
  }
}
