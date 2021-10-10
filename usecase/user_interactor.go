package usecase

import (
	"example/domain"
	"example/interface/repositories"
	"github.com/labstack/echo"
	"log"
)

type UserInteractor struct {
	Echo *echo.Echo
	UserOutput UserOutput
	UserRepo repositories.UserRepository
}

func NewUserInteractor(
	echo *echo.Echo,
	userOutput UserOutput,
	userRepo repositories.UserRepository,
) *UserInteractor {
	return &UserInteractor{
		Echo: echo,
		UserOutput: userOutput,
		UserRepo: userRepo,
	}
}

func (i *UserInteractor) GetUser(ec echo.Context) error {
	u := new(domain.User)
	if err := ec.Bind(u); err != nil {
		return err
	}

	res, err := i.UserRepo.GetUser(ec, u)
	if err != nil {
		log.Fatal(err)
	}

	return i.UserOutput.GetUser(ec, res)
}
