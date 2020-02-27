import 'package:flutter/material.dart';

class ListCard<T> extends StatelessWidget {
  final Widget title;
  final Widget subtitle;
  final Widget trailing;
  final T object;
  final void Function(T) onTap;
  final double verticalPadding;
  final int indentationFactor;

  const ListCard(
      {Key key,
      this.title,
      this.subtitle,
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
          title: title,
          subtitle: subtitle,
          trailing: trailing,
          onTap: onTap != null ? () => onTap(object) : null,
        ),
      ),
    );
  }
}

class CardListData<T> {
  final String title;
  final String subtitle;
  final T object;

  CardListData(this.title, this.subtitle, this.object);
}

class CardList<T> extends StatelessWidget {
  final List<CardListData<T>> dataList;
  final void Function(T) onTap;
  final Future<void> Function() onRefresh;
  final IconData iconData;

  CardList({@required this.dataList, @required this.onTap, @required this.iconData, @required this.onRefresh});

  @override
  Widget build(BuildContext context) {
    var icon = Icon(
      iconData,
      color: Colors.white,
      size: 30.0,
    );
    var listChildren = dataList
        .map(
          (data) => ListCard(
              title: Text(data.title),
              subtitle: data.subtitle != null ? Text(data.subtitle) : null,
              trailing: icon,
              object: data.object,
              onTap: onTap,
              verticalPadding: 10.0),
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
