package services

import (
	"example/domain"
)

type UsersService interface {
	AddUsers(user *domain.User) ([]*domain.User, error)
}
