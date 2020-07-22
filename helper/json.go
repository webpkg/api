package helper

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"path/filepath"
)

// ReadJSON read json to data
func ReadJSON(v interface{}, filename string) error {

	data, err := ioutil.ReadFile(filename)

	if err != nil {
		return err
	}

	if err := json.Unmarshal(data, v); err != nil {
		return err
	}

	return nil
}

// WriteJSON write data to json
func WriteJSON(v interface{}, filename string, force bool) error {

	if force || !FileExist(filename) {

		data, err := json.MarshalIndent(v, "", "  ")

		if err != nil {
			return err
		}

		dir := filepath.Dir(filename)

		if !DirExist(dir) {
			err = os.MkdirAll(dir, 0700)

			if err != nil {
				return err
			}
		}

		return ioutil.WriteFile(filename, data, 0600)
	}

	return os.ErrExist
}
