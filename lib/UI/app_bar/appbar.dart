import 'package:flutter/material.dart';
import 'package:provider/provider.dart';

import '../../data/repository/task_repo.dart';

class NewTaskUI extends StatefulWidget with PreferredSizeWidget {
  @override
  _NewTaskUIState createState() => _NewTaskUIState();

  @override
  final Size preferredSize = Size.fromHeight(60);
}

class _NewTaskUIState extends State<NewTaskUI> {
  Function _addTask;
  String task;
  TextEditingController newTaskController;

  @override
  void initState() {
    super.initState();
    task = "";
    newTaskController = TextEditingController();
  }

  @override
  Widget build(BuildContext context) {
    _addTask = Provider.of<Repository>(context).createTask;
    return AppBar(
      title: TextField(
        onChanged: (String change) => setState(() {
          task = change;
        }),
        decoration: InputDecoration(
        ),
        controller: newTaskController,
      ),
      centerTitle: false,
      elevation: 0,
      backgroundColor: Colors.white12,
      actions: [
        IconButton(
          icon: Icon(Icons.add, color: Colors.black,),
          onPressed: () async {
            bool added = await _addTask(task);
            if (added) {
              newTaskController.text = "";
            }
          },
        ),
      ],
    );
  }
}
