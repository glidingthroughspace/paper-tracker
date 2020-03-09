import 'package:http/http.dart' as http;
import 'package:paper_tracker/config.dart';

class APIClient {
  var config = Config();

  Future<bool> isAvailable() async {
    if (await config.getServerURL() == null) {
      return false;
    }

    try {
      await get("/tracker");
    } catch (err) {
      return false;
    }
    return true;
  }

  Future<Uri> _buildURI(String path, Map<String, String> query) async {
    return Uri.http(await config.getServerURL(), path, query);
  }

  Future<http.Response> get(String path, [Map<String, String> query]) async {
    return http.get(await _buildURI(path, query));
  }

  Future<http.Response> post(String path, String body, [Map<String, String> query]) async {
    return http.post(await _buildURI(path, query), body: body);
  }

  Future<http.Response> put(String path, String body, [Map<String, String> query]) async {
    return http.put(await _buildURI(path, query), body: body);
  }

  Future<http.Response> delete(String path, [Map<String, String> query]) async {
    return http.delete(await _buildURI(path, query));
  }
}
