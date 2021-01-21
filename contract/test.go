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
	// CreateTest ID, TestName, TestDescription, Status, CreatedAt
	// return uint64, error
	CreateTest(test *model.Test) (uint64, error)
	// UpdateTest return rowsAffected, error
	// SET TestName, TestDescription, Status, UpdatedAt
	// WHERE ID
	UpdateTest(test *model.Test) (int64, error)
	// UpdateTestStatus return rowsAffected, error
	// SET status
	// WHERE ID
	UpdateTestStatus(test *model.Test) (int64, error)
	// DestroyTest return rowsAffected, error
	// WHERE id uint64
	DestroyTest(id uint64) (int64, error)
	// DestroyTestSoft return rowsAffected, error
	// WHERE id uint64
	DestroyTestSoft(id uint64) (int64, error)
}
