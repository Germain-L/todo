import 'package:cloud_firestore/cloud_firestore.dart';
import 'package:flutter/material.dart';
import 'package:provider/provider.dart';

import 'UI/main_screen.dart';
import 'UI/new_task_app_bar/new_task_appbar.dart';
import 'data/repository/task_repo.dart';

class TodoApp extends StatelessWidget {
  @override
  Widget build(BuildContext context) {
    FocusScope.of(context).addListener(() {});
    return Provider<Repository>(
      create: (_) => Repository(FirebaseFirestore.instance),
      child: MaterialApp(
        // theme: ThemeData.dark(),
        home: SafeArea(
          child: Scaffold(
            appBar: NewTaskUI(),
            body: MainScreen(),
          ),
        ),
      ),
    );
  }
}
