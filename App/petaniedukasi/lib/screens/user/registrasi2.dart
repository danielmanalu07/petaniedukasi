import 'package:flutter/material.dart';
import 'package:get/get.dart';
import 'package:image_picker/image_picker.dart';

class Registrasi2Screen extends StatefulWidget {
  const Registrasi2Screen({super.key});

  @override
  State<Registrasi2Screen> createState() => _Registrasi2ScreenState();
}

class _Registrasi2ScreenState extends State<Registrasi2Screen> {
  late ImagePicker _imagePicker;
  late PickedFile _image;

  @override
  void initState() {
    super.initState();
    _imagePicker = ImagePicker();
  }

  Future<void> getImage(ImageSource source) async {
    try {
      PickedFile image =
          (await _imagePicker.pickImage(source: source)) as PickedFile;
      setState(() {
        _image = image;
      });
    } catch (e) {
      print('Error picking image: $e');
    }
  }

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      appBar: AppBar(),
      body: SingleChildScrollView(
        child: Container(
          width: MediaQuery.of(context).size.width,
          padding: EdgeInsets.symmetric(horizontal: 40, vertical: 20),
          child: Column(
            crossAxisAlignment: CrossAxisAlignment.start,
            children: [
              Text(
                "Personal Data",
                style: TextStyle(
                  fontSize: 30,
                  fontWeight: FontWeight.w700,
                  color: Colors.black,
                ),
              ),
              SizedBox(
                height: 30,
              ),
              Text(
                "Silahkan isi form dibawah ini\nuntuk memudahkan kami\nmengenal Anda",
                style: TextStyle(
                  fontSize: 20,
                  fontWeight: FontWeight.w400,
                  color: Colors.black,
                ),
              ),
              SizedBox(
                height: 30,
              ),
              TextFormField(
                keyboardType: TextInputType.number,
                decoration: InputDecoration(
                  labelText: "Umur",
                  border: OutlineInputBorder(
                    borderRadius: BorderRadius.circular(10),
                  ),
                ),
              ),
              SizedBox(
                height: 30,
              ),
              Text(
                "Foto KTP",
                style: TextStyle(
                  fontSize: 20,
                  color: Colors.black,
                  fontWeight: FontWeight.bold,
                ),
              ),
              InkWell(
                onTap: () {
                  getImage(ImageSource.gallery);
                },
                child: Container(
                  width: MediaQuery.of(context).size.width - 250,
                  height: MediaQuery.of(context).size.height - 800,
                  padding: EdgeInsets.symmetric(vertical: 16, horizontal: 10),
                  color: Colors.cyan,
                  child: Icon(Icons.camera_alt),
                ),
              ),
              SizedBox(
                height: 30,
              ),
              Text(
                "Foto Selfie",
                style: TextStyle(
                  fontSize: 20,
                  color: Colors.black,
                  fontWeight: FontWeight.bold,
                ),
              ),
              InkWell(
                onTap: () {
                  getImage(ImageSource.gallery);
                },
                child: Container(
                  width: MediaQuery.of(context).size.width - 250,
                  height: MediaQuery.of(context).size.height - 800,
                  padding: EdgeInsets.symmetric(vertical: 16, horizontal: 10),
                  color: Colors.cyan,
                  child: Icon(Icons.camera_alt),
                ),
              ),
              SizedBox(
                height: 50,
              ),
              InkWell(
                onTap: () {
                  Get.to(Registrasi2Screen());
                },
                child: Container(
                  width: MediaQuery.of(context).size.width,
                  padding: EdgeInsets.symmetric(vertical: 16),
                  color: Color(0xFF76B258),
                  child: Center(
                    child: Text(
                      "Simpan",
                      style: TextStyle(
                        fontSize: 20,
                        fontWeight: FontWeight.w600,
                        color: Colors.white,
                      ),
                    ),
                  ),
                ),
              ),
            ],
          ),
        ),
      ),
    );
  }
}
