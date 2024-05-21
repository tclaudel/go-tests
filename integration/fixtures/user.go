package fixtures

import (
	"go-tests/integration"
	"testing"
)

type UserOptType = func(user *integration.User)

type UserOpt struct{}

func (UserOpt) ID(t *testing.T, id int) UserOptType {
	t.Helper()

	return func(user *integration.User) {
		user.ID = id
	}
}

func (UserOpt) Name(t *testing.T, name string) UserOptType {
	t.Helper()

	return func(user *integration.User) {
		user.Name = name
	}
}

func User(t *testing.T, opts ...UserOptType) integration.User {
	t.Helper()

	user := &integration.User{
		ID:   1,
		Name: "John",
	}

	for _, opt := range opts {
		opt(user)
	}

	return *user
}
