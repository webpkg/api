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
./development.bash
```
### Model
```go
package model

import "time"

// CreateTest return *Test
func CreateTest() *Test {

	test := &Test{}

	return test
}

// Test model
// @table tests
type Test struct {
	// @column PrimaryKey
	ID uint64 `json:"id"`
	// @column $dataType=varchar(127)
	TestName        string     `json:"testName"`
	TestDescription *string    `json:"testDescription"`
	CreatedAt       *time.Time `json:"-"`
	UpdatedAt       *time.Time `json:"-"`
}

// CreateTestCollection return *TestCollection
func CreateTestCollection() *TestCollection {

	testCollection := &TestCollection{}

	return testCollection
}

// TestCollection Test list
type TestCollection []Test

// Len return len
func (o *TestCollection) Len() int { return len(*o) }

// Swap swap i, j
func (o *TestCollection) Swap(i, j int) { (*o)[i], (*o)[j] = (*o)[j], (*o)[i] }

// Less compare i, j
func (o *TestCollection) Less(i, j int) bool { return (*o)[i].TestName < (*o)[j].TestName }

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

### Controller
```go
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

```

### Validator
```go
package validator

import (
	"github.com/webpkg/api/model"
)

// CreateTest validate create test
func CreateTest(test *model.Test) error {

	if test.ID != 0 {
		return CreateValidationError("id", "invalid")
	}

	if test.TestName == "" {
		return CreateValidationError("testName", "required")
	}

	return nil
}

// UpdateTest validate update test
func UpdateTest(test *model.Test) error {

	if test.TestName == "" {
		return CreateValidationError("testName", "required")
	}

	return nil
}

```

### Proxy
```go
package proxy

import (
	"github.com/webpkg/api/model"
	"github.com/webpkg/api/repository"
)

// CreateTestID generate a new testID
func CreateTestID() uint64 {
	repo := repository.CreateTestRepository()
	return repo.CreateTestID()
}

// GetTestsByKey get tests by key
func GetTestsByKey(key string, page int, pageSize int) (*model.TestCollection, error) {
	repo := repository.CreateTestRepository()
	return repo.GetTestsByKey(key, page, pageSize)
}

// GetTest by id uint64
func GetTest(id uint64) (*model.Test, error) {
	repo := repository.CreateTestRepository()
	return repo.GetTest(id)
}

// CreateTest ID, TestName, TestDescription, CreatedAt
// return uint64, error
func CreateTest(test *model.Test) (uint64, error) {
	repo := repository.CreateTestRepository()
	return repo.CreateTest(test)
}

// UpdateTest return rowsAffected, error
// SET TestName, TestDescription, UpdatedAt
// WHERE ID
func UpdateTest(test *model.Test) (int64, error) {
	repo := repository.CreateTestRepository()
	return repo.UpdateTest(test)
}

// DestroyTest return rowsAffected, error
// WHERE id uint64
func DestroyTest(id uint64) (int64, error) {
	repo := repository.CreateTestRepository()
	return repo.DestroyTest(id)
}

```

### Contract
```go
package contract

import "github.com/webpkg/api/model"

// TestRepository interface
type TestRepository interface {
	// CreateTestID generate a new testID
	CreateTestID() uint64
	// GetTestsByKey get tests by key
	GetTestsByKey(key string, page int, pageSize int) (*model.TestCollection, error)
	// GetTest by id uint64
	GetTest(id uint64) (*model.Test, error)
	// CreateTest ID, TestName, TestDescription, CreatedAt
	// return uint64, error
	CreateTest(test *model.Test) (uint64, error)
	// UpdateTest return rowsAffected, error
	// SET TestName, TestDescription, UpdatedAt
	// WHERE ID
	UpdateTest(test *model.Test) (int64, error)
	// DestroyTest return rowsAffected, error
	// WHERE id uint64
	DestroyTest(id uint64) (int64, error)
}

```

