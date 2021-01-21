package model

import "time"

// CreateTest return *Test
func CreateTest() *Test {

	test := &Test{}

	return test
}

// Test model
// @table tests
type Test struct {
	// @column PrimaryKey
    ID uint64 `json:"id"`
	// @column $dataType=varchar(127)
	TestName string `json:"testName"`
	TestDescription *string `json:"testDescription"`
	// lt 0 deleted, 0 pendding, 1 valid
	Status    int        `json:"status"`
	DeletedAt *time.Time `json:"-"`
	CreatedAt *time.Time `json:"createdAt"`
	UpdatedAt *time.Time `json:"updatedAt"`
}

// CreateTestCollection return *TestCollection
func CreateTestCollection() *TestCollection {

	testCollection := &TestCollection{}

	return testCollection
}

// TestCollection Test list
type TestCollection []Test

// Len return len
func (o *TestCollection) Len() int { return len(*o) }

// Swap swap i, j
func (o *TestCollection) Swap(i, j int) { (*o)[i], (*o)[j] = (*o)[j], (*o)[i] }

// Less compare i, j
func (o *TestCollection) Less(i, j int) bool { return (*o)[i].TestName < (*o)[j].TestName }
