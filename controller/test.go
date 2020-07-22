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
func (c *TestController) Index(ctx *web.Context) {
	var (
		page     int
		pageSize int
	)

	key := ctx.Query("key")
	ctx.TryParseQuery("page", &page)
	ctx.TryParseQuery("pagesize", &pageSize)

	ctx.AbortIf(proxy.GetTestsByKey(key, page, pageSize))
}

// Create create test
func (c *TestController) Create(ctx *web.Context) {
	test := model.CreateTest()
	ctx.ParseBody(test)
	ctx.Abort(validator.CreateTest(test))

	ctx.AbortIf(proxy.CreateTest(test))
}

// Detail get test detail by id
func (c *TestController) Detail(ctx *web.Context) {
	var (
		id uint64
	)
	ctx.ParseParam("id", &id)

	ctx.AbortIf(proxy.GetTest(id))
}

// Update update test by id
func (c *TestController) Update(ctx *web.Context) {
	test := model.CreateTest()
	ctx.ParseBody(test)
	ctx.Abort(validator.UpdateTest(test))

	ctx.AbortIf(proxy.UpdateTest(test))
}

// Destroy delete test by id
func (c *TestController) Destroy(ctx *web.Context) {
	var (
		id uint64
	)
	ctx.ParseParam("id", &id)

	ctx.AbortIf(proxy.DestroyTest(id))
}
