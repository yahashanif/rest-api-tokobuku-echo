package main

import (
	"rest-api-tokobuku-echo/db"
	"rest-api-tokobuku-echo/migrations"
)

func main() {
	db.Init()
	db.CreateCon().AutoMigrate(&migrations.Buku{})
	db.CreateCon().AutoMigrate(&migrations.Cart{})
	db.CreateCon().AutoMigrate(&migrations.DetailTransaction{})
	db.CreateCon().AutoMigrate(&migrations.Diskon{})
	db.CreateCon().AutoMigrate(&migrations.Kategori{})
	db.CreateCon().AutoMigrate(&migrations.Transaction{})
	db.CreateCon().AutoMigrate(&migrations.User{})

}
