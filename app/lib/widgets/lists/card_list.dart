import 'package:flutter/material.dart';
import 'package:tuple/tuple.dart';

class ListCard<T> extends StatelessWidget {
  final Widget content;
  final Widget trailing;
  final T object;
  final void Function(T) onTap;
  final double verticalPadding;
  final int indentationFactor;

  const ListCard(
      {Key key,
      this.content,
      this.trailing,
      this.object,
      this.onTap,
      this.verticalPadding = 0.0,
      this.indentationFactor = 0})
      : super(key: key);

  @override
  Widget build(BuildContext context) {
    return Card(
      elevation: 8.0,
      margin: EdgeInsets.only(left: 10.0 * (indentationFactor + 1), right: 10.0, top: 6.0, bottom: 6.0),
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
}

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
        .map(
          (tuple) => ListCard(
              content: Text(tuple.item1), trailing: icon, object: tuple.item2, onTap: onTap, verticalPadding: 10.0),
        )
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
