// Copyright 2023 The GoStartKit Authors. All rights reserved.
// Use of this source code is governed by a AGPL
// license that can be found in the LICENSE file.
// https://gostartkit.com
package validator

import (
	"github.com/webpkg/api/model"
)

// CreateTest validate create test
func CreateTest(test *model.Test) error {

	if test.TestName == "" {
		return CreateRequiredError("testName")
	}

	return nil
}

// UpdateTest validate update test
func UpdateTest(test *model.Test) error {

	if test.ID == 0 {
		return CreateRequiredError("id")
	}

	return nil
}

// UpdateTestPartial validate update test part
func UpdateTestPartial(test *model.Test, attrsName ...string) error {

	if test.ID == 0 {
		return CreateRequiredError("id")
	}

	if len(attrsName) == 0 {
		return CreateRequiredError("attrs")
	}

	return nil
}

// UpdateTestStatus validate update test status
func UpdateTestStatus(test *model.Test) error {

	if test.ID == 0 {
		return CreateRequiredError("id")
	}

	return nil
}
