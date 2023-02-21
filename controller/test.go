// Copyright 2023 The GoStartKit Authors. All rights reserved.
// Use of this source code is governed by a AGPL
// license that can be found in the LICENSE file.
// https://gostartkit.com
package controller

import (
	"strings"
	"sync"

	"github.com/gostartkit/api/model"
	"github.com/gostartkit/api/proxy"
	"github.com/gostartkit/api/validator"
	"github.com/webpkg/web"
)

var (
	_testController     *TestController
	_onceTestController sync.Once
)

// CreateTestController return *TestController
func CreateTestController() *TestController {

	_onceTestController.Do(func() {
		_testController = &TestController{}
	})

	return _testController
}

// TestController struct
type TestController struct {
}

// Index get tests
func (c *TestController) Index(ctx *web.Context) (web.Data, error) {

	var (
		page     int
		pageSize int
	)

	filter := ctx.Query(web.QueryFilter)
	orderBy := ctx.Query(web.QueryOrderBy)
	ctx.TryParseQuery(web.QueryPage, &page)
	ctx.TryParseQuery(web.QueryPageSize, &pageSize)

	return proxy.GetTests(filter, orderBy, page, pageSize)
}

// Detail get test
func (c *TestController) Detail(ctx *web.Context) (web.Data, error) {

	var id uint64

	if err := ctx.TryParseParam("id", &id); err != nil {
		return nil, err
	}

	if id == 0 {
		return nil, validator.CreateInvalidError("id")
	}

	return proxy.GetTest(id)
}

// CreateID create test.ID
func (c *TestController) CreateID(ctx *web.Context) (web.Data, error) {
	return proxy.CreateTestID()
}

// Create create test
func (c *TestController) Create(ctx *web.Context) (web.Data, error) {

	test := model.NewTest()

	if err := ctx.TryParseBody(test); err != nil {
		return nil, err
	}

	if err := validator.CreateTest(test); err != nil {
		return nil, err
	}

	if _, err := proxy.CreateTest(test); err != nil {
		return nil, err
	}

	return test.ID, nil
}

// Update update test
func (c *TestController) Update(ctx *web.Context) (web.Data, error) {

	test := model.NewTest()

	if err := ctx.TryParseBody(test); err != nil {
		return nil, err
	}

	if err := ctx.TryParseParam("id", &test.ID); err != nil {
		return nil, err
	}

	if err := validator.UpdateTest(test); err != nil {
		return nil, err
	}

	return proxy.UpdateTest(test)
}

// UpdatePartial update test
func (c *TestController) UpdatePartial(ctx *web.Context) (web.Data, error) {

	attrs := strings.Split(ctx.Get(web.HeaderAttrs), ",")

	if len(attrs) == 0 {
		return nil, validator.CreateRequiredError(web.HeaderAttrs)
	}

	test := model.NewTest()

	if err := ctx.TryParseBody(test); err != nil {
		return nil, err
	}

	if err := ctx.TryParseParam("id", &test.ID); err != nil {
		return nil, err
	}

	if err := validator.UpdateTestPartial(test, attrs...); err != nil {
		return nil, err
	}

	return proxy.UpdateTestPartial(test, attrs...)
}

// UpdateStatus update test.Status
func (c *TestController) UpdateStatus(ctx *web.Context) (web.Data, error) {

	test := model.NewTest()

	if err := ctx.TryParseBody(test); err != nil {
		return nil, err
	}

	if err := ctx.TryParseParam("id", &test.ID); err != nil {
		return nil, err
	}

	if err := validator.UpdateTestStatus(test); err != nil {
		return nil, err
	}

	return proxy.UpdateTestStatus(test)
}

// Destroy delete test
func (c *TestController) Destroy(ctx *web.Context) (web.Data, error) {

	var id uint64

	if err := ctx.TryParseParam("id", &id); err != nil {
		return nil, err
	}

	if id == 0 {
		return nil, validator.CreateInvalidError("id")
	}

	return proxy.DestroyTestSoft(id)
}
