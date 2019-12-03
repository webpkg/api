# api

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