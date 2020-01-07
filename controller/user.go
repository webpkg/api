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
