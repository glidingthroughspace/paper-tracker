import 'package:flutter/material.dart';
import 'package:paper_tracker/client/tracker_client.dart';
import 'package:paper_tracker/widgets/conditional_builder.dart';
import 'package:paper_tracker/widgets/countdown_timer.dart';

class WaitTrackerPollDialog extends StatelessWidget {
  final int trackerID;
  final void Function() onWaitFinished;
  final TrackerClient trackerClient;

  const WaitTrackerPollDialog({Key key, this.trackerID, this.trackerClient, this.onWaitFinished}) : super(key: key);

  @override
  Widget build(BuildContext context) {
    return AlertDialog(
      content: Column(
        mainAxisSize: MainAxisSize.min,
        children: [
          Text(
            "Wait for tracker to answer...",
            style: TextStyle(
              fontSize: 20.0,
            ),
          ),
          Padding(padding: EdgeInsets.only(top: 10.0)),
          buildTimer(),
        ],
      ),
    );
  }

  FutureBuilder<int> buildTimer() {
    return FutureBuilder(
      future: trackerClient.getNextPollSecs(trackerID),
      builder: (context, snapshot) {
        var timeToWait = snapshot.hasData ? snapshot.data : 0;

        return ConditionalBuilder(
          conditional: snapshot.hasData,
          truthy: CountdownTimer(
            duration: Duration(seconds: timeToWait),
            backgroundColor: Theme.of(context).cardColor,
            color: Theme.of(context).accentColor,
            onComplete: () {
              Navigator.of(context).pop();
              onWaitFinished();
            },
          ),
          falsy: Container(),
        );
      },
    );
  }
}
