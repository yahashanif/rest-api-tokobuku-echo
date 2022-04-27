package main

import (
	"rest-api-tokobuku-echo/db"
	"rest-api-tokobuku-echo/routes"
)

func main() {
	db.Init()
	e := routes.Init()

	e.Logger.Fatal(e.Start(":8080"))
}
