package services

import (
	"example/domain"

	"github.com/labstack/echo"
)

type UsersService interface {
	AddUsers(ec echo.Context, user *domain.User) ([]*domain.User, error)
}
