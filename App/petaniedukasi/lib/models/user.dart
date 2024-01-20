class User {
  late int id;
  late String name;
  late String address;
  late String age;
  late String hp;
  late String email;
  late String password;
  late String ktp;
  late String foto;
  late int status;

  User();

  factory User.fromJson(Map<String, dynamic> json) => User._$UserFromJson(json);
  Map<String, dynamic> toJson() => _$UserToJson(this);

  factory User._$UserFromJson(Map<String, dynamic> json) => User()
    ..id = json['id'] as int
    ..name = json['name'] as String
    ..address = json['address'] as String
    ..age = json['age'] as String
    ..hp = json['hp'] as String
    ..email = json['email'] as String
    ..password = json['password'] as String
    ..ktp = json['ktp'] as String
    ..foto = json['foto'] as String
    ..status = json['status'] as int;

  Map<String, dynamic> _$UserToJson(User instance) => <String, dynamic>{
        'id': instance.id,
        'name': instance.ktp,
        'address': instance.address,
        'age': instance.age,
        'hp': instance.hp,
        'email': instance.email,
        'password': instance.password,
        'ktp': instance.ktp,
        'foto': instance.foto,
        'status': instance.status,
      };
}
