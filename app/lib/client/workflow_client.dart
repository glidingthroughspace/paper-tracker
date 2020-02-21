import 'dart:convert';

import 'package:paper_tracker/model/communication/createStepRequest.dart';
import 'package:paper_tracker/model/workflow.dart';

import 'api_client.dart';

class WorkflowClient {
  var apiClient = APIClient();
  static Future<List<Workflow>> futureWorkflows;

  Future<List<Workflow>> getAllWorkflows({bool refresh = false}) async {
    if (futureWorkflows != null && !refresh) {
      return futureWorkflows;
    }

    final response = await apiClient.get("/workflow");
    if (response.statusCode == 200) {
      final rawList = json.decode(response.body) as List;
      futureWorkflows = Future.value(rawList.map((i) => Workflow.fromJson(i)).toList());
      return futureWorkflows;
    } else {
      throw Exception("Failed to load workflows");
    }
  }

  Future<Workflow> getWorkflowByID(int id, {bool refresh = false}) async {
    if (futureWorkflows == null || refresh) {
      getAllWorkflows(refresh: true);
    }

    var workflows = await futureWorkflows;
    return workflows.firstWhere((workflow) => workflow.id == id);
  }

  Future<void> addStartStep(int workflowID, WFStep step) async {
    return apiClient.post("/workflow/$workflowID/start", json.encode(step.toJSON()));
  }

  Future<void> addStep(int workflowID, CreateStepRequest stepRequest) async {
    return apiClient.post("/workflow/$workflowID/step", json.encode(stepRequest.toJson()));
  }
}
