import 'package:flutter/material.dart';

class WaitingTextDialog extends StatelessWidget {
  final String text;

  const WaitingTextDialog({Key key, this.text = "Waiting..."}) : super(key: key);

  @override
  Widget build(BuildContext context) {
    return WillPopScope(
      onWillPop: () async => false,
      child: SimpleDialog(children: [
        Center(
          child: Column(children: [
            CircularProgressIndicator(),
            SizedBox(
              height: 10,
            ),
            Text(text),
          ]),
        )
      ]),
    );
  }
}
