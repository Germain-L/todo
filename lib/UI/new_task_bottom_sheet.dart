import 'package:flutter/material.dart';
import 'package:todos/data/task.dart';

class NewTask extends StatefulWidget {
  @override
  _NewTaskState createState() => _NewTaskState();
}

class _NewTaskState extends State<NewTask> {
  String newTaskText;
  Task newTask;

  @override
  Widget build(BuildContext context) {
    return Container(
      child: Column(
        mainAxisSize: MainAxisSize.min,
        children: [
          Row(
            mainAxisSize: MainAxisSize.max,
            mainAxisAlignment: MainAxisAlignment.center,
            children: [
              TextField(
                autocorrect: true,
                decoration: InputDecoration(),
                autofocus: true,
                onChanged: (String newText) =>
                    setState(() => newTaskText = newText),
              ),
              IconButton(
                icon: Icon(Icons.check),
                onPressed: () async {
                  newTask = Task(data: newTaskText);
                  await newTask.create();
                },
              )
            ],
          )
        ],
      ),
    );
  }
}

//TODO: add to bottomsheet