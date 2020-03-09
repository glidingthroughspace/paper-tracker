import 'dart:convert';

import 'package:paper_tracker/client/api_client.dart';
import 'package:paper_tracker/model/communication/learningFinishRequest.dart';
import 'package:paper_tracker/model/communication/learningStartResponse.dart';
import 'package:paper_tracker/model/communication/learningStatusResponse.dart';
import 'package:paper_tracker/model/tracker.dart';

class TrackerClient {
  var apiClient = APIClient();
  static Future<List<Tracker>> futureTrackers;

  Future<List<Tracker>> getAllTrackers({bool refresh = false}) async {
    if (futureTrackers != null && !refresh) {
      return futureTrackers;
    }

    final response = await apiClient.get("/tracker");
    if (response.statusCode == 200) {
      final rawList = json.decode(response.body) as List;
      futureTrackers = Future.value(rawList.map((i) => Tracker.fromJson(i)).toList());
      return futureTrackers;
    } else {
      throw Exception("Failed to load trackers");
    }
  }

  Future<Tracker> getTrackerByID(int id, {bool refresh = false}) async {
    if (futureTrackers == null || refresh) {
      await getAllTrackers(refresh: true);
    }

    var rooms = await futureTrackers;
    try {
      return rooms.firstWhere((tracker) => tracker.id == id);
    } catch (_) {
      throw Exception("Failed to get tracker with id '$id'");
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

  Future<void> finishLearning(int trackerID, int roomID, List<String> ssids) async {
    final response = await apiClient.post(
        "/tracker/$trackerID/learn/finish", json.encode(LearningFinishRequest(roomID: roomID, ssids: ssids).toJson()));
    if (response.statusCode < 200 || response.statusCode >= 300) {
      throw Exception("Failed to finish learning");
    }
  }

  Future<void> cancelLearning(int trackerID) async {
    final response = await apiClient.post("/tracker/$trackerID/learn/cancel", null);
    if (response.statusCode < 200 || response.statusCode >= 300) {
      throw Exception("Failed to cancel learning");
    }
  }

  Future<void> updateTracker(Tracker tracker) async {
    final response = await apiClient.put("/tracker/${tracker.id}", json.encode(tracker.toJson()));
    if (response.statusCode != 200) {
      throw Exception("Failed to update tracker");
    }
  }

  Future<void> deleteTracker(int trackerID) async {
    final response = await apiClient.delete("/tracker/$trackerID");
    if (response.statusCode != 200) {
      throw Exception("Failed to delete tracker");
    }
  }
}
