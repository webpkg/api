package route

import (
	"sync"

	"github.com/webpkg/web"
)

const (
	prefix = ""
)

var (
	_once sync.Once
)

// Init config
func Init(app *web.Application) {
	_once.Do(func() {
		testRoute(app, prefix)
	})
}
