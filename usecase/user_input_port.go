package usecase

import "github.com/labstack/echo"

type UsersInputPort interface {
	AddUsers(ec echo.Context) error
}
