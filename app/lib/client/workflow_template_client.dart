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
      throw Exception("Failed to load workflows templates");
    }
  }

  Future<WorkflowTemplate> getTemplateByID(int id, {bool refresh = false}) async {
    if (futureTemplates == null || refresh) {
      await getAllTemplates(refresh: true);
    }

    var templates = await futureTemplates;
    try {
      return templates.firstWhere((template) => template.id == id);
    } catch (_) {
      throw Exception("Failed to get template with id '$id'");
    }
  }

  Future<void> addStartStep(int templateID, WFStep step) async {
    var response = await apiClient.post("/workflow/template/$templateID/start", json.encode(step.toJSON()));
    if (response.statusCode != 200) {
      throw Exception("Failed to add start step to template");
    }
  }

  Future<void> addStep(int templateID, CreateStepRequest stepRequest) async {
    var response = await apiClient.post("/workflow/template/$templateID/step", json.encode(stepRequest.toJson()));
    if (response.statusCode != 200) {
      throw Exception("Failed to add step to template");
    }
  }

  Future<void> createTemplate(WorkflowTemplate template) async {
    var response = await apiClient.post("/workflow/template", json.encode(template.toJson()));
    if (response.statusCode != 200) {
      throw Exception("Failed to create template");
    }
  }

  Future<WFStep> getStepByID(int templateID, int stepID) async {
    final response = await apiClient.get("/workflow/template/$templateID/step/$stepID");
    if (response.statusCode == 200) {
      return WFStep.fromJson(json.decode(response.body));
    } else {
      throw Exception("Failed to get step with id '$stepID'");
    }
  }

  Future<void> updateStep(int templateID, WFStep step) async {
    var response = await apiClient.put("/workflow/template/$templateID/step/${step.id}", json.encode(step.toJSON()));
    if (response.statusCode != 200) {
      throw Exception("Failed to update step in template");
    }
  }

  Future<void> deleteStep(int templateID, int stepID) async {
    var response = await apiClient.delete("/workflow/template/$templateID/step/$stepID");
    if (response.statusCode != 200) {
      throw Exception("Failed to delete step in template");
    }
  }
}
