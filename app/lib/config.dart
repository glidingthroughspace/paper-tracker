import 'package:shared_preferences/shared_preferences.dart';

class Config {
  final keyServerURL = "serverURL";

  Future<bool> setServerURL(String url) async {
    final SharedPreferences prefs = await SharedPreferences.getInstance();
    return prefs.setString(keyServerURL, url);
  }

  Future<String> getServerURL() async {
    final SharedPreferences prefs = await SharedPreferences.getInstance();
    return prefs.getString(keyServerURL);
  }
}