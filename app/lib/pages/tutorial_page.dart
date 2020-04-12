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
              onPressed: () => Navigator.of(context).pop(),
              minWidth: double.infinity,
            ),
          ],
        ),
      ),
    );
  }
}

var tutorialList = <TutorialStep>[
  TutorialStep("Welcome", "Swipe to begin with the tutorial!"),
  TutorialStep("Configuration",
      "Enter the URL and optionally the port of the server where the backend software is running on. The default port is 8080."),
  TutorialStep("Main Screen",
      "The app is divided into 4 different tabs. Each tab provides you with a list of: workflows in execution, workflow templates, rooms and trackers. On the top right you find this tutorial."),
  TutorialStep("Tracker List",
      "Switch to the fourth tab for all trackers. If your list is empty use the Flasher tool to setup your tracker. To download the tool visit the URL you entered for configuration."),
  TutorialStep("Pull to refresh",
      "Pull down the empty list to refresh. It should the show your set up tracker. This works on all tabs and also detail pages."),
  TutorialStep("Tracker Detail",
      "Tap on the tracker to open its Detail Page. Tap on the edit button to enable editing and the same button again to save. The delete button deletes the tracker."),
  TutorialStep("Room List",
      "The room tab lists all rooms. Press the add button in the lower right corner to start adding a new one. In the dialog simply enter the name of the room and press 'Create'."),
  TutorialStep("Learn Room 1",
      "On the room detail page tap 'Learn now' to learn the room. On the learn page also select your tracker to learn the room with and start learning."),
  TutorialStep("Learn Room 2",
      "After the tracker got the command, walk the room until the timer runs out. After that select the SSIDs you want to use for tracking and finish with the button on the end of the page."),
  TutorialStep("Workflow Template List",
      "The workflow template list works the same as the room list. Also use the button in the bottom right corner to create a new template."),
  TutorialStep("Edit Workflow Template",
      "Use the button with the plus to add steps to the template. You need to enter a label, if you want a decision label and assign a learned room. Tap on a created step to edit or move it."),
  TutorialStep("Workflow Template Revisions",
      "If you have an execution for your template, the editing is locked. In that case you can create a new revision of this template with the bottom middle button and edit that."),
  TutorialStep("Workflow Exec List",
      "On the workflow exec list, press the bottom right button to start a new worklow. On the following page enter a label and select a tracker, workflow template and decisions."),
  TutorialStep("Workflow Exec",
      "The detail page of an execution shows the current status and the steps. Highlighted in yellow is the current step. Tap on a step to skip or move to this step.")
];

class TutorialStep {
  final String title;
  final String description;

  TutorialStep(this.title, this.description);
}

class TutorialItem extends StatelessWidget {
  final int index;
  TutorialItem(this.index);

  @override
  Widget build(BuildContext context) {
    TutorialStep tutorial = tutorialList[index];

    return Card(
      elevation: 4.0,
      child: Container(
        padding: EdgeInsets.all(20.0),
        child: Column(
          crossAxisAlignment: CrossAxisAlignment.start,
          children: [
            Text(tutorial.title, style: TextStyle(fontSize: 30.0, fontWeight: FontWeight.bold)),
            Padding(padding: EdgeInsets.only(top: 30.0)),
            Container(
              decoration: BoxDecoration(border: Border.all(color: Color.fromRGBO(42, 47, 57, 1.0), width: 4)),
              child: Image(image: AssetImage("assets/images/tutorial_$index.png")),
            ),
            Padding(padding: EdgeInsets.only(top: 30.0)),
            Text(tutorial.description, style: TextStyle(fontSize: 25.0)),
          ],
        ),
      ),
    );
  }
}
