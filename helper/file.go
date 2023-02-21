// Copyright 2023 The GoStartKit Authors. All rights reserved.
// Use of this source code is governed by a AGPL
// license that can be found in the LICENSE file.
// https://gostartkit.com
package helper

import (
	"os"
	"path/filepath"
)

// FileExist check file
func FileExist(filename string) bool {
	info, err := os.Stat(filename)

	if err != nil && os.IsNotExist(err) {
		return false
	}

	return !info.IsDir()
}

// DirExist check dir
func DirExist(dir string) bool {
	info, err := os.Stat(dir)

	if err != nil && os.IsNotExist(err) {
		return false
	}

	return info.IsDir()
}

// HomeDir join path with app home dir
func HomeDir(path string) (string, error) {
	if !filepath.IsAbs(path) {
		dir, err := os.Getwd()

		if err != nil {
			return "", err
		}

		path = filepath.Join(dir, path)
	}

	return path, nil
}
