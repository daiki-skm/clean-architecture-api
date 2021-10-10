package usecase

import (
	"log"

	"example/domain"
	"example/interface/repositories"

	"github.com/labstack/echo"
)

type UsersInteractor struct {
	UsersOutputPort UsersOutputPort
	UsersRepository repositories.UsersRepository
}

func NewUsersInteractor(
	usersOutputPort UsersOutputPort,
	usersRepository repositories.UsersRepository,
) *UsersInteractor {
	return &UsersInteractor{
		UsersOutputPort: usersOutputPort,
		UsersRepository: usersRepository,
	}
}

func (i *UsersInteractor) AddUsers(ec echo.Context) error {
	u := new(domain.User)
	if err := ec.Bind(u); err != nil {
		return err
	}

	res, err := i.UsersRepository.AddUsers(ec, u)
	if err != nil {
		log.Fatal(err)
	}

	return i.UsersOutputPort.AddUsers(ec, res)
}
