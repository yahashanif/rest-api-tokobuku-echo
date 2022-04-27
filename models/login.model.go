package models

import (
	"fmt"
	"net/http"
	"rest-api-tokobuku-echo/db"
	"rest-api-tokobuku-echo/helpers"
	"rest-api-tokobuku-echo/migrations"
)

func RegisterUser(username, password, nama string) (Response, error) {
	var res Response
	d := db.CreateCon()
	hash, _ := helpers.HashPassword(password)
	user := migrations.User{
		Username: username,
		Password: hash,
		Nama:     nama,
	}
	err := d.Create(&user).Error

	if err != nil {
		fmt.Println(err)
		return res, err
	}

	res.Status = http.StatusOK
	res.Message = "SUKSES"
	res.Data = UserResponse{
		ID:       user.ID,
		Nama:     user.Nama,
		Username: user.Username,
	}

	return res, err

}
