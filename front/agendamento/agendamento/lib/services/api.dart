import 'package:dio/dio.dart';
getHttp() async {
  try {
    var response = await Dio().get('http://www.google.com');
    print(response);
    return response;
  } catch (e) {
    print(e);
    return e;
  }
}