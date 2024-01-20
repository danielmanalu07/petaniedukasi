import 'dart:convert';
import 'package:flutter/material.dart';
import 'package:get/get.dart';
import 'package:get_storage/get_storage.dart';
import 'package:petaniedukasi/controllers/adminController.dart';
import 'package:petaniedukasi/databases/constants.dart';
import 'package:petaniedukasi/models/admin.dart';
import 'package:http/http.dart' as http;

class ProfileAdmin extends StatefulWidget {
  const ProfileAdmin({Key? key}) : super(key: key);

  @override
  State<ProfileAdmin> createState() => _ProfileAdminState();
}

class _ProfileAdminState extends State<ProfileAdmin> {
  final AdminController adminController = Get.put(AdminController());
  Admin? adminData;
  final box = GetStorage();
  final isLoading = true.obs;

  @override
  void initState() {
    super.initState();
    fetchDataAdmin();
  }

  Future<void> fetchDataAdmin() async {
    isLoading(true);
    String? authToken = box.read('token');
    try {
      final response = await http.get(
        Uri.parse(UrlAdmin + 'admin'),
        headers: {
          'Cookie': 'jwt=$authToken',
          'Authorization': 'Bearer $authToken',
          'Content-Type': 'application/json',
          'Access-Control-Allow-Credentials': 'true',
        },
      );

      if (response.statusCode == 200) {
        Map<String, dynamic>? responseData = jsonDecode(response.body);
        if (responseData != null && responseData['Admin'] != null) {
          setState(() {
            adminData = Admin.fromJson(responseData['Admin']);
          });
        } else {
          throw Exception('Invalid response format');
        }
      } else if (response.statusCode == 401) {
        throw Exception("Unauthenticated");
      } else {
        throw Exception('Failed to load admin data: ${response.statusCode}');
      }
    } catch (error) {
      print(error);
    } finally {
      isLoading(false);
    }
  }

  AlertDialog _showDialog() {
    return AlertDialog(
      title: Text("Konfirmasi"),
      content: Text("Apakah Anda ingin keluar dari aplikasi?"),
      actions: [
        TextButton(
          onPressed: () {
            Navigator.of(context).pop();
          },
          child: Text("Batal"),
        ),
        TextButton(
          onPressed: () {
            adminController.Logout();
          },
          child: Text("Keluar"),
        ),
      ],
    );
  }

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      appBar: AppBar(
        backgroundColor: Colors.blueAccent,
        title: Row(
          children: [
            Text(
              'Profile',
              style: TextStyle(
                color: Colors.white,
                fontSize: 25,
                fontWeight: FontWeight.w600,
              ),
            ),
            SizedBox(
              width: MediaQuery.of(context).size.width - 160,
            ),
            IconButton(
              onPressed: () {
                showDialog(
                  context: context,
                  builder: (BuildContext context) {
                    return _showDialog();
                  },
                );
              },
              icon: Icon(
                Icons.exit_to_app,
                size: 30,
              ),
              color: Colors.white,
            ),
          ],
        ),
      ),
      body: Obx(() {
        return isLoading.value
            ? Center(
                child: CircularProgressIndicator(),
              )
            : Center(
                child: Padding(
                  padding: const EdgeInsets.all(16.0),
                  child: Column(
                    mainAxisAlignment: MainAxisAlignment.center,
                    crossAxisAlignment: CrossAxisAlignment.center,
                    children: [
                      CircleAvatar(
                        radius: 50,
                        backgroundColor: Colors.grey,
                        child: Icon(
                          Icons.person,
                          size: 60,
                          color: Colors.white,
                        ),
                      ),
                      SizedBox(height: 20),
                      Text(
                        adminData?.name ?? 'Name not Available',
                        style: TextStyle(
                          fontSize: 20,
                          fontWeight: FontWeight.bold,
                        ),
                      ),
                      SizedBox(height: 10),
                      Text(
                        adminData?.email ?? 'Email not Available',
                        style: TextStyle(
                          fontSize: 16,
                          color: Colors.grey,
                        ),
                      ),
                    ],
                  ),
                ),
              );
      }),
    );
  }
}
