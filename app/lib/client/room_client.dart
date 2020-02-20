import 'dart:convert';

import 'package:paper_tracker/client/api_client.dart';
import 'package:paper_tracker/model/room.dart';

class RoomClient {
  var apiClient = APIClient();
  static Future<List<Room>> futureRooms;

  Future<List<Room>> getAllRooms({bool refresh = false}) async {
    if (futureRooms != null && !refresh) {
      return futureRooms;
    }

    final response = await apiClient.get("/room");
    if (response.statusCode == 200) {
      final rawList = json.decode(response.body) as List;
      futureRooms = Future.value(rawList.map((i) => Room.fromJson(i)).toList());
      return futureRooms;
    } else {
      throw Exception("Failed to load client");
    }
  }

  Future<Room> getRoomByID(int id, {bool refresh = false}) async {
    if (futureRooms == null || refresh) {
      getAllRooms(refresh: true);
    }

    var rooms = await futureRooms;
    return rooms.firstWhere((room) => room.id == id);
  }

  Future<void> addRoom(Room room) async {
    return apiClient.post("/room", json.encode(room));
  }

  Future<void> updateRoom(Room room) async {
    return apiClient.put("/room/${room.id}", json.encode(room));
  }

  Future<void> deleteRoom(int id) async {
    return apiClient.delete("/room/$id");
  }
}