### Repository
```go
package repository

import (
	"sync"
	"sync/atomic"

	"github.com/webpkg/api/contract"
	"github.com/webpkg/api/model"
)

var (
	_testRepository     contract.TestRepository
	_onceTestRepository sync.Once
)

// CreateTestRepository return contract.TestRepository
func CreateTestRepository() contract.TestRepository {

	_onceTestRepository.Do(func() {
		_testRepository = &TestRepository{}

		if IDConfig().TestID == 0 {
			IDConfig().TestID = AppConfig().AppID - AppConfig().AppNum
		}
	})

	return _testRepository
}

// TestRepository struct
type TestRepository struct {
}

// CreateTestID generate a new testID
func (r *TestRepository) CreateTestID() uint64 {
	return atomic.AddUint64(&IDConfig().TestID, AppConfig().AppNum)
}

// GetTestsByKey get tests by key
func (r *TestRepository) GetTestsByKey(key string, page int, pageSize int) (*model.TestCollection, error) {
	sql := "SELECT `id`, `test_name`, `test_description`, `created_at`, `updated_at` " +
		"FROM `tests` " +
		"WHERE `test_name` like ? " +
		"limit ? offset ? "

	key = "%" + key + "%"

	if pageSize > _maxPageSize {
		pageSize = _maxPageSize
	} else if pageSize <= 0 {
		pageSize = _pageSize
	}

	offset := 0

	if page > 1 {
		offset = (page - 1) * pageSize
	}

	rows, err := query(sql, key, pageSize, offset)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	tests := model.CreateTestCollection()

	for rows.Next() {

		test := model.CreateTest()

		err := rows.Scan(&test.ID, &test.TestName, &test.TestDescription, &test.CreatedAt, &test.UpdatedAt)

		if err != nil {
			return nil, err
		}

		*tests = append(*tests, *test)
	}

	return tests, rows.Err()
}

// GetTest by id uint64
func (r *TestRepository) GetTest(id uint64) (*model.Test, error) {
	sql := "SELECT `id`, `test_name`, `test_description`, `created_at`, `updated_at` " +
		"FROM `tests` " +
		"WHERE `id` = ? " +
		"limit 1 "

	row := queryRow(sql, id)

	test := model.CreateTest()

	err := row.Scan(&test.ID, &test.TestName, &test.TestDescription, &test.CreatedAt, &test.UpdatedAt)

	if err != nil {
		return nil, err
	}

	return test, nil
}

// CreateTest ID, TestName, TestDescription, CreatedAt
// return uint64, error
func (r *TestRepository) CreateTest(test *model.Test) (uint64, error) {
	sql := "INSERT INTO `tests` " +
		"(`id`, `test_name`, `test_description`, `created_at`) " +
		"VALUES(?, ?, ?, ?) "

	if test.ID == 0 {
		test.ID = r.CreateTestID()
	}

	_, err := exec(sql, test.ID, test.TestName, test.TestDescription, now())

	if err != nil {
		return 0, err
	}

	return test.ID, nil
}

// UpdateTest return rowsAffected, error
// SET TestName, TestDescription, UpdatedAt
// WHERE ID
func (r *TestRepository) UpdateTest(test *model.Test) (int64, error) {
	sql := "UPDATE `tests` " +
		"SET `test_name` = ?, `test_description` = ?, `updated_at` = ? " +
		"WHERE `id` = ? "

	result, err := exec(sql, test.TestName, test.TestDescription, now(), test.ID)

	if err != nil {
		return 0, err
	}

	rowsAffected, err := result.RowsAffected()

	if err != nil {
		return 0, err
	}

	return rowsAffected, nil
}

// DestroyTest return rowsAffected, error
// WHERE id uint64
func (r *TestRepository) DestroyTest(id uint64) (int64, error) {
	sql := "DELETE FROM `tests` WHERE `id` = ? "

	result, err := exec(sql, id)

	if err != nil {
		return 0, err
	}

	rowsAffected, err := result.RowsAffected()

	if err != nil {
		return 0, err
	}

	return rowsAffected, nil
}

```

--------------------------------------------------
# Api Doc

## Get Tests

### URL

  GET /test/?key=&page=1&pagesize=20

### Header
```bash
Content-Type: application/json
Authorization: Bearer $ACCESSTOKEN
```

### Success Response:
#### Status Code: 200
#### Data (TestCollection):
```json
[
    {
        "id": "uint64",
        "testName": "string",
        "testDescription": "string",
        "createdAt": "time.Time",
        "updatedAt": "time.Time"
    }
]
```

