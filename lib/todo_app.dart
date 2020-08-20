import 'package:flutter/material.dart';

import 'UI/all.dart';

class TodoApp extends StatelessWidget {
  @override
  Widget build(BuildContext context) {
    return MaterialApp(
      theme: ThemeData.dark(),
      home: Scaffold(
        body: AllTasks()
      ),
    );
  }
}
