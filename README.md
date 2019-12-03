# api

a start point for your go project base on webpkg/web.

### Install
```bash
git clone git@github.com:webpkg/api.git
```
replace github.com/webpkg/api to your module path
```bash
./development.bash
```
### Controller
```go
package controller

import (
	"sync"

	"github.com/webpkg/web"
)

var (
	_userController     web.Controller
	_onceUserController sync.Once
)

// CreateUserController return web.Controller
func CreateUserController() web.Controller {

	_onceUserController.Do(func() {
		_userController = &userController{}
	})

	return _userController
}

// userController struct
type userController struct {
}

// Index get users
func (uc *userController) Index(ctx *web.Context) {
	ctx.WriteString("user.index")
}

// Create create user
func (uc *userController) Create(ctx *web.Context) {
	ctx.WriteString("user.create")
}

// Detail get user detail by id
func (uc *userController) Detail(ctx *web.Context) {
	ctx.WriteString("user.detail")
}

// Update update user by id
func (uc *userController) Update(ctx *web.Context) {
	ctx.WriteString("user.update")
}

// Destroy delete user by id
func (uc *userController) Destroy(ctx *web.Context) {
	ctx.WriteString("user.destroy")
}

```

### Route
```go
package route

import (
	"sync"

	"github.com/webpkg/web"
	"github.com/webpkg/api/controller"
)

var (
	_once sync.Once
)

// Init config
func Init(app *web.Application) {

	_once.Do(func() {
		app.Resource("/user/", controller.CreateUserController())
	})
}

```

### Contract
```go
package contract

import "github.com/webpkg/api/model"

// UserRepository interface
type UserRepository interface {
	Find(id uint64) (*model.User, error)
	Create(user *model.User) error
}

```

### Repository
```go
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

```