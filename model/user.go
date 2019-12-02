package model

import "time"

// User model
type User struct {
	ID              uint64
	FirstName       string
	LastName        string
	Email           string
	EmailEerifiedAt time.Time
	Password        string
	RememberToken   string
	Right           int64
	DeletedAt       time.Time
	CreatedAt       time.Time
	UpdatedAt       time.Time
}
