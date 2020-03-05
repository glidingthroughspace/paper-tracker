import 'dart:convert';

import 'package:paper_tracker/model/workflow.dart';

import 'api_client.dart';

class WorkflowExecClient {
  var apiClient = APIClient();
  static Future<List<WorkflowExec>> futureExecs;

  Future<List<WorkflowExec>> getAllExecs({bool refresh = false}) async {
    if (futureExecs != null && !refresh) {
      return futureExecs;
    }

    final response = await apiClient.get("/workflow/exec");
    if (response.statusCode == 200) {
      final rawList = json.decode(response.body) as List;
      futureExecs = Future.value(rawList.map((i) => WorkflowExec.fromJson(i)).toList());
      return futureExecs;
    } else {
      throw Exception("Failed to load workflows templates");
    }
  }

  Future<WorkflowExec> getExecByID(int id, {bool refresh = false}) async {
    if (futureExecs == null || refresh) {
      await getAllExecs(refresh: true);
    }

    var execs = await futureExecs;
    try {
      return execs.firstWhere((exec) => exec.id == id);
    } catch (_) {
      throw Exception("Failed to get exec with id '$id'");
    }
  }

  Future<void> startExec(WorkflowExec exec) async {
    var response = await apiClient.post("/workflow/exec", json.encode(exec.toJSON()));
    if (response.statusCode != 200) {
      throw Exception("Failed to start exec");
    }
  }

  Future<void> progressToStep(int execID, int stepID) async {
    var response = await apiClient.post("/workflow/exec/$execID/progress/$stepID", null);
    if (response.statusCode != 200) {
      throw Exception("Failed to progress to step in exec");
    }
  }

  Future<void> cancelExec(int execID) async {
    var response = await apiClient.post("/workflow/exec/$execID/cancel", null);
    if (response.statusCode != 200) {
      throw Exception("Failed to cancel exec");
    }
  }
}
