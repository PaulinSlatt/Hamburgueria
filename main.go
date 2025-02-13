package main

import (
	"GoEcho/database"
	"GoEcho/routes"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()

	database.ConnectDB()

	routes.HandleRequest(e)

	e.Logger.Fatal(e.Start(":8082"))
}
