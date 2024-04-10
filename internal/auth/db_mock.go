package auth

import (
	"github.com/stretchr/testify/mock"
)

// DbMock is a mock for the UserDB interface
type DbMock struct {
	mock.Mock
}

func (m *DbMock) CreateUser(user *User) error {
	args := m.Called(user)
	return args.Error(0)
}

func (m *DbMock) GetUserByEmail(email string) (*User, error) {
	args := m.Called(email)
	return args.Get(0).(*User), args.Error(1)
}
