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
  Tutorial("Welcome", "assets/images/tutorial_01.png", "Swipe to begin with the tutorial!"),
  Tutorial("Tabs", "assets/images/tutorial_02.png",
      "The app is divided into 4 different tabs. Each tab provides you with a list of: workflows in execution, workflow templates, rooms and trackers."),
];

class Tutorial {
  final String title;
  final String imageURL;
  final String description;

  Tutorial(this.title, this.imageURL, this.description);
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
            Image(image: AssetImage(tutorial.imageURL)),
            Padding(padding: EdgeInsets.only(top: 30.0)),
            Text(tutorial.description, style: TextStyle(fontSize: 25.0)),
          ],
        ),
      ),
    );
  }
}
