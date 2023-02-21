// Copyright 2023 The GoStartKit Authors. All rights reserved.
// Use of this source code is governed by a AGPL
// license that can be found in the LICENSE file.
// https://gostartkit.com
package repository

import (
	"fmt"
	"reflect"
	"strings"
	"sync"
	"sync/atomic"

	"github.com/webpkg/api/config"
	"github.com/webpkg/api/contract"
	"github.com/webpkg/api/helper"
	"github.com/webpkg/api/model"
	"github.com/webpkg/web"
)

var (
	_testRepository     contract.TestRepository
	_onceTestRepository sync.Once
)

// CreateTestRepository return contract.TestRepository
func CreateTestRepository() contract.TestRepository {

	_onceTestRepository.Do(func() {
		_testRepository = &TestRepository{}
	})

	return _testRepository
}

// TestRepository struct
type TestRepository struct {
	mu     sync.Mutex
	testID uint64
}

// CreateTestID return test.ID error
func (r *TestRepository) CreateTestID() (uint64, error) {
	r.mu.Lock()
	if r.testID == 0 {
		var err error
		r.testID, err = max("tests", "id", config.App().AppID, config.App().AppNum)
		if err != nil {
			r.mu.Unlock()
			return 0, err
		}
		if r.testID == 0 {
			r.testID = config.App().AppID - config.App().AppNum
		}
	}
	r.mu.Unlock()
	testID := atomic.AddUint64(&r.testID, config.App().AppNum)
	return testID, nil
}

// GetTests return *model.TestCollection, error
func (r *TestRepository) GetTests(filter string, orderBy string, page int, pageSize int) (*model.TestCollection, error) {

	var sqlx strings.Builder
	var args []interface{}

	sqlx.WriteString("SELECT `id`, `test_name`, `test_description`, `status`, `deleted_at`, `created_at`, `updated_at` ")
	sqlx.WriteString("FROM `tests` ")
	sqlx.WriteString("WHERE `status` >= 0 ")

	if filter != "" {
		sqlx.WriteString("AND ")
		if err := web.SqlFilter(filter, &sqlx, &args, "", r.tryParse); err != nil {
			return nil, err
		}
		sqlx.WriteString(" ")
	}

	if orderBy != "" {
		sqlx.WriteString("ORDER BY ")
		if err := web.SqlOrderBy(orderBy, &sqlx, "", r.tryParseKey); err != nil {
			return nil, err
		}
		sqlx.WriteString(" ")
	}

	sqlx.WriteString("limit ? offset ?")

	if pageSize > _maxPageSize {
		pageSize = _maxPageSize
	} else if pageSize <= 0 {
		pageSize = _pageSize
	}

	offset := 0

	if page > 1 {
		offset = (page - 1) * pageSize
	}

	args = append(args, pageSize, offset)

	rows, err := query(sqlx.String(), args...)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	tests := model.NewTestCollection()

	for rows.Next() {

		test := model.NewTest()

		err := rows.Scan(&test.ID, &test.TestName, &test.TestDescription, &test.Status, &test.DeletedAt, &test.CreatedAt, &test.UpdatedAt)

		if err != nil {
			return nil, err
		}

		*tests = append(*tests, *test)
	}

	return tests, rows.Err()
}

// GetTest return *model.Test, error
func (r *TestRepository) GetTest(id uint64) (*model.Test, error) {

	sqlx := "SELECT `id`, `test_name`, `test_description`, `status`, `deleted_at`, `created_at`, `updated_at` " +
		"FROM `tests` " +
		"WHERE `id` = ? AND `status` >= 0"

	row := queryRow(sqlx, id)

	test := model.NewTest()

	err := row.Scan(&test.ID, &test.TestName, &test.TestDescription, &test.Status, &test.DeletedAt, &test.CreatedAt, &test.UpdatedAt)

	if err != nil {
		return nil, err
	}

	return test, nil
}

// CreateTest return int64, error
// Attributes: ID uint64, TestName string, TestDescription *string, Status int
func (r *TestRepository) CreateTest(test *model.Test) (int64, error) {

	sqlx := "INSERT INTO `tests` " +
		"(`id`, `test_name`, `test_description`, `status`, `created_at`) " +
		"VALUES(?, ?, ?, ?, ?)"

	var err error

	if test.ID == 0 {

		test.ID, err = r.CreateTestID()

		if err != nil {
			return 0, err
		}
	}

	test.CreatedAt = now()

	result, err := exec(sqlx, test.ID, test.TestName, test.TestDescription, test.Status, test.CreatedAt)

	if err != nil {
		return 0, err
	}

	rowsAffected, err := result.RowsAffected()

	if err != nil {
		return 0, err
	}

	return rowsAffected, nil
}

