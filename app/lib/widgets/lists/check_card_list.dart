import 'package:flutter/material.dart';
import 'package:paper_tracker/widgets/lists/card_list.dart';

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
    Map<String, bool> cc = Map.of(contentMap);
    cc.removeWhere((key, value) => !value);
    return cc.keys.toList();
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
        .map((title, checked) => MapEntry(
            ListCard(
                title: Text(title), trailing: Checkbox(value: checked, onChanged: null), object: title, onTap: onTap),
            null))
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
