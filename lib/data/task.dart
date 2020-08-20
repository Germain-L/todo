import 'package:cloud_firestore/cloud_firestore.dart';
import 'package:flutter/foundation.dart';

import '../consts.dart';

class Task {
  final String data;
  Timestamp currentTimeStamp;

  Task({@required this.data});


  Future<void> create() async {
    currentTimeStamp = Timestamp.now();

    Map firestoreData = {"task": data, "data": currentTimeStamp};

    await firestore.collection("tasks").add(firestoreData).catchError((e) => print(e));
  }
}