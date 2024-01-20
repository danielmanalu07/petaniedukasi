import 'package:flutter/material.dart';
import 'package:petaniedukasi/screens/admin/dashboard.dart';
import 'package:petaniedukasi/screens/admin/profile.dart';
import 'package:petaniedukasi/screens/admin/users.dart';

class HomeAdmin extends StatefulWidget {
  const HomeAdmin({super.key});

  @override
  State<HomeAdmin> createState() => _HomeAdminState();
}

class _HomeAdminState extends State<HomeAdmin> {
  int pageIndex = 0;

  final pages = [
    const DashboardAdmin(),
    const ApproveUsers(),
    const ProfileAdmin(),
  ];

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      bottomNavigationBar: Container(
        height: 60,
        decoration: BoxDecoration(
            color: Theme.of(context).primaryColor,
            borderRadius: const BorderRadius.only(
              topLeft: Radius.circular(20),
              topRight: Radius.circular(20),
            )),
        child: Row(
          mainAxisAlignment: MainAxisAlignment.spaceAround,
          children: [
            IconButton(
              onPressed: () {
                setState(() {
                  pageIndex = 0;
                });
              },
              icon: pageIndex == 0
                  ? const Icon(
                      Icons.home_outlined,
                      color: Colors.orange,
                      size: 35,
                    )
                  : const Icon(
                      Icons.home_outlined,
                      color: Colors.white,
                      size: 35,
                    ),
            ),
            IconButton(
              onPressed: () {
                setState(() {
                  pageIndex = 1;
                });
              },
              icon: pageIndex == 1
                  ? const Icon(
                      Icons.group_outlined,
                      color: Colors.orange,
                      size: 35,
                    )
                  : const Icon(
                      Icons.group_outlined,
                      color: Colors.white,
                      size: 35,
                    ),
            ),
            IconButton(
              onPressed: () {
                setState(() {
                  pageIndex = 2;
                });
              },
              icon: pageIndex == 2
                  ? const Icon(
                      Icons.person,
                      color: Colors.orange,
                      size: 35,
                    )
                  : const Icon(
                      Icons.person,
                      color: Colors.white,
                      size: 35,
                    ),
            ),
          ],
        ),
      ),
      drawer: Drawer(),
      backgroundColor: const Color(0xffC4DFCB),
      body: pages[pageIndex],
    );
  }
}
