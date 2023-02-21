// Copyright 2023 The GoStartKit Authors. All rights reserved.
// Use of this source code is governed by a AGPL
// license that can be found in the LICENSE file.
// https://gostartkit.com
package helper

import (
	"errors"
)

var (
	ErrModelInvalid  = errors.New("model invalid")
	ErrModelExist    = errors.New("model already exists")
	ErrModelNotExist = errors.New("model does not exist")
)
