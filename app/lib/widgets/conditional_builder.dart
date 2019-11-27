import 'package:flutter/material.dart';

class ConditionalBuilder extends StatelessWidget {
  final bool conditional;
  final Widget falsy;
  final Widget truthy;

  ConditionalBuilder({
    @required this.conditional,
    @required this.truthy,
    @required this.falsy,
  })  : assert(conditional != null),
        assert(truthy != null),
        assert(falsy != null);

  @override
  Widget build(BuildContext context) => conditional ? truthy : falsy;
}