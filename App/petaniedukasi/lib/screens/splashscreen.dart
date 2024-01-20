import 'package:flutter/material.dart';
import 'package:get/get.dart';
import 'package:petaniedukasi/screens/admin/login.dart';
import 'package:petaniedukasi/screens/user/login.dart';

class SplashScreen extends StatefulWidget {
  const SplashScreen({Key? key}) : super(key: key);

  @override
  State<SplashScreen> createState() => _SplashScreenState();
}

class _SplashScreenState extends State<SplashScreen> {
  @override
  Widget build(BuildContext context) {
    return Scaffold(
      body: SingleChildScrollView(
        child: Container(
          width: MediaQuery.of(context).size.width,
          child: Padding(
            padding: EdgeInsets.symmetric(vertical: 100),
            child: Column(
              crossAxisAlignment: CrossAxisAlignment.center,
              children: [
                Image(
                  image: AssetImage("assets/images/Splash.png"),
                  width: MediaQuery.of(context).size.width,
                ),
                SizedBox(
                  height: 30,
                ),
                Text(
                  "Petani - Edukasi",
                  style: TextStyle(
                    fontSize: 20,
                    fontWeight: FontWeight.w600,
                    color: Color(0xFF76B258),
                    letterSpacing: 2,
                  ),
                ),
                SizedBox(
                  height: 30,
                ),
                Text(
                  "Kembangkan Pertanianmu \n      dengan Cerdas dan \n    Terhubung langsung🌱",
                  style: TextStyle(
                    fontSize: 25,
                    fontWeight: FontWeight.w600,
                    color: Colors.black,
                  ),
                ),
                SizedBox(
                  height: 30,
                ),
                ElevatedButton(
                  onPressed: () {
                    Get.to(LoginUser());
                  },
                  style: ElevatedButton.styleFrom(
                    primary: Color(0xFF76B258),
                    onPrimary: Colors.white,
                    padding: EdgeInsets.symmetric(vertical: 16),
                    textStyle: TextStyle(
                      fontSize: 20,
                      fontWeight: FontWeight.w600,
                    ),
                    minimumSize: Size(
                      MediaQuery.of(context).size.width - 100,
                      0,
                    ),
                  ),
                  child: Text("SIGN IN USER"),
                ),
                SizedBox(
                  height: 20,
                ),
                ElevatedButton(
                  onPressed: () {
                    Get.to(LoginAdmin());
                  },
                  style: ElevatedButton.styleFrom(
                    primary: Color(0xFF76B258),
                    onPrimary: Colors.white,
                    padding: EdgeInsets.symmetric(vertical: 16),
                    textStyle: TextStyle(
                      fontSize: 20,
                      fontWeight: FontWeight.w600,
                    ),
                    minimumSize: Size(
                      MediaQuery.of(context).size.width - 100,
                      0,
                    ),
                  ),
                  child: Text("SIGN IN ADMIN"),
                ),
              ],
            ),
          ),
        ),
      ),
    );
  }
}
