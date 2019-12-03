import 'package:flutter/material.dart';

class CardList<T> extends StatelessWidget {
  final Map<String, T> titleObjectMap;
  final void Function(T) onTap;
  final IconData iconData;

  CardList({this.titleObjectMap, this.onTap, this.iconData});

  @override
  Widget build(BuildContext context) {
    var icon = Icon(
      iconData,
      color: Colors.white,
      size: 30.0,
    );
    var listChildren = titleObjectMap.map((label, object) => MapEntry(_buildCard(context, label, icon, object, onTap, 10.0), null)).keys.toList();
    return ListView(
      padding: EdgeInsets.only(top: 15.0),
      children: listChildren,
      shrinkWrap: true,
    );
  }
}

class CheckCardListController {
  Map<String, bool> contentMap;

  void flipState(String title) {
    contentMap[title] = !contentMap[title];
  }

  void fromTitles(List<String> titles) {
    contentMap = titles.asMap().map((_, title) => MapEntry(title, false));
  }
}

class CheckCardList extends StatefulWidget {
  final List<String> titles;
  final CheckCardListController controller;

  const CheckCardList({Key key, @required this.titles, @required this.controller}) : super(key: key);

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
    var listChildren = widget.controller.contentMap.map((title, checked) => MapEntry(_buildCard(context, title, Checkbox(value: checked, onChanged: null), title, onTap, 0.0), null)).keys.toList();
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

Card _buildCard<T>(BuildContext context, String title, Widget trailing, T object, void Function(T) onTap, double verticalPadding) {
  return Card(
    elevation: 8.0,
    margin: EdgeInsets.symmetric(horizontal: 10.0, vertical: 6.0),
    child: Container(
      decoration: BoxDecoration(color: Theme.of(context).cardColor),
      child: ListTile(
        contentPadding: EdgeInsets.symmetric(horizontal: 20.0, vertical: verticalPadding),
        title: Text(
          title,
        ),
        trailing: trailing,
        onTap: () => onTap(object),
      ),
    ),
  );
}
