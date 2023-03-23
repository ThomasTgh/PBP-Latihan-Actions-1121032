package main

import (
	"github.com/labstack/echo/v4"
	"github.com/praktikum/Modul_Action/controllers"
)

func main() {
	e := echo.New()

	e.GET("/users", controllers.GetUsers) //Read

	e.Logger.Fatal(e.Start(":8080"))
}