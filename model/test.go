
package model

import "time"

// NewTest return *Test
func NewTest() *Test {

	test := &Test{}

	return test
}

// Test model
// @Entity tableName="tests"
type Test struct {
	// @PrimaryKey
    ID uint64 `json:"id"`
	// @DataType varchar(127)
	TestName string `json:"testName"`
	// @Column dataType=varchar(255)
	TestDescription *string `json:"testDescription"`
	// @Comment "lt 0 deleted, 0 pendding, 1 valid"
	Status    int        `json:"status"`
	DeletedAt *time.Time `json:"deletedAt"`
	CreatedAt *time.Time `json:"createdAt"`
	UpdatedAt *time.Time `json:"updatedAt"`
}

// NewTestCollection return *TestCollection
func NewTestCollection() *TestCollection {

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
