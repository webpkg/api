// Copyright 2023 The GoStartKit Authors. All rights reserved.
// Use of this source code is governed by a AGPL
// license that can be found in the LICENSE file.
// https://gostartkit.com
package helper

import (
	"crypto/sha256"
	"encoding/hex"
)

// CreateToken generate token
func CreateToken() string {
	return RandString(32)
}

// Hash val
func Hash(val string) string {
	h := sha256.New()
	h.Write([]byte(val))
	return hex.EncodeToString(h.Sum(nil))
}
