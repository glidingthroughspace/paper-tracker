import 'package:flutter/material.dart';

import 'conditional_builder.dart';

abstract class DropdownCapable {
  int get id;
  String get label;
}

class DropdownController {
  DropdownCapable selectedItem;
  int defaultID;
}

class Dropdown extends StatefulWidget {
  final DropdownController controller;
  final Future<List<DropdownCapable>> Function() getItems;
  final String hintName;
  final IconData icon;
  final bool itemFixed;

  const Dropdown(
      {Key key,
      @required this.getItems,
      @required this.controller,
      this.hintName = "",
      this.icon = Icons.adb,
      this.itemFixed = false})
      : super(key: key);

  @override
  _DropdownState createState() => _DropdownState();
}

class _DropdownState extends State<Dropdown> {
  @override
  Widget build(BuildContext context) {
    return ConditionalBuilder(
      conditional: widget.itemFixed,
      truthy: buildFixedItem(context),
      falsy: buildDropdown(context),
    );
  }

  Widget buildDropdown(BuildContext context) {
    return FutureBuilder(
      future: widget.getItems(),
      builder: (context, snapshot) {
        List<DropdownCapable> itemList = snapshot.hasData ? snapshot.data : [];
        if (snapshot.hasData && widget.controller.selectedItem == null && widget.controller.defaultID != null) {
          widget.controller.selectedItem =
              itemList.firstWhere((room) => room.id == widget.controller.defaultID, orElse: null);
        }

        return DropdownButton(
          icon: Icon(widget.icon),
          items: itemList.map((item) => DropdownMenuItem(value: item, child: Text(item.label))).toList(),
          value: snapshot.hasData ? widget.controller.selectedItem : null,
          isExpanded: true,
          onChanged: (value) {
            setState(() {
              widget.controller.selectedItem = value;
            });
          },
          hint: Text("Please select a ${widget.hintName}"),
          disabledHint: Text("No ${widget.hintName}s found"),
        );
      },
    );
  }

  Widget buildFixedItem(BuildContext context) {
    return Row(
      mainAxisAlignment: MainAxisAlignment.spaceBetween,
      children: [
        Text(
          widget.controller.selectedItem != null ? widget.controller.selectedItem.label : "",
          style: TextStyle(fontSize: 18.0),
        ),
        Icon(widget.icon, size: 25.0),
      ],
    );
  }
}
