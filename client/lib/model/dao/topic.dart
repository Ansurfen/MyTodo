import 'package:my_todo/model/entity/topic.dart';
import 'package:my_todo/utils/db.dart';

class TopicDao {
  static const String tableName = "topic";

  static Future<int> create(Topic t) async {
    return await DBProvider.db.insert(tableName, t.toJson());
  }

  static Future<Row?> findOne({String? where, List<Object?>? whereArgs}) async {
    Rows rows = await DBProvider.db
        .query(tableName, where: where, whereArgs: whereArgs);
    if (rows.isNotEmpty) {
      return rows[0];
    }
    return null;
  }

  static Future<List<Topic>> findMany() async {
    List<Topic> topics = [];
    Rows rows = await DBProvider.db.query(tableName);
    for (Row row in rows) {
      topics.add(Topic.fromJson(row));
    }
    return topics;
  }

  static Future<int> deleteAll() async {
    return await DBProvider.db.delete(tableName);
  }
}
