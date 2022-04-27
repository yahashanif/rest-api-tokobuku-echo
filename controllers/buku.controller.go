package controllers

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"rest-api-tokobuku-echo/migrations"
	"rest-api-tokobuku-echo/models"
	"strconv"

	"github.com/labstack/echo/v4"
)

func FetchAllBuku(c echo.Context) error {
	result, err := models.FetchAllBuku()

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"message": err.Error(),
		})
	}
	return c.JSON(http.StatusOK, result)
}

func FindBukuByID(c echo.Context) error {
	idString := c.Param("id")

	id, _ := strconv.Atoi(idString)
	result, err := models.FindBukuByID(id)

	if err != nil {
		c.JSON(http.StatusInternalServerError, map[string]string{
			"message": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, result)
}

func CreateBuku(c echo.Context) error {

	// Gambar
	file, err := c.FormFile("file")
	if err != nil {
		c.JSON(http.StatusInternalServerError, map[string]string{
			"message": err.Error(),
		})
	}
	src, err := file.Open()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"message": err.Error(),
		})
	}

	defer src.Close()

	// Destination

	dst, err := os.Create("files/" + file.Filename)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"message": err.Error(),
		})
	}
	defer dst.Close()
	// Copy
	if _, err = io.Copy(dst, src); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"message": err.Error(),
		})
	}

	gambar := file.Filename
	idkat, _ := strconv.Atoi(c.FormValue("id_kategori"))
	thn, _ := strconv.Atoi(c.FormValue("tahun_terbit"))
	jml, _ := strconv.Atoi(c.FormValue("jumlah_buku"))
	hrg, _ := strconv.Atoi(c.FormValue("harga"))

	// ===========================
	buku := migrations.Buku{
		Judul:       c.FormValue("judul"),
		IdKategori:  idkat,
		Pengarang:   c.FormValue("pengarang"),
		Penerbit:    c.FormValue("penerbit"),
		TahunTerbit: thn,
		JumlahBuku:  jml,
		Harga:       hrg,
		Gambar:      gambar,
		Deskripsi:   c.FormValue("deskripsi"),
	}
	result, err := models.CreateBuku(buku)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"message": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, result)
}

func UpdateBuku(c echo.Context) error {
	id, _ := strconv.Atoi(c.FormValue("ID"))
	idkat, _ := strconv.Atoi(c.FormValue("id_kategori"))
	thn, _ := strconv.Atoi(c.FormValue("id_kategori"))
	jml, _ := strconv.Atoi(c.FormValue("jumlah_buku"))
	hrg, _ := strconv.Atoi(c.FormValue("harga"))
	buku := migrations.Buku{
		Judul:       c.FormValue("judul"),
		IdKategori:  idkat,
		Pengarang:   c.FormValue("pengarang"),
		Penerbit:    c.FormValue("penerbit"),
		TahunTerbit: thn,
		JumlahBuku:  jml,
		Harga:       hrg,
		Deskripsi:   c.FormValue("deskripsi"),
	}
	result, err := models.UpdateBuku(id, buku)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"message": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, result)
}

func DeleteBuku(c echo.Context) error {
	var buku migrations.Buku
	id, _ := strconv.Atoi(c.Param("id"))
	buku.ID = id

	result, err := models.DeleteBuku(id)
	if err != nil {
		fmt.Println(err)
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"message": err.Error(),
		})
	}
	return c.JSON(http.StatusOK, result)
}
