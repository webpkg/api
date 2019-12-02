package repository

import (
	"sync"

	"github.com/webpkg/api/contract"
	"github.com/webpkg/api/model"
)

var (
	_userRepository     contract.UserRepository
	_onceUserRepository sync.Once
)

// CreateUserRepository return contract.IUserRepository
func CreateUserRepository() contract.UserRepository {

	_onceUserRepository.Do(func() {
		_userRepository = &UserRepository{}
	})

	return _userRepository
}

// UserRepository struct
type UserRepository struct {
}

// Find user by id
func (u *UserRepository) Find(id uint64) (*model.User, error) {
	return nil, nil
}

// Create user
func (u *UserRepository) Create(user *model.User) error {
	return nil
}
