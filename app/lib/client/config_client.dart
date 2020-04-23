import 'dart:convert';

import 'package:paper_tracker/client/api_client.dart';

class ConfigClient {
  var apiClient = APIClient();

  Future<Map<String, dynamic>> getConfig() async {
    final response = await apiClient.get("/config");
    if (response.statusCode == 200) {
      return json.decode(response.body);
    } else {
      throw Exception("Failed to load config");
    }
  }

  Future<void> setConfig(Map<String, dynamic> config) async {
    var response = await apiClient.post("/config", json.encode(config));
    if (response.statusCode < 200 || response.statusCode >= 300) {
      throw Exception("Failed to set config");
    }
  }
}
