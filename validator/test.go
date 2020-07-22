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
