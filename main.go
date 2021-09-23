package main

import (
	"myproperty-api/config"
	m "myproperty-api/middlewares"
	"myproperty-api/routes"
)

func main() {
	config.InitDB()
	e := routes.New()
	m.LogMiddlewares(e)
	e.Logger.Fatal(e.Start(":8080"))
}
