package main

import (
	"example/interface/controllers"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func main()  {
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	usersController := controllers.NewUsersController(e)
	e.POST("/users", usersController.POST)

	e.Logger.Fatal(e.Start(":1323"))
}
