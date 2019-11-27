import 'package:flutter/material.dart';

class ConditionalBuilder extends StatelessWidget {
  final bool conditional;
  final Function falsyBuilder;
  final Function truthyBuilder;

  ConditionalBuilder({
    @required this.conditional,
    @required this.truthyBuilder,
    @required this.falsyBuilder,
  })  : assert(conditional != null),
        assert(truthyBuilder != null),
        assert(falsyBuilder != null);

  @override
  Widget build(BuildContext context) => conditional ? truthyBuilder() : falsyBuilder();
}