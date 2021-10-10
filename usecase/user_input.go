package usecase

import "github.com/labstack/echo"

type UserInput interface {
	GetUser(ec echo.Context) error
}
