import 'package:flutter/material.dart';
import 'package:provider/provider.dart';

import '../../data/repository/task_repo.dart';
import '../../data/task.dart';

class TaskUI extends StatefulWidget {
  final Task task;

  const TaskUI({Key key, this.task}) : super(key: key);

  @override
  _TaskUIState createState() => _TaskUIState();
}

class _TaskUIState extends State<TaskUI> {
  Function _deleteTask;
  Function _markTaskDone;
  Function _unmarkTaskDone;
  @override
  Widget build(BuildContext context) {
    _deleteTask = Provider.of<Repository>(context).deleteTask;
    _markTaskDone = Provider.of<Repository>(context).markTaskDone;
    _unmarkTaskDone = Provider.of<Repository>(context).unmarkTaskDone;
    return ListTile(
      title: Text(
        widget.task.data,
        style: TextStyle(
          fontSize: 30,
          decoration: widget.task.done ? TextDecoration.lineThrough : TextDecoration.none,
        ),
      ),
      subtitle: Text(widget.task.time()),
      trailing: Row(
        mainAxisSize: MainAxisSize.min,
        children: [
          IconButton(
            icon: Icon(Icons.delete, color: Colors.red,),
            onPressed: () => setState(() {
              _deleteTask(widget.task);
            }),
          ),
          IconButton(
            icon: Icon(widget.task.done ? Icons.undo : Icons.check, color: Colors.green,),
            onPressed: () {
              if(widget.task.done) {
                _unmarkTaskDone(widget.task);
              } else {
                _markTaskDone(widget.task);
              }
            },
          ),
        ],
      ),
    );
  }
}