### Error Response:
#### Status Code: 400
#### Data:
```json
"error message"
```

### Error Response:
#### Status Code: 401
#### Data:
```json
"invalid token"
```

### Error Response:
#### Status Code: 403
#### Data:
```json
"permission denied"
```

--------------------

### Query:

| QueryName | Required | Type |
|-----------|---------:|-----:|
|  key      |    No    |string|
|  page     |    No    |  int |
|  pagesize |    No    |  int |

### Data (Test):

| AttributeName | Required | Type |
|---------------|---------:|-----:|
|id|Yes|uint64|
|testName|Yes|string|
|testDescription|No|string|
|createdAt|No|time.Time|
|updatedAt|No|time.Time|

--------------------------
## Get Test

### URL

  GET /test/:id

### Header
```bash
Content-Type: application/json
Authorization: Bearer $ACCESSTOKEN
```

### Success Response:
#### Status Code: 200
#### Data (Test):
```json
{
    "id": "uint64",
    "testName": "string",
    "testDescription": "string",
    "createdAt": "time.Time",
    "updatedAt": "time.Time"
}
```

### Error Response:
#### Status Code: 400
#### Data:
```json
"error message"
```

### Error Response:
#### Status Code: 401
#### Data:
```json
"invalid token"
```

### Error Response:
#### Status Code: 403
#### Data:
```json
"permission denied"
```

--------------------

### Data (Test):

| AttributeName | Required | Type |
|---------------|---------:|-----:|
|id|Yes|uint64|
|testName|Yes|string|
|testDescription|No|string|
|createdAt|No|time.Time|
|updatedAt|No|time.Time|

--------------------------------

## Create Test

### URL

  POST /test/

### Header
```bash
Content-Type: application/json
Authorization: Bearer $ACCESSTOKEN
```

### Payload (Test)
```json
{
    "id": "uint64",
    "testName": "string",
    "testDescription": "string",
    "createdAt": "time.Time"
}
```

### Success Response:
#### Status Code: 200
#### Data (id uint64):
```json

```

### Error Response:
#### Status Code: 400
#### Data:
```json
"error message"
```

### Error Response:
#### Status Code: 401
#### Data:
```json
"invalid token"
```

### Error Response:
#### Status Code: 403
#### Data:
```json
"permission denied"
```

--------------------

### Data (Test):

| AttributeName | Required | Type | Validator |
|---------------|---------:|-----:|----------:|
|id|Yes|uint64||
|testName|Yes|string||
|testDescription|No|string||
|createdAt|No|time.Time||

-----------------------------

## Updata Test

### URL

  PUT /test/:id

### Header
```bash
Content-Type: application/json
Authorization: Bearer $ACCESSTOKEN
```

### Payload (Test)
```json
{
    "id": "uint64",
    "testName": "string",
    "testDescription": "string",
    "updatedAt": "time.Time"
}
```

### Success Response:
#### Status Code: 200
#### Data (rowsAffected int64):
```json

```

### Error Response:
#### Status Code: 400
#### Data:
```json
"error message"
```

### Error Response:
#### Status Code: 401
#### Data:
```json
"invalid token"
```

### Error Response:
#### Status Code: 403
#### Data:
```json
"permission denied"
```

--------------------

### Data (Test):

| AttributeName | Required | Type | Validator |
|---------------|---------:|-----:|----------:|
|testName|Yes|string||
|testDescription|No|string||
|updatedAt|No|time.Time||

### Where:

| AttributeName | Required | Type |
|---------------|---------:|-----:|
|id|Yes|uint64|

--------------------------------------

## Destroy Test

### URL

  DELETE /test/:id

### Header
```bash
Content-Type: application/json
Authorization: Bearer $ACCESSTOKEN
```

### Success Response:
#### Status Code: 200
#### Data (rowsAffected int64):
```json

```

### Error Response:
#### Status Code: 400
#### Data:
```json
"error message"
```

### Error Response:
#### Status Code: 401
#### Data:
```json
"invalid token"
```

### Error Response:
#### Status Code: 403
#### Data:
```json
"permission denied"
```

--------------------

### Where:

| AttributeName | Required | Type |
|---------------|---------:|-----:|
|id|Yes|uint64|