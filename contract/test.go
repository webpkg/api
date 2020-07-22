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
