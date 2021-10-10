package usecases

import (
	"log"

	"example/domain"
	"example/usecase/services"

	"github.com/labstack/echo"
)

type UsersInteractor struct {
	UsersOutputPort UsersOutputPort
	UsersService services.UsersService
}

func NewUsersInteractor(
	usersOutputPort UsersOutputPort,
	usersService services.UsersService,
) *UsersInteractor {
	return &UsersInteractor{
		UsersOutputPort: usersOutputPort,
		UsersService: usersService,
	}
}

func (i *UsersInteractor) AddUsers(ec echo.Context) error {
	u := new(domain.User)
	if err := ec.Bind(u); err != nil {
		return err
	}

	res, err := i.UsersService.AddUsers(ec, u)
	if err != nil {
		log.Fatal(err)
	}

	return i.UsersOutputPort.AddUsers(ec, res)
}
