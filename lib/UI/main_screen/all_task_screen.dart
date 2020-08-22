import 'package:flutter/material.dart';

import '../../data/task.dart';
import '../all_tasks/task_UI.dart';

class AllTasks extends StatelessWidget {
  final List<Task> tasks;

  AllTasks({Key key, @required this.tasks}) : super(key: key);

  @override
  Widget build(BuildContext context) {
    return Container(
      child: ListView.separated(
        separatorBuilder: (context, index) {
          return Divider();
        },
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
