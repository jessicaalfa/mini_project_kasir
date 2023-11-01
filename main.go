package main

import (
	"kasir/config"
	"kasir/routes"
)

func main() {

	config.ConnectDB()

	e := routes.Init()

	e.Logger.Fatal(e.Start(":8000"))
}
