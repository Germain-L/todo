import 'package:cloud_firestore/cloud_firestore.dart';
import 'package:flutter/material.dart';
import 'package:provider/provider.dart';
import 'package:todos/data/repository/task_repo.dart';

import 'UI/app_bar/appbar.dart';
import 'UI/main_screen.dart';

class TodoApp extends StatelessWidget {
  @override
  Widget build(BuildContext context) {
    return Provider<Repository>(
      create: (_) => Repository(FirebaseFirestore.instance),
      child: MaterialApp(
        theme: ThemeData.dark(),
        home: SafeArea(
          child: Scaffold(
              // appBar: AppBar(
              //   title: TextField(),
              //   centerTitle: false,
              //   elevation: 0,
              //   backgroundColor: Theme.of(context).primaryColor,
              //   actions: [
              //     IconButton(
              //       icon: Icon(Icons.add),
              //       onPressed: () {},
              //     ),
              //   ],
              // ),
              appBar: NewTaskUI(),
              body: MainScreen()),
        ),
      ),
    );
  }
}
