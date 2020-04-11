import 'package:flutter/cupertino.dart';
import 'package:flutter/material.dart';

class TutorialPage extends StatefulWidget {
  static const Route = "/tutorial";

  @override
  _TutorialPageState createState() => _TutorialPageState();
}

class _TutorialPageState extends State<TutorialPage> {
  int currentPage = 0;

  onPageChanged(int index) {
    setState(() {
      currentPage = index;
    });
  }

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      appBar: AppBar(
        title: Text("Paper-Tracker Tutorial"),
        centerTitle: true,
        leading: Container(),
      ),
      body: Container(
        padding: EdgeInsets.all(30.0),
        child: Column(
          children: [
            Expanded(
              child: PageView.builder(
                itemCount: tutorialList.length,
                itemBuilder: (ctx, i) => TutorialItem(i),
                scrollDirection: Axis.horizontal,
                onPageChanged: onPageChanged,
              ),
            ),
            Padding(
              padding: EdgeInsets.only(top: 5.0),
            ),
            Text("${currentPage + 1} / ${tutorialList.length}", style: TextStyle(fontSize: 20.0)),
            Padding(
              padding: EdgeInsets.only(top: 60.0),
            ),
            MaterialButton(
              child: Text("Done"),
              color: Theme.of(context).accentColor,
              onPressed: currentPage == tutorialList.length - 1 ? () => Navigator.of(context).pop() : null,
              minWidth: double.infinity,
            ),
          ],
        ),
      ),
    );
  }
}

var tutorialList = <Tutorial>[
  Tutorial("Welcome", "Swipe to begin with the tutorial!"),
  Tutorial("Configuration",
      "Enter the URL and optionally the port of the server where the backend software is running on. The default port is 8080."),
  Tutorial("Main Screen",
      "The app is divided into 4 different tabs. Each tab provides you with a list of: workflows in execution, workflow templates, rooms and trackers. On the top right you find this tutorial."),
  Tutorial("Tracker List",
      "Switch to the fourth tab for all trackers. If your list is empty use the Flasher tool to setup your tracker."),
  Tutorial("Pull to refresh",
      "Pull down the empty list to refresh. It should the show your set up tracker. This works on all tabs and also detail pages."),
  Tutorial("Tracker Detail",
      "Tap on the tracker to open its Detail Page. Tap on the edit button to enable editing and the same button again to save. The delete button deletes the tracker."),
  Tutorial("Room List", "tutorial_07.png"),
];

class Tutorial {
  final String title;
  final String description;

  Tutorial(this.title, this.description);
}

class TutorialItem extends StatelessWidget {
  final int index;
  TutorialItem(this.index);

  @override
  Widget build(BuildContext context) {
    Tutorial tutorial = tutorialList[index];

    return Card(
      elevation: 4.0,
      child: Container(
        padding: EdgeInsets.all(20.0),
        child: Column(
          crossAxisAlignment: CrossAxisAlignment.start,
          children: [
            Text(tutorial.title, style: TextStyle(fontSize: 30.0, fontWeight: FontWeight.bold)),
            Padding(padding: EdgeInsets.only(top: 30.0)),
            Image(image: AssetImage("assets/images/tutorial_$index.png")),
            Padding(padding: EdgeInsets.only(top: 30.0)),
            Text(tutorial.description, style: TextStyle(fontSize: 25.0)),
          ],
        ),
      ),
    );
  }
}
