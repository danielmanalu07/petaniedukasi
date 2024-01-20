import 'package:flutter/material.dart';
import 'package:get/get.dart';
import 'package:petaniedukasi/controllers/userController.dart';

class HomeUser extends StatefulWidget {
  const HomeUser({super.key});

  @override
  State<HomeUser> createState() => _HomeUserState();
}

class _HomeUserState extends State<HomeUser> {
  UserController userController = Get.put(UserController());
  @override
  Widget build(BuildContext context) {
    return Scaffold(
      body: Center(
        child: Column(
          mainAxisAlignment: MainAxisAlignment.center,
          children: [
            IconButton(
              onPressed: () async {
                userController.Logout();
              },
              icon: Icon(
                Icons.exit_to_app,
                color: Colors.blue,
              ),
            ),
            Text("Logout"),
          ],
        ),
      ),
    );
  }
}
