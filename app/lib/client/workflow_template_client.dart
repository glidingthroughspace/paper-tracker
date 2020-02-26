import 'dart:convert';

import 'package:paper_tracker/model/communication/createStepRequest.dart';
import 'package:paper_tracker/model/workflow.dart';

import 'api_client.dart';

class WorkflowTemplateClient {
  var apiClient = APIClient();
  static Future<List<WorkflowTemplate>> futureTemplates;

  Future<List<WorkflowTemplate>> getAllTemplates({bool refresh = false}) async {
    if (futureTemplates != null && !refresh) {
      return futureTemplates;
    }

    final response = await apiClient.get("/workflow/template");
    if (response.statusCode == 200) {
      final rawList = json.decode(response.body) as List;
      futureTemplates = Future.value(rawList.map((i) => WorkflowTemplate.fromJson(i)).toList());
      return futureTemplates;
    } else {
      throw Exception("Failed to load workflows execs");
    }
  }

  Future<WorkflowTemplate> getTemplateByID(int id, {bool refresh = false}) async {
    if (futureTemplates == null || refresh) {
      await getAllTemplates(refresh: true);
    }

    var templates = await futureTemplates;
    return templates.firstWhere((template) => template.id == id);
  }

  Future<void> addStartStep(int templateID, WFStep step) async {
    return apiClient.post("/workflow/template/$templateID/start", json.encode(step.toJSON()));
  }

  Future<void> addStep(int templateID, CreateStepRequest stepRequest) async {
    return apiClient.post("/workflow/template/$templateID/step", json.encode(stepRequest.toJson()));
  }

  Future<void> createTemplate(WorkflowTemplate template) async {
    return apiClient.post("/workflow/template", json.encode(template.toJson()));
  }
}
