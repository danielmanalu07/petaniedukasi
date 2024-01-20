import 'dart:convert';

import 'package:flutter/material.dart';
import 'package:get/get.dart';
import 'package:get_storage/get_storage.dart';
import 'package:http/http.dart' as http;
import 'package:petaniedukasi/databases/constants.dart';
import 'package:petaniedukasi/models/admin.dart';
import 'package:petaniedukasi/screens/admin/home.dart';
import 'package:petaniedukasi/screens/admin/login.dart';

class AdminController extends GetxController {
  final isLoading = false.obs;
  final token = ''.obs;
  final box = GetStorage();

  Admin? adminData;

  Future<void> Login({
    required String email,
    required String password,
  }) async {
    try {
      isLoading.value = true;
      var admin = {
        'email': email,
        'password': password,
      };

      var response = await http.post(
        Uri.parse(UrlAdmin + 'loginadm'),
        body: jsonEncode(admin),
        headers: {'Content-Type': 'application/json'},
      );

      if (response.statusCode == 200) {
        isLoading.value = false;
        var responseBody = json.decode(response.body);
        String? tkn = responseBody['token'];

        if (tkn != null) {
          token.value = tkn;
          box.write('token', token.value);
          Get.snackbar(
            'Success',
            'Login successful',
            snackPosition: SnackPosition.TOP,
            backgroundColor: Colors.green,
            colorText: Colors.white,
          );
          Get.offAll(() => const HomeAdmin());
          print("Success :  ${response.statusCode}");
          print(response.body);
        } else {
          print('Error: JWT token is null or not found in the response body');
        }
      } else if (response.statusCode == 404) {
        isLoading.value = false;
        Get.snackbar(
          'Error',
          'Admin Not Found.',
          snackPosition: SnackPosition.TOP,
          backgroundColor: Colors.red,
          colorText: Colors.white,
        );
        print("Error :  ${response.statusCode}");
        print(response.body);
      } else if (response.statusCode == 400) {
        isLoading.value = false;
        Get.snackbar(
          'Error',
          'incorrect email or password',
          snackPosition: SnackPosition.TOP,
          backgroundColor: Colors.red,
          colorText: Colors.white,
        );
        print("Error :  ${response.statusCode}");
        print(response.body);
      } else {
        isLoading.value = false;
        Get.snackbar(
          'Error',
          'Login Failed.',
          snackPosition: SnackPosition.TOP,
          backgroundColor: Colors.red,
          colorText: Colors.white,
        );
        print("Error :  ${response.statusCode}");
        print(response.body);
      }
    } catch (e) {
      isLoading.value = false;
      print(e.toString());
    }
  }

  void Logout() async {
    try {
      box.remove('token');
      Get.snackbar(
        'Logout',
        'Logout successful',
        snackPosition: SnackPosition.TOP,
        backgroundColor: Colors.green,
        colorText: Colors.white,
      );
      Get.offAll(() => LoginAdmin());
    } catch (e) {
      print(e.toString());
    }
  }
}
