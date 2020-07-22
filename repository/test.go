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
