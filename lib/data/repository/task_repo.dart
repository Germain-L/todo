import 'package:cloud_firestore/cloud_firestore.dart';

import '../task.dart';

class Repository {
  final FirebaseFirestore _firestore;

  Repository(this._firestore) : assert(_firestore != null);

  Stream<List<Task>> getTasks() {
    return _firestore.collection("tasks").snapshots().map((snapshot) {
      return snapshot.docs.map((document) {
        return Task(data: document.data()['task'], id: document.id);
      }).toList();
    });
  }

  Future<void> deleteTask(Task task) async {
    await _firestore
        .collection("deleted")
        .doc(task.id)
        .set({"task": task.data});
    await _firestore.collection("tasks").doc(task.id).delete();
  }

  Future<void> createTask(String taskString) async {
    Map firestoreData = Map.from({"task": taskString});

    await _firestore
        .collection("tasks")
        .add(firestoreData)
        .catchError((e) => throw (e));
  }
}
