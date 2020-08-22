import 'package:cloud_firestore/cloud_firestore.dart';
import 'package:flutter/foundation.dart';

class Task {
  final String data;
  final String id;
  final Timestamp timeCreated;
  final bool done;

  String time() {
    DateTime dateTime = timeCreated.toDate();
    String minute = dateTime.minute.toString().length == 1 ? "0${dateTime.minute.toString()}" : dateTime.minute.toString();
    String hour = dateTime.hour.toString().length == 1 ? "0${dateTime.hour.toString()}" : dateTime.hour.toString();
    String day = dateTime.day.toString();
    String month = dateTime.month.toString();
    String year = dateTime.year.toString();

    return "$day/$month/$year at  $hour:$minute";
  }

  Task({
    @required this.done,
    @required this.timeCreated,
    @required this.id,
    @required this.data,
  })  : assert(data != null),
        assert(id != null),
        assert(timeCreated != null),
        assert(done != null);
}
