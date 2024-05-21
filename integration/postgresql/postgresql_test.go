//go:build integration

package postgresql

import (
	_ "github.com/lib/pq"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"go-tests/integration"
	"testing"
)

func TestNewPostgresqlRepo(t *testing.T) {
	_, err := NewPostgresqlRepo()

	assert.NoError(t, err)
}

type TestSuite struct {
	suite.Suite
	repo *UsersPostgreSqlRepo
}

func (suite *TestSuite) SetupSuite() {
	repo, err := NewPostgresqlRepo()
	if err != nil {
		suite.FailNow(err.Error())
	}

	suite.repo = repo
}

func (suite *TestSuite) TearDownSuite() {
	suite.repo.db.Close()
}

func (suite *TestSuite) TearDownTest() {
	_, err := suite.repo.db.Exec("DELETE FROM users")
	if err != nil {
		suite.FailNow(err.Error())
	}
}

func TestIntegrationSuite(t *testing.T) {
	suite.Run(t, new(TestSuite))
}

func (suite *TestSuite) TestCreateUser() {
	tests := []struct {
		name             string
		users            []integration.User
		expectedUsersLen int
		expectedUsers    []integration.User
		expectedErr      error
	}{
		{
			name:             "Create one users",
			users:            []integration.User{{ID: 1, Name: "John"}},
			expectedUsersLen: 1,
			expectedUsers:    []integration.User{{ID: 1, Name: "John"}},
			expectedErr:      nil,
		},
		{
			name:             "Create multiple users",
			users:            []integration.User{{ID: 1, Name: "John"}, {ID: 2, Name: "Doe"}},
			expectedUsersLen: 2,
			expectedUsers:    []integration.User{{ID: 1, Name: "John"}, {ID: 2, Name: "Doe"}},
			expectedErr:      nil,
		},
	}

	for _, tt := range tests {
		suite.T().Run(tt.name, func(t *testing.T) {
			err := suite.repo.CreateUsers(tt.users)
			assert.ErrorIs(t, err, tt.expectedErr)

			// count
			var count int

			err = suite.repo.db.QueryRow("SELECT COUNT(*) FROM users").Scan(&count)
			assert.Equal(t, tt.expectedUsersLen, count)

			// find all
			res, err := suite.repo.db.Query("SELECT id, name FROM users")
			assert.NoError(t, err)

			var users []integration.User
			for res.Next() {
				var user integration.User
				err := res.Scan(&user.ID, &user.Name)
				assert.NoError(t, err)
				users = append(users, user)
			}

			assert.Equal(t, tt.expectedUsers, users)
			suite.TearDownTest()
		})
	}
}
