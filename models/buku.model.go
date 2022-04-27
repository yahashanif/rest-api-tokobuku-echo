package models

import (
	"fmt"
	"net/http"
	"rest-api-tokobuku-echo/db"
	"rest-api-tokobuku-echo/migrations"
)

// var d *gorm.DB

func FetchAllBuku() (Response, error) {
	// var obj migrations.Buku

	var arrObj []Buku

	var res Response
	d := db.CreateCon()
	d.Raw("SELECT bukus.*,kategoris.* from bukus join kategoris on bukus.id_kategori = kategoris.id").Scan(&arrObj)

	res.Status = http.StatusOK
	res.Message = "Sukses"
	res.Data = arrObj

	return res, nil

}

func FindBukuByID(ID int) (Response, error) {
	var obj Buku

	var res Response

	d := db.CreateCon()

	err := d.Find(&obj, ID).Error

	if err != nil {
		fmt.Println(err)
		return res, err
	}

	res.Status = http.StatusOK
	res.Message = "Sukses"
	res.Data = obj

	return res, nil
}

func CreateBuku(buku migrations.Buku) (Response, error) {
	var res Response

	d := db.CreateCon()

	err := d.Create(&buku).Error

	if err != nil {
		fmt.Println(err)
		return res, err
	}
	res.Status = http.StatusOK
	res.Message = "Sukses"
	res.Data = buku

	return res, err
}

func UpdateBuku(ID int, buku migrations.Buku) (Response, error) {
	var res Response

	d := db.CreateCon()

	err := d.Model(migrations.Buku{}).Where("id = ?", ID).Updates(buku).Error

	if err != nil {
		fmt.Println(err)
		return res, err
	}

	res.Status = http.StatusOK
	res.Message = "SUKSES"
	res.Data = map[string]string{
		"Message": "Berhasil Update Data ",
	}

	return res, err
}

func DeleteBuku(ID int) (Response, error) {
	var res Response
	var buku migrations.Buku
	d := db.CreateCon()
	err := d.Where("id = ?", ID).Delete(&buku).Error

	if err != nil {
		fmt.Println(err)
		return res, err
	}

	res.Status = http.StatusOK
	res.Message = "SUKSES"
	res.Data = map[string]string{
		"Message": "Sukses Delete Data",
	}

	return res, err
}
