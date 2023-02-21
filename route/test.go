// Copyright 2023 The GoStartKit Authors. All rights reserved.
// Use of this source code is governed by a AGPL
// license that can be found in the LICENSE file.
// https://gostartkit.com
package route

import (
	"github.com/webpkg/api/controller"
	"github.com/webpkg/api/middleware"
	"github.com/webpkg/web"
)

func testRoute(app *web.Application, prefix string) {

	test := controller.CreateTestController()

	app.Get(prefix+"/test/", middleware.Chain(test.Index, "test.all"))
	app.Get(prefix+"/test/:id", middleware.Chain(test.Detail, "test.all"))
	app.Post(prefix+"/apply/test/id/", middleware.Chain(test.CreateID, "test.edit"))
	app.Post(prefix+"/test/", middleware.Chain(test.Create, "test.edit"))
	app.Put(prefix+"/test/:id", middleware.Chain(test.Update, "test.edit"))
	app.Patch(prefix+"/test/:id", middleware.Chain(test.UpdatePartial, "test.edit"))
	app.Patch(prefix+"/test/:id/status/", middleware.Chain(test.UpdateStatus, "test.edit"))
	app.Delete(prefix+"/test/:id", middleware.Chain(test.Destroy, "test.edit"))
}