// UpdateTest return int64, error
// Attributes: TestName string, TestDescription *string, Status int
func (r *TestRepository) UpdateTest(test *model.Test) (int64, error) {

	sqlx := "UPDATE `tests` " +
		"SET `test_name` = ?, `test_description` = ?, `status` = ?, `updated_at` = ? " +
		"WHERE `id` = ?"

	test.UpdatedAt = now()

	result, err := exec(sqlx, test.TestName, test.TestDescription, test.Status, test.UpdatedAt, test.ID)

	if err != nil {
		return 0, err
	}

	rowsAffected, err := result.RowsAffected()

	if err != nil {
		return 0, err
	}

	return rowsAffected, nil
}

// UpdateTestPartial return int64, error
// Attributes: TestName string, TestDescription *string, Status int
func (r *TestRepository) UpdateTestPartial(test *model.Test, attrsName ...string) (int64, error) {

	var sqlx strings.Builder
	var args []interface{}

	rv := reflect.Indirect(reflect.ValueOf(test))

	sqlx.WriteString("UPDATE `tests` SET ")

	for i, n := range attrsName {

		columnName, attributeName, _, err := r.tryParseKey(n)

		if err != nil {
			return 0, err
		}

		if i > 0 {
			sqlx.WriteString(", ")
		}

		fmt.Fprintf(&sqlx, "`%s` = ?", columnName)

		v := rv.FieldByName(attributeName).Interface()

		args = append(args, v)
	}

	sqlx.WriteString(", `updated_at` = ?")

	test.UpdatedAt = now()

	args = append(args, test.UpdatedAt)

	sqlx.WriteString(" WHERE `id` = ?")

	args = append(args, test.ID)

	result, err := exec(sqlx.String(), args...)

	if err != nil {
		return 0, err
	}

	rowsAffected, err := result.RowsAffected()

	if err != nil {
		return 0, err
	}

	return rowsAffected, nil
}

// UpdateTestStatus return int64, error
// Attributes: Status int
func (r *TestRepository) UpdateTestStatus(test *model.Test) (int64, error) {

	sqlx := "UPDATE `tests` " +
		"SET `status` = ?, `updated_at` = ? " +
		"WHERE `id` = ?"

	test.UpdatedAt = now()

	result, err := exec(sqlx, test.Status, test.UpdatedAt, test.ID)

	if err != nil {
		return 0, err
	}

	rowsAffected, err := result.RowsAffected()

	if err != nil {
		return 0, err
	}

	return rowsAffected, nil
}

// DestroyTest return int64, error
func (r *TestRepository) DestroyTest(id uint64) (int64, error) {

	sqlx := "DELETE FROM `tests` WHERE `id` = ?"

	result, err := exec(sqlx, id)

	if err != nil {
		return 0, err
	}

	rowsAffected, err := result.RowsAffected()

	if err != nil {
		return 0, err
	}

	return rowsAffected, nil
}

// DestroyTest return int64, error
func (r *TestRepository) DestroyTestSoft(id uint64) (int64, error) {

	sqlx := "UPDATE `tests` " +
		"SET `status` = -ABS(`status`) " +
		"WHERE `id` = ?"

	result, err := exec(sqlx, id)

	if err != nil {
		return 0, err
	}

	rowsAffected, err := result.RowsAffected()

	if err != nil {
		return 0, err
	}

	return rowsAffected, nil
}

// tryParse return columnName, attributeValue, error
func (r *TestRepository) tryParse(key string, val string) (string, interface{}, error) {

	columnName, _, attributeType, err := r.tryParseKey(key)

	if err != nil {
		return "", nil, err
	}

	v, err := helper.TryParse(val, attributeType)

	if err != nil {
		return "", nil, err
	}

	return columnName, v, nil
}

// tryParseKey return columnName, attributeName, attributeType, error
func (r *TestRepository) tryParseKey(key string) (string, string, string, error) {

	switch key {
	case "id", "ID":
		return "id", "ID", "uint64", nil
	case "testName", "TestName":
		return "test_name", "TestName", "string", nil
	case "testDescription", "TestDescription":
		return "test_description", "TestDescription", "*string", nil
	case "status", "Status":
		return "status", "Status", "int", nil
	case "deletedAt", "DeletedAt":
		return "deleted_at", "DeletedAt", "*time.Time", nil
	case "createdAt", "CreatedAt":
		return "created_at", "CreatedAt", "*time.Time", nil
	case "updatedAt", "UpdatedAt":
		return "updated_at", "UpdatedAt", "*time.Time", nil
	default:
		return "", "", "", fmt.Errorf("'test.%s' not exists", key)
	}
}
