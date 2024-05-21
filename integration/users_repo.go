//go:generate mockgen -destination=mocks/users_repo_mock.go -package=mocks go-tests/integration UsersRepo

package integration

type UsersRepo interface {
	CreateUsers(users []User) error
	FindAllUsers() ([]User, error)
}
