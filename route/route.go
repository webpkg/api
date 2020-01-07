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
