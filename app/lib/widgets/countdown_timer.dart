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
  void dispose() {
    controller.dispose();
    super.dispose();
  }

  @override
  void initState() {
    super.initState();
    controller = AnimationController(
      vsync: this,
      duration: widget.duration,
    );
    animation = Tween<double>(begin: 1, end: 0).animate(controller);
    animation.addStatusListener((status) {
      if (status == AnimationStatus.completed) widget.onComplete();
    });
    controller.forward();
  }

  @override
  Widget build(BuildContext context) {
    return _AnimatedCountdownTimer(
      duration: widget.duration,
      animation: animation,
      backgroundColor: widget.backgroundColor,
      color: widget.color,
    );
  }
}

class _AnimatedCountdownTimer extends AnimatedWidget {
  final Color backgroundColor, color;
  final Duration duration;

  _AnimatedCountdownTimer({Key key, Animation<double> animation, this.backgroundColor, this.color, this.duration})
      : super(key: key, listenable: animation);

  String get countdownString {
    final animation = listenable as Animation<double>;
    Duration remaining = duration * animation.value;
    return "${remaining.inSeconds} seconds left";
  }

  @override
  Widget build(BuildContext context) {
    final animation = listenable as Animation<double>;
    return Column(
      children: [
        SizedBox(
          height: 30.0,
          child: LinearProgressIndicator(
            backgroundColor: backgroundColor,
            valueColor: AlwaysStoppedAnimation<Color>(color),
            value: animation.value,
          ),
        ),
        SizedBox(height: 5.0),
        Text("$countdownString"),
      ],
    );
  }
}
