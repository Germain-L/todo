import 'package:flutter/material.dart';
import 'package:provider/provider.dart';
import 'package:todos/data/repository/task_repo.dart';
import 'package:todos/data/task.dart';

class TaskUI extends StatefulWidget {
  final Task task;

  const TaskUI({Key key, this.task}) : super(key: key);

  @override
  _TaskUIState createState() => _TaskUIState();
}

class _TaskUIState extends State<TaskUI> {
  bool toDelete = false;

  Function _deleteTask;
  @override
  Widget build(BuildContext context) {
    _deleteTask = Provider.of<Repository>(context).deleteTask;
    return Column(
      mainAxisSize: MainAxisSize.min,
      children: <Widget>[
        Padding(
          padding: const EdgeInsets.symmetric(horizontal: 15.0),
          child: Row(
            mainAxisAlignment: MainAxisAlignment.spaceBetween,
            children: <Widget>[
              SizedBox(
                width: MediaQuery.of(context).size.width * 0.68,
                child: Text(
                  widget.task.data,
                  style: TextStyle(
                      fontSize: 30,
                      decoration: toDelete ? TextDecoration.lineThrough : null),
                ),
              ),
              IconButton(
                icon: Icon(Icons.delete),
                onPressed: () => setState(() {
                  _deleteTask(widget.task);
                }),
              ),
              IconButton(
                icon: Icon(toDelete ? Icons.undo : Icons.check),
                onPressed: () => setState(() {
                  toDelete = !toDelete;
                }),
              ),
            ],
          ),
        )
      ],
    );
  }
}
