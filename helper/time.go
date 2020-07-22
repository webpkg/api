package helper

import "time"

// Now return *time.Time
func Now() *time.Time {
	now := time.Now()
	return &now
}
