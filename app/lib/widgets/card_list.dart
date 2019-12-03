import 'package:flutter/material.dart';

class CardList<T> extends StatelessWidget {
  final Map<String, T> titleObjectMap;
  final void Function(T) onTap;

  CardList({this.titleObjectMap, this.onTap});

  @override
  Widget build(BuildContext context) {
    var listChildren = titleObjectMap.map((label, object) => MapEntry(buildCard(context, label, object), null)).keys.toList();
    return ListView(
      padding: EdgeInsets.only(top: 15.0),
      children: listChildren,
      shrinkWrap: true,
    );
  }

  Card buildCard(BuildContext context, String title, T object) {
    return Card(
      elevation: 8.0,
      margin: EdgeInsets.symmetric(horizontal: 10.0, vertical: 6.0),
      child: Container(
        decoration: BoxDecoration(color: Theme.of(context).cardColor),
        child: ListTile(
          contentPadding: EdgeInsets.symmetric(horizontal: 20.0, vertical: 10.0),
          title: Text(
            title,
          ),
          trailing: Icon(
            Icons.keyboard_arrow_right,
            color: Colors.white,
            size: 30.0,
          ),
          onTap: () => onTap(object),
        ),
      ),
    );
  }
}
