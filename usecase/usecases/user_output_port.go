package usecases

import (
	"example/domain"

	"github.com/labstack/echo"
)

type UsersOutputPort interface {
	AddUsers(ec echo.Context, user []*domain.User) error
}
