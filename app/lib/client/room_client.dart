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
      await getAllRooms(refresh: true);
    }

    var rooms = await futureRooms;
    try {
      return rooms.firstWhere((room) => room.id == id);
    } catch (_) {
      throw Exception("Failed to get room with id '$id'");
    }
  }

  Future<void> addRoom(Room room) async {
    var response = await apiClient.post("/room", json.encode(room.toJson()));
    if (response.statusCode < 200 || response.statusCode >= 300) {
      throw Exception("Failed to add room");
    }
  }

  Future<void> updateRoom(Room room) async {
    var response = await apiClient.put("/room/${room.id}", json.encode(room.toJson()));
    if (response.statusCode != 200) {
      throw Exception("Failed to update room");
    }
  }

  Future<void> deleteRoom(int id) async {
    var response = await apiClient.delete("/room/$id");
    if (response.statusCode != 200) {
      throw Exception("Failed to delete room");
    }
  }
}
