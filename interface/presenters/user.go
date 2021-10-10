package presenters

import (
	"example/domain"
	"example/usecase"
	"github.com/labstack/echo"
	"net/http"
)

type UserPresenters struct {
	echo *echo.Echo
	usecase.UserOutput
}

func NewUserPresenters(echo *echo.Echo) *UserPresenters {
	return &UserPresenters{
		echo: echo,
	}
}

func (p *UserPresenters) GetUser(ec echo.Context, user []*domain.User) error {
	return ec.JSON(http.StatusOK, user)
}
