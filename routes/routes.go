package routes

import (
	"rest-api-tokobuku-echo/controllers"

	"github.com/labstack/echo/v4"
)

func Init() *echo.Echo {
	e := echo.New()

	// Buku
	e.GET("/buku", controllers.FetchAllBuku)
	e.GET("/buku/:id", controllers.FindBukuByID)
	e.POST("/buku", controllers.CreateBuku)
	e.PUT("/buku", controllers.UpdateBuku)
	e.DELETE("/buku/:id", controllers.DeleteBuku)

	e.POST("/register", controllers.RegisterUser)

	return e
}
