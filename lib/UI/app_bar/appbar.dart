import 'package:flutter/material.dart';
import 'package:provider/provider.dart';
import 'package:todos/data/repository/task_repo.dart';

class NewTaskUI extends StatefulWidget with PreferredSizeWidget {
  @override
  _NewTaskUIState createState() => _NewTaskUIState();

  @override
  final Size preferredSize = Size.fromHeight(60);
}

class _NewTaskUIState extends State<NewTaskUI> {
  Function _addTask;
  String task;

  @override
  void initState() {
    super.initState();
    task = "";
  }

  @override
  Widget build(BuildContext context) {
    _addTask = Provider.of<Repository>(context).createTask;
    return AppBar(
      title: TextField(
        onChanged: (String change) => setState(() {
          task = change;
        }),
      ),
      centerTitle: false,
      elevation: 0,
      backgroundColor: Theme.of(context).primaryColor,
      actions: [
        IconButton(
          icon: Icon(Icons.add),
          onPressed: () async {
            await _addTask(task);
          },
        ),
      ],
    );
  }
}
