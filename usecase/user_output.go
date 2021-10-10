package usecase

import (
	"example/domain"
	"github.com/labstack/echo"
)

type UserOutput interface {
	GetUser(ec echo.Context, user []*domain.User) error
}
