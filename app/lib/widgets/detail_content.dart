import 'package:auto_size_text/auto_size_text.dart';
import 'package:flutter/material.dart';

class DetailContent extends StatelessWidget {
  final IconData iconData;
  final String title;
  final Widget content;
  final List<Widget> bottomButtons;

  DetailContent({this.iconData, this.title, this.content, this.bottomButtons});

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      backgroundColor: Theme.of(context).backgroundColor,
      body: Column(
        children: [
          buildTopContent(context),
          content
        ],
      ),
      bottomNavigationBar: buildBottomNavigation(context),
    );
  }

  Widget buildTopContent(BuildContext context) {
    return Stack(
      children: [
        Container(
          color: Theme.of(context).cardColor,
          width: MediaQuery.of(context).size.width,
          padding: EdgeInsets.only(left: 50.0, right: 10.0, top: 80.0, bottom: 30.0),
          child: Row(
            crossAxisAlignment: CrossAxisAlignment.center,
            children: [
              Icon(
                iconData,
                color: Theme.of(context).accentColor,
                size: 40.0,
              ),
              SizedBox(width: 15.0),
              Expanded(
                child: AutoSizeText(
                  title,
                  maxLines: 1,
                  style: TextStyle(
                    fontSize: 45.0,
                    color: Colors.white,
                  ),
                ),
              ),
            ],
          ),
        ),
        Positioned(
          left: 15.0,
          top: 60.0,
          child: InkWell(
            onTap: () => Navigator.of(context).pop(),
            child: Icon(Icons.arrow_back, color: Colors.white),
          ),
        ),
      ],
    );
  }

  Widget buildBottomNavigation(BuildContext context) {
    return Container(
      height: 55.0,
      child: BottomAppBar(
        color: Theme.of(context).cardColor,
        child: Container(
          padding: EdgeInsets.symmetric(horizontal: 10.0),
          child: Row(
            mainAxisAlignment: MainAxisAlignment.spaceAround,
            children: bottomButtons,
          ),
        ),
      ),
    );
  }
}
