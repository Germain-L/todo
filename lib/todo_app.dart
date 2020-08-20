import 'package:flutter/material.dart';

import 'UI/all.dart';
import 'UI/new_task_bar.dart';

class TodoApp extends StatelessWidget {
  @override
  Widget build(BuildContext context) {
    return MaterialApp(
      theme: ThemeData.dark(),
      home: SafeArea(
        child: Scaffold(
          appBar: NewTaskAppBar(),
          body: AllTasks()
        ),
      ),
    );
  }
}
