package controller

import (
	"sync"

	"github.com/webpkg/api/model"
	"github.com/webpkg/api/proxy"
	"github.com/webpkg/api/validator"
	"github.com/webpkg/web"
)

var (
	_testController     *TestController
	_onceTestController sync.Once
)

// CreateTestController return TestController
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

	key := ctx.Query("key")
	ctx.TryParseQuery("page", &page)
	ctx.TryParseQuery("pagesize", &pageSize)

	return proxy.GetTestsByKey(key, page, pageSize)
}

// Create create test
func (c *TestController) Create(ctx *web.Context) (web.Data, error) {
	test := model.CreateTest()

	if err := ctx.TryParseBody(test); err != nil {
		return nil, err
	}

	if err := validator.CreateTest(test); err != nil {
		return nil, err
	}

	return proxy.CreateTest(test)
}

// Detail get test detail by id
func (c *TestController) Detail(ctx *web.Context) (web.Data, error) {
	var (
		id uint64
	)
	
	if err := ctx.TryParseParam("id", &id); err != nil {
		return nil, err
	}

	return proxy.GetTest(id)
}

// Update update test by id
func (c *TestController) Update(ctx *web.Context) (web.Data, error) {
	test := model.CreateTest()

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

// UpdateStatus update test status by id
func (c *TestController) UpdateStatus(ctx *web.Context) (web.Data, error) {
	test := model.CreateTest()

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

// Destroy delete test by id
func (c *TestController) Destroy(ctx *web.Context) (web.Data, error) {
	var (
		id uint64
	)
	if err := ctx.TryParseParam("id", &id); err != nil {
		return nil, err
	}

	return proxy.DestroyTestSoft(id)
}
