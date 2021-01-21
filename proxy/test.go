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

// CreateTest ID, TestName, TestDescription, Status, CreatedAt
// return uint64, error
func CreateTest(test *model.Test) (uint64, error) {
	repo := repository.CreateTestRepository()
	return repo.CreateTest(test)
}

// UpdateTest return rowsAffected, error
// SET TestName, TestDescription, Status, UpdatedAt
// WHERE ID
func UpdateTest(test *model.Test) (int64, error) {
	repo := repository.CreateTestRepository()
	return repo.UpdateTest(test)
}

// UpdateTestStatus return rowsAffected, error
// SET status
// WHERE ID
func UpdateTestStatus(test *model.Test) (int64, error) {
	repo := repository.CreateTestRepository()
	return repo.UpdateTestStatus(test)
}

// DestroyTest return rowsAffected, error
// WHERE id uint64
func DestroyTest(id uint64) (int64, error) {
	repo := repository.CreateTestRepository()
	return repo.DestroyTest(id)
}

// DestroyTestSoft return rowsAffected, error
// WHERE id uint64
func DestroyTestSoft(id uint64) (int64, error) {
	repo := repository.CreateTestRepository()
	return repo.DestroyTestSoft(id)
}
