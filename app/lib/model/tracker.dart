class Tracker {
  int id;
  String label;
  String lastPoll;
  String lastSleepTime;

  Tracker({this.id, this.label, this.lastPoll, this.lastSleepTime});

  factory Tracker.fromJson(Map<String, dynamic> json) {
    return Tracker(
      id: json['ID'],
      label: json['Label'],
      lastPoll: json['LastPoll'],
      lastSleepTime: json['LastSleepTime'],
    );
  }
}