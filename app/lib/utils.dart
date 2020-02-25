import 'package:flutter/material.dart';

TableRow getTableSpacing(double padding) {
  return TableRow(children: [Padding(padding: EdgeInsets.only(top: padding)), Container()]);
}
