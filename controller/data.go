// Copyright 2023 The GoStartKit Authors. All rights reserved.
// Use of this source code is governed by a AGPL
// license that can be found in the LICENSE file.
// https://gostartkit.com
package controller

import (
	"sync"

	"github.com/webpkg/api/config"
	"github.com/webpkg/web"
)

var (
	_dataController     *DataController
	_onceDataController sync.Once
)

// CreateDataController return *DataController
func CreateDataController() *DataController {

	_onceDataController.Do(func() {
		_dataController = &DataController{}
	})

	return _dataController
}

// DataController struct
type DataController struct {
}

// Index get data
func (c *DataController) Index(ctx *web.Context) (web.Data, error) {
	return nil, nil
}

// Rbac get rbac
func (c *DataController) Rbac(ctx *web.Context) (web.Data, error) {
	return config.Rbac(), nil
}

// RbacUserRight get current user right
func (c *DataController) RbacUserRight(ctx *web.Context) (web.Data, error) {
	return config.Rbac().Keys(ctx.UserRight()), nil
}
