import 'package:flutter/material.dart';
import 'package:tuple/tuple.dart';

class CardList<T> extends StatelessWidget {
  final List<Tuple2<String, T>> titleObjectList;
  final void Function(T) onTap;
  final IconData iconData;

  CardList({this.titleObjectList, this.onTap, this.iconData});

  @override
  Widget build(BuildContext context) {
    var icon = Icon(
      iconData,
      color: Colors.white,
      size: 30.0,
    );
    var listChildren =
        titleObjectList.map((tuple) => _buildCard(context, tuple.item1, icon, tuple.item2, onTap, 10.0)).toList();
    return ListView(
      padding: EdgeInsets.only(top: 15.0),
      children: listChildren,
      shrinkWrap: true,
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
        .map((title, checked) =>
            MapEntry(_buildCard(context, title, Checkbox(value: checked, onChanged: null), title, onTap, 0.0), null))
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

Card _buildCard<T>(
    BuildContext context, String title, Widget trailing, T object, void Function(T) onTap, double verticalPadding) {
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
