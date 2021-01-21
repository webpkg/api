# api

a start point for your go project base on webpkg/web.

1. support one write multiple read databases

2. support multiple instance whith same/different database instance

3. support soft delete

### Install
```bash
git clone git@github.com:webpkg/api.git
```
replace github.com/webpkg/api to your module path
```bash
make clean  && make
```

### Route
```go
package route

import (
	"github.com/webpkg/api/controller"
	"github.com/webpkg/api/middleware"
	"github.com/webpkg/web"
)

func testRoute(app *web.Application, prefix string) {

	test := controller.CreateTestController()

	app.Get(prefix+"/test/", middleware.Chain(test.Index, "test.all"))
	app.Post(prefix+"/test/", middleware.Chain(test.Create, "test.edit"))
	app.Get(prefix+"/test/:id", middleware.Chain(test.Detail, "test.all"))
	app.Patch(prefix+"/test/:id", middleware.Chain(test.Update, "test.edit"))
	app.Put(prefix+"/test/:id", middleware.Chain(test.Update, "test.edit"))
	app.Delete(prefix+"/test/:id", middleware.Chain(test.Destroy, "test.edit"))
}

```