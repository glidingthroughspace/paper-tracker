import 'dart:convert';

import 'package:paper_tracker/client/api_client.dart';
import 'package:paper_tracker/model/communication/learningStartResponse.dart';
import 'package:paper_tracker/model/communication/learningStatusResponse.dart';
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

  Future<LearningStartResponse> startLearning(int id) async {
    final response = await apiClient.post("/tracker/$id/learn/start", null);
    if (response.statusCode == 200) {
      return LearningStartResponse.fromJson(json.decode(response.body));
    } else {
      throw Exception("Failed to start learning");
    }
  }

  Future<LearningStatusResponse> getLearningStatus(int id) async {
    final response = await apiClient.get("/tracker/$id/learn/status");
    if (response.statusCode == 200) {
      return LearningStatusResponse.fromJson(json.decode(response.body));
    } else {
      throw Exception("Failed to get learning status");
    }
  }
}