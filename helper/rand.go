// Copyright 2023 The GoStartKit Authors. All rights reserved.
// Use of this source code is governed by a AGPL
// license that can be found in the LICENSE file.
// https://gostartkit.com
package helper

import (
	"math/rand"
	"time"
)

const _charsetRand = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

var _seededRand = rand.New(rand.NewSource(time.Now().Unix()))

// RandStringWithCharset rand string with charset
func RandStringWithCharset(length int, charset string) string {
	b := make([]byte, length)
	l := len(charset)
	for i := range b {
		b[i] = charset[_seededRand.Intn(l)]
	}
	return string(b)
}

// RandString rand string
func RandString(length int) string {
	return RandStringWithCharset(length, _charsetRand)
}

// RandInt rand int between [min, max)
func RandInt(min int, max int) int {
	if min <= 0 || max <= 0 {
		return 0
	}

	if min >= max {
		return max
	}

	return _seededRand.Intn(max-min) + min
}

// RandMax rand int between [0, max)
func RandMax(max int) int {
	if max <= 1 {
		return 0
	}

	return _seededRand.Intn(max)
}
