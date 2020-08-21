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

  Future<void> deleteTask(String docId) async {
    await _firestore.collection("tasks").doc(docId).delete();
  }

  Future<void> createTask(Map data) async {
    Timestamp currentTimeStamp = Timestamp.now();

    Map firestoreData = {"task": data, "data": currentTimeStamp};

    await _firestore
        .collection("tasks")
        .add(firestoreData)
        .catchError((e) => throw (e));
  }
}
