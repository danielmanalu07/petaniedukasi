import 'dart:convert';

import 'package:flutter/material.dart';
import 'package:get/get.dart';
import 'package:get_storage/get_storage.dart';
import 'package:image_picker/image_picker.dart';
import 'package:petaniedukasi/databases/constants.dart';
import 'package:petaniedukasi/models/user.dart';
import 'package:http/http.dart' as http;
import 'package:petaniedukasi/screens/user/home.dart';
import 'package:petaniedukasi/screens/user/login.dart';

class UserController extends GetxController {
  final isLoading = false.obs;

  User? userData;
  final box = GetStorage();

  Future<void> Register({
    required String name,
    required String address,
    required String age,
    required String hp,
    required String email,
    required String password,
    required XFile ktpFile,
    required XFile fotoFile,
    required String status,
  }) async {
    try {
      isLoading.value = true;

      var user = {
        'name': name,
        'address': address,
        'age': age,
        'hp': hp,
        'email': email,
        'password': password,
        'ktp': ktpFile.path,
        'foto': fotoFile.path,
        'status': status,
      };

      var response = await http.post(
        Uri.parse(UrlUser + 'register'),
        body: user,
      );

      if (response.statusCode == 200) {
        isLoading.value = false;

        Get.offAll(() => LoginUser());
        Get.snackbar(
          'Success',
          jsonDecode(response.body)['message'],
          snackPosition: SnackPosition.TOP,
          backgroundColor: Colors.green,
          colorText: Colors.white,
        );
        print("Success :  ${response.statusCode}");
        print(response.body);
      } else {
        isLoading.value = false;
        Get.snackbar(
          'Error',
          jsonDecode(response.body)['message'],
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

  Future<void> Login({
    required String email,
    required String password,
  }) async {
    try {
      isLoading.value = true;
      var user = {
        'email': email,
        'password': password,
      };

      var response = await http.post(
        Uri.parse(UrlUser + 'login'),
        body: jsonEncode(user),
        headers: {'Content-Type': 'application/json'},
      );

      if (response.statusCode == 200) {
        isLoading.value = false;
        Get.offAll(() => const HomeUser());
        Get.snackbar(
          'Success',
          'Login successful',
          snackPosition: SnackPosition.TOP,
          backgroundColor: Colors.green,
          colorText: Colors.white,
        );

        print("Success :  ${response.statusCode}");
        print(response.body);
      } else if (response.statusCode == 404) {
        isLoading.value = false;
        Get.snackbar(
          'Error',
          'User Not Found.',
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
      Get.offAll(() => LoginUser());
    } catch (e) {
      print(e.toString());
    }
  }
}
