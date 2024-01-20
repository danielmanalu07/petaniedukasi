class Admin {
  late int id;
  late String name;
  late String email;
  late String password;

  Admin({
    required this.email,
    required this.id,
    required this.name,
    required this.password,
  });

  factory Admin.fromJson(Map<String, dynamic> json) {
    return Admin(
      id: json['id'] ?? '',
      name: json['name'] ?? '',
      email: json['Email'] ?? '',
      password: json['password'] ?? '',
    );
  }
}
