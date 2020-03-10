import 'package:flutter/material.dart';

class AttributeTable extends StatelessWidget {
  final List<TableRow> children;

  const AttributeTable({Key key, this.children}) : super(key: key);

  @override
  Widget build(BuildContext context) {
    return Container(
      padding: EdgeInsets.all(15.0),
      child: Table(
        defaultVerticalAlignment: TableCellVerticalAlignment.middle,
        columnWidths: {0: FractionColumnWidth(0.3)},
        children: children,
      ),
    );
  }
}
