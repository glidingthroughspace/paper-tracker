import 'package:auto_size_text/auto_size_text.dart';
import 'package:flutter/material.dart';
import 'package:paper_tracker/model/room.dart';

class RoomPage extends StatelessWidget {
  static const String Route = "/page";

  @override
  Widget build(BuildContext context) {
    final Room room = ModalRoute.of(context).settings.arguments;
    return Scaffold(
      backgroundColor: Theme.of(context).backgroundColor,
      body: Column(
        children: [
          buildTopContent(context, room),
        ],
      ),
      bottomNavigationBar: buildBottomNavigation(context, room),
    );
  }

  Widget buildTopContent(BuildContext context, Room room) {
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
                Icons.room,
                color: Theme.of(context).accentColor,
                size: 40.0,
              ),
              SizedBox(width: 15.0),
              Expanded(
                child: AutoSizeText(
                  room.label,
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

  Widget buildBottomNavigation(BuildContext context, Room room) {
    return Container(
      height: 55.0,
      child: BottomAppBar(
        color: Theme.of(context).cardColor,
        child: Container(
          padding: EdgeInsets.symmetric(horizontal: 10.0),
          child: Row(
            mainAxisAlignment: MainAxisAlignment.spaceAround,
            children: [
              IconButton(
                icon: Icon(Icons.save, color: Colors.white),
                onPressed: () {},
              ),
              IconButton(
                icon: Icon(Icons.delete_forever, color: Colors.white),
                onPressed: () {},
              ),
            ],
          ),
        ),
      ),
    );
  }
}
