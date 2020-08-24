import 'package:cloud_firestore/cloud_firestore.dart';

import '../task.dart';

class Repository {
  final FirebaseFirestore _firestore;

  Repository(this._firestore) : assert(_firestore != null);

  Stream<List<Task>> getTasks() {
    return _firestore.collection("tasks").snapshots().map((snapshot) {
      return snapshot.docs.map((document) {
        return Task(
          data: document.data()['task'],
          timeCreated: document.data()['time'],
          done: document.data()['done'],
          id: document.id,
        );
      }).toList();
    });
  }

  Future<void> deleteTask(Task task) async {
    await _firestore.collection("deleted").doc(task.id).set({
      "task": task.data,
      "done": true,
      "time": task.timeCreated,
    });
    await _firestore.collection("tasks").doc(task.id).delete();
  }

  Future<bool> createTask(String taskString) async {
    Timestamp timeCreated = Timestamp.now();
    await _firestore
        .collection("tasks")
        .add({"task": taskString, "done": false, "time": timeCreated})
        .catchError((e) => throw (e));
    return true;
  }

  Future<void> markTaskDone(Task task) async {
    await _firestore.collection("tasks").doc(task.id).set({
      "task": task.data,
      "done": true,
      "time": task.timeCreated,
    });
  }

  Future<void> unmarkTaskDone(Task task) async {
    await _firestore.collection("tasks").doc(task.id).set({
      "task": task.data,
      "done": false,
      "time": task.timeCreated,
    });
  }
}
