package service

import (
	"github.com/stretchr/testify/assert"
	"go-tests/integration"
	"go-tests/integration/fixtures"
	"go-tests/integration/mocks"
	"go.uber.org/mock/gomock"
	"testing"
)

func TestUserService_CreateUsers(t *testing.T) {
	userOpt := fixtures.UserOpt{}

	tests := []struct {
		name      string
		users     []integration.User
		usersRepo func(ctrl *gomock.Controller) *mocks.MockUsersRepo
		expected  error
	}{
		{
			name:  "Create users",
			users: []integration.User{fixtures.User(t)},
			usersRepo: func(ctrl *gomock.Controller) *mocks.MockUsersRepo {
				repo := mocks.NewMockUsersRepo(ctrl)
				repo.EXPECT().CreateUsers(gomock.Any()).Return(nil)
				return repo
			},
			expected: nil,
		},
		{
			name:  "Create multiple users",
			users: []integration.User{fixtures.User(t), fixtures.User(t, userOpt.Name(t, "Jane"))},
			usersRepo: func(ctrl *gomock.Controller) *mocks.MockUsersRepo {
				repo := mocks.NewMockUsersRepo(ctrl)
				repo.EXPECT().CreateUsers(gomock.Any()).Return(nil)
				return repo
			},
			expected: nil,
		},
		{
			name:  "Create users fails",
			users: []integration.User{fixtures.User(t)},
			usersRepo: func(ctrl *gomock.Controller) *mocks.MockUsersRepo {
				repo := mocks.NewMockUsersRepo(ctrl)
				repo.EXPECT().CreateUsers([]integration.User{fixtures.User(t)}).Return(assert.AnError)
				return repo
			},
			expected: ErrCreatingUsers,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			service := NewUserService(tt.usersRepo(ctrl))
			err := service.CreateUsers(tt.users)
			assert.ErrorIs(t, err, tt.expected)
		})
	}
}
