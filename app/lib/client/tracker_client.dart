import 'dart:convert';

import 'package:paper_tracker/client/api_client.dart';
import 'package:paper_tracker/model/tracker.dart';

class TrackerClient {
  var apiClient = APIClient();

  Future<List<Tracker>> fetchTrackers() async {
    final response = await apiClient.get("/tracker");
    if (response.statusCode == 200) {
      final rawList = json.decode(response.body) as List;
      return rawList.map((i) => Tracker.fromJson(i)).toList();
    } else {
      throw Exception("Failed to load trackers");
    }
  }
}