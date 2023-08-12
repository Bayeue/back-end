# Documentation

## Bayeue

Bayeue merupakan aplikasi PPBOB berbasih Web dan Mobile menggunakan teknologi VueJs (web), Flutter (mobile), dan Go Echo (Backend).

## Teknologi

1. Echo
2. ORM Gorm
3. Claudinary
4. Xendit
5. Mailjet

## Swagger

BaseUrl = https://virtserver.swaggerhub.com/gozza/Payment-Point/1.0.0

1. masukan baseUrl
2. Buka swagger dokumentasi di https://app.swaggerhub.com/apis/gozza/Payment-Point/1.0.0
3. Masukan baseurl sesuai endpoint. Contoh : https://virtserver.swaggerhub.com/gozza/Payment-Point/1.0.0/api/product/1

## Me-run docker compose

1. buka git bash
2. lalu masukan perintah

```
    docker-compose up --build
```

3. server langsung dirun dan database langsung dibuat dalam container

## Cara Melakukan Unittest

```
go test ./users/domain/abstraction_test.go -coverpkg=./users/service/...
go test ./products/domain/abstraction_test.go -coverpkg=./products/service/...
go test ./transaction/domain/abstraction_test.go -coverpkg=./transaction/service/...

```

## cara import json postman

1. login akun postman
2. pergi ke workspace
3. import file ke workspace (silahkan cari import data dan upload)

file nya saya simpan di discord

# ERD Bayeue app

!["erd-bayeue"](./assets/erd/bayeue%20ERD.png)

## untuk endpoint beserta postman

route yang sudah ada di branch ini, kalian tinggal import ke postman!
["link-gdrive"](https://drive.google.com/file/d/1CgzMJpNxILzcdepSIJu6r9wFex1iV7IE/view?usp=sharing)
