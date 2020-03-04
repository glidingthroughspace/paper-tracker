import 'package:flutter/material.dart';

class OptionsDialog<T> extends StatelessWidget {
  final T object;
  final Map<String, void Function(T)> options;

  const OptionsDialog({Key key, this.object, this.options}) : super(key: key);

  @override
  Widget build(BuildContext context) {
    var buttons = options
        .map(
          (title, action) => MapEntry(
              MaterialButton(
                child: Text(title),
                shape: Border.all(color: Colors.grey),
                onPressed: () => action(object),
              ),
              null),
        )
        .keys
        .toList();

    return AlertDialog(
      content: Column(
        mainAxisSize: MainAxisSize.min,
        crossAxisAlignment: CrossAxisAlignment.stretch,
        children: buttons,
      ),
    );
  }
}
