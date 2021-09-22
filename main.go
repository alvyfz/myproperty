package main

import (
	"myproperty-api/config"
	"myproperty-api/routes"
)

func main() {
	config.InitDB()
	e := routes.New()
	e.Logger.Fatal(e.Start(":8080"))
}