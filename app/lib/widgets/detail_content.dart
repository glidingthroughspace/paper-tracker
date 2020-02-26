import 'package:auto_size_text/auto_size_text.dart';
import 'package:flutter/material.dart';
import 'package:paper_tracker/widgets/conditional_builder.dart';

class DetailContent extends StatelessWidget {
  final IconData iconData;
  final String title;
  final Widget content;
  final List<Widget> bottomButtons;
  final bool disableBackNav;
  final Future<void> Function() onRefresh;

  DetailContent(
      {this.iconData, this.title, this.content, this.bottomButtons, this.disableBackNav = false, this.onRefresh});

  @override
  Widget build(BuildContext context) {
    var scrollConfigChild = CustomScrollView(
      physics: AlwaysScrollableScrollPhysics(),
      slivers: [
        SliverList(
          delegate: SliverChildListDelegate([buildTopContent(context), content]),
        ),
      ],
    );

    return Scaffold(
      body: ScrollConfiguration(
        behavior: NoGlowScrollBehavior(),
        child: onRefresh != null ? RefreshIndicator(onRefresh: onRefresh, child: scrollConfigChild) : scrollConfigChild,
      ),
      bottomNavigationBar: ConditionalBuilder(
        conditional: bottomButtons != null,
        truthy: buildBottomNavigation(context),
        falsy: Container(width: 0, height: 0),
      ),
    );
  }

  Widget buildTopContent(BuildContext context) {
    return Stack(
      children: [
        Container(
          color: Theme.of(context).cardColor,
          width: MediaQuery.of(context).size.width,
          padding: EdgeInsets.only(left: 50.0, right: 20.0, top: 80.0, bottom: 30.0),
          child: Row(
            crossAxisAlignment: CrossAxisAlignment.center,
            children: [
              Icon(iconData, size: 40.0),
              SizedBox(width: 15.0),
              Expanded(
                child: AutoSizeText(
                  title,
                  maxLines: 1,
                  style: TextStyle(
                    fontSize: 45.0,
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
            onTap: disableBackNav ? null : () => Navigator.of(context).pop(),
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
            children: bottomButtons != null ? bottomButtons : [],
          ),
        ),
      ),
    );
  }
}

class NoGlowScrollBehavior extends ScrollBehavior {
  @override
  Widget buildViewportChrome(BuildContext context, Widget child, AxisDirection axisDirection) {
    return child;
  }
}
