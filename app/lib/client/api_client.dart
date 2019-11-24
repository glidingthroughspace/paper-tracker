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

  Future<http.Response> get(String path) async {
    var url = Uri.http(await config.getServerURL(), path);
    return http.get(url);
  }
}