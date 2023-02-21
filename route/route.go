// Copyright 2023 The GoStartKit Authors. All rights reserved.
// Use of this source code is governed by a AGPL
// license that can be found in the LICENSE file.
// https://gostartkit.com
package route

import (
	"sync"

	"github.com/webpkg/web"
)

var (
	_once   sync.Once
	_prefix string
)

// Init route init
func Init(app *web.Application) {
	_once.Do(func() {
		dataRoute(app, _prefix)
		testRoute(app, _prefix)
	})
}
