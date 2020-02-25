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

  Future<Uri> _buildURI(String path) async {
    return Uri.http(await config.getServerURL(), path);
  }

  Future<http.Response> get(String path) async {
    return http.get(await _buildURI(path));
  }

  Future<http.Response> post(String path, String body) async {
    return http.post(await _buildURI(path), body: body);
  }

  Future<http.Response> put(String path, String body) async {
    return http.put(await _buildURI(path), body: body);
  }

  Future<http.Response> delete(String path) async {
    return http.delete(await _buildURI(path));
  }
}
