// Copyright 2023 The GoStartKit Authors. All rights reserved.
// Use of this source code is governed by a AGPL
// license that can be found in the LICENSE file.
// https://gostartkit.com
package proxy

import (
	"github.com/gostartkit/api/model"
	"github.com/gostartkit/api/repository"
)

var (
	testRepository = repository.CreateTestRepository()
)

// CreateTestID return test.ID error
func CreateTestID() (uint64, error) {
	return testRepository.CreateTestID()
}

// GetTests return *model.TestCollection, error
func GetTests(filter string, orderBy string, page int, pageSize int) (*model.TestCollection, error) {
	return testRepository.GetTests(filter, orderBy, page, pageSize)
}

// GetTest return *model.Test, error
func GetTest(id uint64) (*model.Test, error) {
	return testRepository.GetTest(id)
}

// CreateTest return int64, error
// Attributes: ID uint64, TestName string, TestDescription *string, Status int
func CreateTest(test *model.Test) (int64, error) {
	return testRepository.CreateTest(test)
}

// UpdateTest return int64, error
// Attributes: TestName string, TestDescription *string, Status int
func UpdateTest(test *model.Test) (int64, error) {
	return testRepository.UpdateTest(test)
}

// UpdateTestPartial return int64, error
// Attributes: TestName string, TestDescription *string, Status int
func UpdateTestPartial(test *model.Test, attrsName ...string) (int64, error) {
	return testRepository.UpdateTestPartial(test, attrsName...)
}

// UpdateTestStatus return int64, error
// Attributes: Status int
func UpdateTestStatus(test *model.Test) (int64, error) {
	return testRepository.UpdateTestStatus(test)
}

// DestroyTest return int64, error
func DestroyTest(id uint64) (int64, error) {
	return testRepository.DestroyTest(id)
}

// DestroyTest return int64, error
func DestroyTestSoft(id uint64) (int64, error) {
	return testRepository.DestroyTestSoft(id)
}
