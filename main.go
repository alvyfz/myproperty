package main

import (
	"myproperty-api/config"
	route "myproperty-api/routes"

	"github.com/labstack/echo"
)

func main() {
	config.InitDB()
	app := echo.New()
	route.NewUser(app)
	app.Logger.Fatal(app.Start(":8080"))
}
