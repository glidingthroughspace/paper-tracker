import 'package:flutter/material.dart';
import 'package:paper_tracker/client/room_client.dart';
import 'package:paper_tracker/model/room.dart';
import 'package:paper_tracker/model/workflow.dart';
import 'package:paper_tracker/widgets/lists/card_list.dart';

class WorkflowStepsListController {
  var selectedDecisionMap = Map<int, String>();
}

class WorkflowStepsList extends StatefulWidget {
  final List<WFStep> steps;
  final RoomClient roomClient;
  final void Function(WFStep prevStep) onStepAdd;
  final WorkflowStepsListController controller;
  final Map<int, ExecStepInfo> stepInfos;
  final int currentStep;
  final bool primaryScroll;

  const WorkflowStepsList(
      {Key key,
      @required this.steps,
      @required this.roomClient,
      this.onStepAdd,
      this.controller,
      this.stepInfos,
      this.currentStep,
      this.primaryScroll = true})
      : super(key: key);

  @override
  _WorkflowStepsListState createState() => _WorkflowStepsListState();
}

class _WorkflowStepsListState extends State<WorkflowStepsList> {
  Map<int, String> selectedDecisionMap;

  @override
  void initState() {
    super.initState();
    if (widget.controller != null)
      selectedDecisionMap = widget.controller.selectedDecisionMap;
    else
      selectedDecisionMap = Map<int, String>();
  }

  @override
  Widget build(BuildContext context) {
    var listChildren = getChildrenListFromSteps(widget.steps, 0);

    return ListView(
      padding: EdgeInsets.only(top: 15.0),
      children: listChildren,
      shrinkWrap: true,
      primary: widget.primaryScroll,
    );
  }

  List<Widget> getChildrenListFromSteps(List<WFStep> steps, int indentation) {
    var listChildren = List<Widget>();
    for (WFStep step in steps) {
      var stepInfo = widget.stepInfos != null ? widget.stepInfos[step.id] : null;
      var nestedSteps = getNestedSteps(step, stepInfo);

      listChildren.add(ListCard(
        title: buildCardContent(step, stepInfo),
        indentationFactor: indentation,
        verticalPadding: 5.0,
        color: widget.currentStep == step.id ? WFStep.CurrentStepColor : null,
      ));

      if (nestedSteps != null) {
        listChildren.addAll(getChildrenListFromSteps(nestedSteps, indentation + 1));
      }
    }

    if (widget.onStepAdd != null) {
      listChildren.add(
        ListCard(
          title: Center(child: Icon(Icons.add)),
          object: steps.length > 0 ? steps.last : null,
          onTap: widget.onStepAdd,
          indentationFactor: indentation,
        ),
      );
    }

    return listChildren;
  }

  List<WFStep> getNestedSteps(WFStep step, ExecStepInfo stepInfo) {
    if (step.options.isEmpty) {
      return null;
    }

    if (stepInfo != null && stepInfo.decision != null) {
      selectedDecisionMap[step.id] = widget.stepInfos[step.id].decision;
    } else if (!selectedDecisionMap.containsKey(step.id)) {
      selectedDecisionMap[step.id] = step.options.keys.elementAt(0);
    }

    return step.options[selectedDecisionMap[step.id]];
  }

  Widget buildCardContent(WFStep step, ExecStepInfo stepInfo) {
    List<Row> children = [
      Row(
        children: [
          Text(step.label),
        ],
      ),
      Row(
        children: [
          FutureBuilder(
            future: widget.roomClient.getRoomByID(step.roomID),
            builder: (context, snapshot) {
              if (snapshot.hasData) {
                Room room = snapshot.data;
                return Text("Room: ${room.label}");
              }
              return Text("Room: ");
            },
          ),
        ],
      )
    ];

    if (step.options.isNotEmpty) {
      children.addAll(buildOptionsButtons(step, stepInfo));
    }

    return Column(
      children: children,
    );
  }

  List<Row> buildOptionsButtons(WFStep step, ExecStepInfo stepInfo) {
    var children = <Row>[];

    List<Widget> buttonContents;
    List<bool> isSelected;
    if (stepInfo != null && stepInfo.decision != null) {
      buttonContents = [Text(widget.stepInfos[step.id].decision)];
      isSelected = [true];
    } else {
      buttonContents = List.of(step.options.keys.map((label) => Text(label)));
      isSelected = step.options.keys.map((label) => selectedDecisionMap[step.id] == label).toList();
    }

    if (step.options.length < 2 && widget.onStepAdd != null) {
      buttonContents.add(Icon(Icons.add));
      isSelected.add(false);
    }

    children.add(Row(children: [Padding(padding: EdgeInsets.only(top: 10.0))]));
    children.add(Row(
      children: [
        ToggleButtons(
          children: buttonContents,
          isSelected: isSelected,
          constraints: BoxConstraints.expand(width: 100, height: 40),
          color: Colors.white,
          selectedColor: Colors.deepOrange,
          fillColor: Colors.blueGrey,
          borderColor: Colors.grey,
          selectedBorderColor: Colors.grey,
          onPressed: (it) {
            setState(() {
              if (buttonContents.elementAt(it) is Text) {
                Text text = buttonContents.elementAt(it);
                selectedDecisionMap[step.id] = text.data;
              } else {
                widget.onStepAdd(step);
              }
            });
          },
        ),
      ],
      mainAxisAlignment: MainAxisAlignment.spaceEvenly,
    ));

    return children;
  }
}
