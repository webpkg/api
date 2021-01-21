package repository

import (
	"sync"
	"sync/atomic"

	"github.com/webpkg/api/contract"
	"github.com/webpkg/api/model"
)

var (
	_testID             uint64
	_testRepository     contract.TestRepository
	_onceTestRepository sync.Once
)

// CreateTestRepository return contract.TestRepository
func CreateTestRepository() contract.TestRepository {

	_onceTestRepository.Do(func() {
		_testRepository = &TestRepository{}

		if _testID == 0 {
			_testID, _ = max("tests", "id")

			if _testID == 0 {
				_testID = WebConfig().App.AppID - WebConfig().App.AppNum
			}
		}
	})

	return _testRepository
}

// TestRepository struct
type TestRepository struct {
}

// CreateTestID generate a new testID
func (r *TestRepository) CreateTestID() uint64 {
	return atomic.AddUint64(&_testID, WebConfig().App.AppNum)
}

// GetTestsByKey get tests by key
func (r *TestRepository) GetTestsByKey(key string, page int, pageSize int) (*model.TestCollection, error) {
	sqlx := "SELECT `id`, `test_name`, `test_description`, `status`, `created_at`, `updated_at` " +
		"FROM `tests` " +
		"WHERE `test_name` like ? and `status` > 0 " +
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

	rows, err := query(sqlx, key, pageSize, offset)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	tests := model.CreateTestCollection()

	for rows.Next() {

		test := model.CreateTest()

		err := rows.Scan(&test.ID, &test.TestName, &test.TestDescription, &test.Status, &test.CreatedAt, &test.UpdatedAt)

		if err != nil {
			return nil, err
		}

		*tests = append(*tests, *test)
	}

	return tests, rows.Err()
}

// GetTest by id uint64
func (r *TestRepository) GetTest(id uint64) (*model.Test, error) {
	sqlx := "SELECT `id`, `test_name`, `test_description`, `status`, `created_at`, `updated_at` " +
		"FROM `tests` " +
		"WHERE `id` = ? and `status` > 0 " +
		"limit 1 "

	row := queryRow(sqlx, id)

	test := model.CreateTest()

	err := row.Scan(&test.ID, &test.TestName, &test.TestDescription, &test.Status, &test.CreatedAt, &test.UpdatedAt)

	if err != nil {
		return nil, err
	}

	return test, nil
}

// CreateTest ID, TestName, TestDescription, Status, CreatedAt
// return uint64, error
func (r *TestRepository) CreateTest(test *model.Test) (uint64, error) {
	sqlx := "INSERT INTO `tests` " +
		"(`id`, `test_name`, `test_description`, `status`, `created_at`) " +
		"VALUES(?, ?, ?, ?, ?) "

	if test.ID == 0 {
		test.ID = r.CreateTestID()
	}

	_, err := exec(sqlx, test.ID, test.TestName, test.TestDescription, test.Status, now())

	if err != nil {
		return 0, err
	}

	return test.ID, nil
}

// UpdateTest return rowsAffected, error
// SET TestName, TestDescription, Status, UpdatedAt
// WHERE ID
func (r *TestRepository) UpdateTest(test *model.Test) (int64, error) {
	sqlx := "UPDATE `tests` " +
		"SET `test_name` = ?, `test_description` = ?, `status` = ?, `updated_at` = ? " +
		"WHERE `id` = ? "

	result, err := exec(sqlx, test.TestName, test.TestDescription, test.Status, now(), test.ID)

	if err != nil {
		return 0, err
	}

	rowsAffected, err := result.RowsAffected()

	if err != nil {
		return 0, err
	}

	return rowsAffected, nil
}

// UpdateTestStatus return rowsAffected, error
// SET status
// WHERE ID
func (r *TestRepository) UpdateTestStatus(test *model.Test) (int64, error) {
	sqlx := "UPDATE `tests` " +
		"SET `status` = ?, `updated_at` = ? " +
		"WHERE `id` = ? "

	result, err := exec(sqlx, test.Status, now(), test.ID)

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
	sqlx := "DELETE FROM `tests` WHERE `id` = ? "

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

// DestroyTestSoft return rowsAffected, error
// WHERE id uint64
func (r *TestRepository) DestroyTestSoft(id uint64) (int64, error) {
	sqlx := "UPDATE `tests` SET `deleted_at` = ?, status=-ABS(status) " +
		"WHERE `id` = ? "

	result, err := exec(sqlx, now(), id)

	if err != nil {
		return 0, err
	}

	rowsAffected, err := result.RowsAffected()

	if err != nil {
		return 0, err
	}

	return rowsAffected, nil
}
