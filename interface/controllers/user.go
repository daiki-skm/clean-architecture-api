package controllers

import (
	"example/interface/presenters"
	"example/interface/repositories"
	"example/usecase/usecases"

	"github.com/labstack/echo"
)

type UsersController struct {
	UsersInputPort usecases.UsersInputPort
}

func NewUsersController(e *echo.Echo) *UsersController {
	return &UsersController{
		UsersInputPort: usecases.NewUsersInteractor(
			presenters.NewUsersPresenters(e),
			repositories.NewUsersRepository(e),
		),
	}
}

func (c *UsersController) POST(ec echo.Context) error {
	return c.UsersInputPort.AddUsers(ec)
}
