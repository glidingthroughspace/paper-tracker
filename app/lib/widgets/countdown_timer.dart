import 'package:flutter/material.dart';

class CountdownTimer extends StatefulWidget {
  final Duration duration;
  final Color backgroundColor, color;
  final void Function() onComplete;

  const CountdownTimer({Key key, this.duration, this.backgroundColor, this.color, this.onComplete}) : super(key: key);

  @override
  _CountdownTimerState createState() => _CountdownTimerState();
}

class _CountdownTimerState extends State<CountdownTimer> with SingleTickerProviderStateMixin {
  AnimationController controller;
  Animation<double> animation;

  @override
  void initState() {
    super.initState();
    controller = AnimationController (
      vsync: this,
      duration: widget.duration,
    );
    animation = Tween<double>(begin: 1, end: 0).animate(controller);
    animation.addStatusListener((status) {
      if (status == AnimationStatus.completed)
        widget.onComplete();
    });
    controller.forward();
  }

  @override
  Widget build(BuildContext context) {
    return _AnimatedProgressIndicator(
      animation: animation,
      backgroundColor: widget.backgroundColor,
      color: widget.color,
    );
  }
}

class _AnimatedProgressIndicator extends AnimatedWidget {
  final Color backgroundColor, color;

  _AnimatedProgressIndicator({Key key, Animation<double> animation, this.backgroundColor, this.color})
      : super(key: key, listenable: animation);

  @override
  Widget build(BuildContext context) {
    final animation = listenable as Animation<double>;
    return LinearProgressIndicator(
      backgroundColor: backgroundColor,
      valueColor: AlwaysStoppedAnimation<Color>(color),
      value: animation.value,
    );
  }
}
