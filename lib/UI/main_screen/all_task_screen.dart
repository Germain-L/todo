import 'package:flutter/material.dart';
import 'package:provider/provider.dart';
import 'package:todos/UI/all_tasks/task_UI.dart';
import 'package:todos/data/repository/task_repo.dart';
import 'package:todos/data/task.dart';

class AllTasks extends StatelessWidget {
  final List<Task> tasks;

  AllTasks({Key key, @required this.tasks}) : super(key: key);

  @override
  Widget build(BuildContext context) {
    return Container(
      child: ListView.builder(
        itemCount: tasks.length,
        itemBuilder: (context, index) {
          return TaskUI(
            task: tasks[index],
          );
        },
      ),
    );
  }
}
