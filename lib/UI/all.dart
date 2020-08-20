import 'package:cloud_firestore/cloud_firestore.dart';
import 'package:flutter/material.dart';
import 'package:todos/consts.dart';

class AllTasks extends StatelessWidget {
  @override
  Widget build(BuildContext context) {
    return Container(
      child: StreamBuilder<QuerySnapshot>(
        stream: firestore.collection("tasks").get().asStream(),
        builder: (BuildContext context, AsyncSnapshot<QuerySnapshot> snapshot) {
          print(snapshot.data);
          return Container(
            child: Center(
              child: Text(snapshot.data.docs[0].data()["task"]),
            ),
          );
        },
      ),
    );
  }
}
