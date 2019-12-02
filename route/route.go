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
