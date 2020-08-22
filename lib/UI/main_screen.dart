import 'package:flutter/material.dart';
import 'package:provider/provider.dart';

import '../data/repository/task_repo.dart';
import '../data/task.dart';
import 'main_screen/all_task_screen.dart';
import 'main_screen/error_screen.dart';
import 'main_screen/loading_screen.dart';

class MainScreen extends StatefulWidget {
  @override
  _MainScreenState createState() => _MainScreenState();
}

class _MainScreenState extends State<MainScreen> {
  Stream<List<Task>> _tasksList;

  @override
  Widget build(BuildContext context) {
    _tasksList = Provider.of<Repository>(context).getTasks();
    return Container(
      child: StreamBuilder<List<Task>>(
        stream: _tasksList,
        builder: (BuildContext context, AsyncSnapshot<List<Task>> snapshot) {
          if (snapshot.hasError) {
            return ErrorScreen(error: snapshot.error.toString());
          }

          switch (snapshot.connectionState) {
            case ConnectionState.waiting:
              return LoadingScreen();
            default:
              return AllTasks(tasks: snapshot.data);
          }
        },
      ),
    );
  }
}
