package main

import (
	"github.com/labstack/echo/v4"
	"github.com/praktikum/Modul_Action/controllers"
)

func main() {
	e := echo.New()

	e.GET("/users", controllers.GetUsers) //Read
	e.POST("/users", controllers.InsertUsers) //Create

	e.Logger.Fatal(e.Start(":8080"))
}