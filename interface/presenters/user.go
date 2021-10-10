package presenters

import (
	"net/http"

	"example/domain"
	"example/usecase/usecases"

	"github.com/labstack/echo"
)

type UsersPresenters struct {
	echo *echo.Echo
	usecases.UsersOutputPort
}

func NewUsersPresenters(echo *echo.Echo) *UsersPresenters {
	return &UsersPresenters{
		echo: echo,
	}
}

func (p *UsersPresenters) AddUsers(ec echo.Context, user []*domain.User) error {
	return ec.JSON(http.StatusOK, user)
}
