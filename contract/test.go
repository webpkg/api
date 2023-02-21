// Copyright 2023 The GoStartKit Authors. All rights reserved.
// Use of this source code is governed by a AGPL
// license that can be found in the LICENSE file.
// https://gostartkit.com
package contract

import "github.com/webpkg/api/model"

// TestRepository interface
type TestRepository interface {
	// CreateTestID return test.ID error
	CreateTestID() (uint64, error)
	// GetTests return *model.TestCollection, error
	GetTests(filter string, orderBy string, page int, pageSize int) (*model.TestCollection, error)
	// GetTest return *model.Test, error
	GetTest(id uint64) (*model.Test, error)
	// CreateTest return int64, error
	// Attributes: ID uint64, TestName string, TestDescription *string, Status int
	CreateTest(test *model.Test) (int64, error)
	// UpdateTest return int64, error
	// Attributes: TestName string, TestDescription *string, Status int
	UpdateTest(test *model.Test) (int64, error)
	// UpdateTestPartial return int64, error
	// Attributes: TestName string, TestDescription *string, Status int
	UpdateTestPartial(test *model.Test, attrsName ...string) (int64, error)
	// UpdateTestStatus return int64, error
	// Attributes: Status int
	UpdateTestStatus(test *model.Test) (int64, error)
	// DestroyTest return int64, error
	DestroyTest(id uint64) (int64, error)
	// DestroyTest return int64, error
	DestroyTestSoft(id uint64) (int64, error)
}
