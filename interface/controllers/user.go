package controllers

import (
	"example/interface/presenters"
	"example/interface/repositories"
	"example/usecase"
	"github.com/labstack/echo"
)

type UserController struct {
	UserInput usecase.UserInput
}

func NewUserController(e *echo.Echo) *UserController {
	return &UserController{
		UserInput: usecase.NewUserInteractor(
			e,
			presenters.NewUserPresenters(e),
			repositories.NewUserRepositoryAdapter(e),
		),
	}
}

func (c *UserController) Get(ec echo.Context) error {
	return c.UserInput.GetUser(ec)
}
