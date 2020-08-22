import 'package:cloud_firestore/cloud_firestore.dart';
import 'package:flutter/material.dart';
import 'package:provider/provider.dart';
import 'package:todos/data/repository/task_repo.dart';

import 'UI/app_bar/appbar.dart';
import 'UI/main_screen.dart';

class TodoApp extends StatelessWidget {
  @override
  Widget build(BuildContext context) {
    FocusScope.of(context).addListener(() {});
    return Provider<Repository>(
      create: (_) => Repository(FirebaseFirestore.instance),
      child: MaterialApp(
        theme: ThemeData.dark(),
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
