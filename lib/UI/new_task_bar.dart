import 'package:flutter/material.dart';

class NewTaskAppBar extends StatelessWidget with PreferredSizeWidget {
  NewTaskAppBar({
    Key key,
  })  : preferredSize = Size(double.infinity, 100),
        super(key: key);

  @override
  Widget build(BuildContext context) {
    return SizedBox(
      height: 100,
      width: double.maxFinite,
      child: Row(
        mainAxisSize: MainAxisSize.min,
        mainAxisAlignment: MainAxisAlignment.start,
        children: <Widget>[
          Padding(
            padding: const EdgeInsets.all(8.0),
            child: TextField(
              decoration: InputDecoration(
                border: OutlineInputBorder(),
              ),
            ),
          ),
          IconButton(
            icon: Icon(Icons.check),
            onPressed: () {},
          ),
        ],
      ),
    );
  }

  @override
  final Size preferredSize;
}
