import 'dart:convert';

import 'package:paper_tracker/client/api_client.dart';
import 'package:paper_tracker/model/room.dart';

class RoomClient {
  var apiClient = APIClient();

  Future<List<Room>> fetchRooms() async {
    final response = await apiClient.get("/room");
    if (response.statusCode == 200) {
      final rawList = json.decode(response.body) as List;
      return rawList.map((i) => Room.fromJson(i)).toList();
    } else {
      throw Exception("Failed to load trackers");
    }
  }

  Future<void> addRoom(Room room) async {
    return apiClient.post("/room", json.encode(room));
  }
}