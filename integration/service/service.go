package service

import (
	"errors"
	"fmt"
	"go-tests/integration"
)

var ErrCreatingUsers = errors.New("error creating users")

type UserService struct {
	usersRepo integration.UsersRepo
}

func NewUserService(userRepo integration.UsersRepo) *UserService {
	return &UserService{usersRepo: userRepo}
}

func (s *UserService) CreateUsers(users []integration.User) error {
	err := s.usersRepo.CreateUsers(users)
	if err != nil {
		return fmt.Errorf("%w: %v", ErrCreatingUsers, err)
	}

	return nil
}
