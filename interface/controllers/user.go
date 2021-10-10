package controllers

import (
	"example/interface/presenters"
	"example/interface/repositories"
	"example/usecase"

	"github.com/labstack/echo"
)

type UsersController struct {
	UsersInputPort usecase.UsersInputPort
}

func NewUsersController(e *echo.Echo) *UsersController {
	return &UsersController{
		UsersInputPort: usecase.NewUsersInteractor(
			presenters.NewUsersPresenters(e),
			repositories.NewUsersRepositoryAdapter(e),
		),
	}
}

func (c *UsersController) POST(ec echo.Context) error {
	return c.UsersInputPort.AddUsers(ec)
}
