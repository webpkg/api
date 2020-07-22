package route

import (
	"sync"

	"github.com/webpkg/web"
)

var (
	_once sync.Once
)

// Init config
func Init(app *web.Application) {
	_once.Do(func() {

	})
}
