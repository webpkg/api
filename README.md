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
	"log"
	"sync"

	"github.com/webpkg/web"
)

var (
	_userController     *UserController
	_onceUserController sync.Once
)

// CreateUserController return *UserController
func CreateUserController() *UserController {

	_onceUserController.Do(func() {
		_userController = &UserController{}
	})

	return _userController
}

// UserController struct
type UserController struct {
}

// Index get users
func (uc *UserController) Index(ctx *web.Context) {
	ctx.WriteString("user.index")
}

// Create create user
func (uc *UserController) Create(ctx *web.Context) {

	name := ctx.Form("name")
	log.Printf("%s", name)
	ctx.WriteString(name)
}

// Detail get user detail by id
func (uc *UserController) Detail(ctx *web.Context) {
	ctx.WriteString("user.detail")
}

// Update update user by id
func (uc *UserController) Update(ctx *web.Context) {
	ctx.WriteString("user.update")
}

// Destroy delete user by id
func (uc *UserController) Destroy(ctx *web.Context) {
	ctx.WriteString("user.destroy")
}

```

### Route
```go
package route

import (
	"sync"

	"github.com/webpkg/api/controller"
	"github.com/webpkg/web"
)

var (
	_once sync.Once
)

// Init config
func Init(app *web.Application) {

	_once.Do(func() {
		user := controller.CreateUserController()
		app.Get("/user/", user.Index)
		app.Post("/user/", user.Create)
		app.Get("/user/:id", user.Detail)
		app.Patch("/user/:id", user.Update)
		app.Put("/user/:id", user.Update)
		app.Delete("/user/:id", user.Destroy)
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