package helper

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"path/filepath"
)

// ReadJSON read json to data
func ReadJSON(v interface{}, filename string) error {

	if !filepath.IsAbs(filename) {
		dir, err := os.Getwd()

		if err != nil {
			return err
		}

		filename = filepath.Join(dir, filename)
	}

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

	if !filepath.IsAbs(filename) {
		dir, err := os.Getwd()

		if err != nil {
			return err
		}

		filename = filepath.Join(dir, filename)
	}

	if force || !FileExists(filename) {

		data, err := json.MarshalIndent(v, "", "  ")

		if err != nil {
			return err
		}

		return ioutil.WriteFile(filename, data, 0600)
	}

	return os.ErrExist
}

// FileExists check file, exist return true else return false
func FileExists(filename string) bool {
	info, err := os.Stat(filename)
	if err != nil {
		return false
	}

	return !info.IsDir()
}
