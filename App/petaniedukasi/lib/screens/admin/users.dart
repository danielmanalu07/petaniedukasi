import 'package:flutter/material.dart';

class ApproveUsers extends StatefulWidget {
  const ApproveUsers({Key? key}) : super(key: key);

  @override
  State<ApproveUsers> createState() => _ApproveUsersState();
}

class _ApproveUsersState extends State<ApproveUsers> {
  @override
  Widget build(BuildContext context) {
    return Scaffold(
      appBar: AppBar(
        backgroundColor: Colors.blueAccent,
        title: Center(
          child: Text(
            "Users Account",
            style: TextStyle(
              color: Colors.white,
              fontSize: 25,
              fontWeight: FontWeight.w600,
            ),
          ),
        ),
      ),
      body: Container(
        padding: EdgeInsets.symmetric(vertical: 20),
        width: MediaQuery.of(context).size.width,
        child: Row(
          mainAxisAlignment: MainAxisAlignment.center,
          children: [
            Card(
              color: Colors.grey,
              child: SizedBox(
                width: MediaQuery.of(context).size.width - 80,
                child: Padding(
                  padding: const EdgeInsets.only(left: 130),
                  child: Text(
                    "Contoh",
                    style: TextStyle(color: Colors.white),
                  ),
                ),
              ),
            ),
          ],
        ),
      ),
    );
  }
}
