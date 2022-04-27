package controllers

import (
	"net/http"
	"rest-api-tokobuku-echo/models"

	"github.com/labstack/echo/v4"
)

func RegisterUser(c echo.Context) error {
	username := c.FormValue("username")
	nama := c.FormValue("nama")
	password := c.FormValue("password")

	result, err := models.RegisterUser(username, password, nama)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"message": err.Error(),
		})
	}
	return c.JSON(http.StatusOK, result)
}
