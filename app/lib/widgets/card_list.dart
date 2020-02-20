import 'package:flutter/material.dart';
import 'package:paper_tracker/client/room_client.dart';
import 'package:paper_tracker/model/room.dart';
import 'package:paper_tracker/model/workflow.dart';
import 'package:tuple/tuple.dart';

class CardList<T> extends StatelessWidget {
  final List<Tuple2<String, T>> titleObjectList;
  final void Function(T) onTap;
  final Future<void> Function() onRefresh;
  final IconData iconData;

  CardList({@required this.titleObjectList, @required this.onTap, @required this.iconData, @required this.onRefresh});

  @override
  Widget build(BuildContext context) {
    var icon = Icon(
      iconData,
      color: Colors.white,
      size: 30.0,
    );
    var listChildren = titleObjectList
        .map((tuple) => _buildCard(context, Text(tuple.item1),
            trailing: icon, object: tuple.item2, onTap: onTap, verticalPadding: 10.0))
        .toList();
    return RefreshIndicator(
      onRefresh: onRefresh,
      child: ListView(
        padding: EdgeInsets.only(top: 15.0),
        children: listChildren,
      ),
    );
  }
}

class CheckCardListController {
  Map<String, bool> contentMap = Map();

  void flipState(String title) {
    contentMap[title] = !contentMap[title];
  }

  void fromTitles(List<String> titles) {
    if (titles != null) contentMap = titles.asMap().map((_, title) => MapEntry(title, false));
  }

  void updateFromTitles(List<String> titles) {
    for (var title in titles) {
      if (!contentMap.containsKey(title)) {
        contentMap[title] = false;
      }
    }
  }

  List<String> get checked {
    return contentMap.map((key, value) => value ? MapEntry(key, value) : null).keys.toList();
  }
}

class CheckCardList extends StatefulWidget {
  final List<String> titles;
  final CheckCardListController controller;

  const CheckCardList({Key key, this.titles, @required this.controller}) : super(key: key);

  @override
  _CheckCardListState createState() => _CheckCardListState();
}

class _CheckCardListState extends State<CheckCardList> {
  @override
  void initState() {
    super.initState();
    widget.controller.fromTitles(widget.titles);
  }

  @override
  Widget build(BuildContext context) {
    var listChildren = widget.controller.contentMap
        .map(
          (title, checked) => MapEntry(
              _buildCard(context, Text(title),
                  trailing: Checkbox(value: checked, onChanged: null), object: title, onTap: onTap),
              null),
        )
        .keys
        .toList();
    return ListView(
      padding: EdgeInsets.only(top: 15.0),
      children: listChildren,
      shrinkWrap: true,
    );
  }

  void onTap(String title) {
    setState(() {
      widget.controller.flipState(title);
    });
  }
}

class WorkflowStepsList extends StatefulWidget {
  final List<WFStep> steps;
  final RoomClient roomClient;
  final void Function(WFStep prevStep) onStepAdd;

  const WorkflowStepsList({Key key, @required this.steps, @required this.roomClient, this.onStepAdd}) : super(key: key);

  @override
  _WorkflowStepsListState createState() => _WorkflowStepsListState();
}

class _WorkflowStepsListState extends State<WorkflowStepsList> {
  var selectedDecisionMap = Map<int, String>();

  @override
  Widget build(BuildContext context) {
    var listChildren = getChildrenListFromSteps(widget.steps, 0);

    return ListView(
      padding: EdgeInsets.only(top: 15.0),
      children: listChildren,
      shrinkWrap: true,
    );
  }

  Widget buildContent(WFStep step) {
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
      var texts = step.options.keys.map((label) => Text(label)).toList();
      var isSelected = step.options.keys.map((label) => selectedDecisionMap[step.id] == label).toList();
      children.add(Row(
        children: [
          ToggleButtons(
            children: texts,
            isSelected: isSelected,
            constraints: BoxConstraints.expand(width: 100, height: 40),
            onPressed: (it) {
              setState(() {
                selectedDecisionMap[step.id] = texts.elementAt(it).data;
              });
            },
          ),
        ],
        mainAxisAlignment: MainAxisAlignment.spaceEvenly,
      ));
    }

    return Column(
      children: children,
    );
  }

  List<Widget> getChildrenListFromSteps(List<WFStep> steps, int indentation) {
    var listChildren = List<Widget>();
    for (WFStep step in steps) {
      var nestedSteps = getNestedSteps(step);

      listChildren.add(_buildCard(context, buildContent(step), leftMarginFactor: indentation));

      if (nestedSteps != null) {
        listChildren.addAll(getChildrenListFromSteps(nestedSteps, indentation + 1));
      }
    }

    if (widget.onStepAdd != null) {
      listChildren.add(
        _buildCard(
          context,
          Center(child: Icon(Icons.add)),
          object: steps.last,
          onTap: widget.onStepAdd,
          leftMarginFactor: indentation,
        ),
      );
    }

    return listChildren;
  }

  List<WFStep> getNestedSteps(WFStep step) {
    if (step.options.isEmpty) {
      return null;
    }

    if (!selectedDecisionMap.containsKey(step.id)) {
      selectedDecisionMap[step.id] = step.options.keys.elementAt(0);
    }

    return step.options[selectedDecisionMap[step.id]];
  }
}

Card _buildCard<T>(BuildContext context, Widget content,
    {Widget trailing, T object, void Function(T) onTap, double verticalPadding = 0.0, int leftMarginFactor = 0}) {
  return Card(
    elevation: 8.0,
    margin: EdgeInsets.only(left: 10.0 * (leftMarginFactor + 1), right: 10.0, top: 6.0, bottom: 6.0),
    child: Container(
      decoration: BoxDecoration(color: Theme.of(context).cardColor),
      child: ListTile(
        contentPadding: EdgeInsets.symmetric(horizontal: 20.0, vertical: verticalPadding),
        title: content,
        trailing: trailing,
        onTap: onTap != null ? () => onTap(object) : null,
      ),
    ),
  );
}
