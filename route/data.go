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

func dataRoute(app *web.Application, prefix string) {

	data := controller.CreateDataController()

	app.Get(prefix+"/config/rbac/", middleware.Chain(data.Rbac))
	app.Get(prefix+"/config/rbac/user/right/", middleware.Chain(data.RbacUserRight))
}
