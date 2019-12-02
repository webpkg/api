package contract

import "github.com/webpkg/api/model"

// UserRepository interface
type UserRepository interface {
	Find(id uint64) (*model.User, error)
	Create(user *model.User) error
}
