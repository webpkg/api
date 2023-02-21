// Copyright 2023 The GoStartKit Authors. All rights reserved.
// Use of this source code is governed by a AGPL
// license that can be found in the LICENSE file.
// https://gostartkit.com
package helper

import (
	"encoding/json"
	"os"
	"path/filepath"
)

// ReadJSON read json to data
func ReadJSON(filename string, v interface{}) error {

	f, err := os.Open(filename)

	if err != nil {
		return err
	}

	defer f.Close()

	return json.NewDecoder(f).Decode(v)
}

// WriteJSON write data to json
func WriteJSON(filename string, v interface{}, overwrite bool) error {

	if !overwrite && FileExist(filename) {
		return os.ErrExist
	}

	dir := filepath.Dir(filename)

	if !DirExist(dir) {
		if err := os.MkdirAll(dir, 0700); err != nil {
			return err
		}
	}

	f, err := os.OpenFile(filename, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644)

	if err != nil {
		return err
	}

	defer f.Close()

	enc := json.NewEncoder(f)

	enc.SetIndent("", "  ")

	return enc.Encode(v)
}
