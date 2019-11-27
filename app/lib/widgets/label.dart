import 'package:flutter/material.dart';

class Label extends StatelessWidget {
  final String text;

  Label(this.text);

  @override
  Widget build(BuildContext context) {
    return Text(
      this.text,
      style: TextStyle(color: Colors.white, fontSize: 20.0, fontWeight: FontWeight.bold),
    );
  }
}