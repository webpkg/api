package helper

import (
	"encoding/json"
	"errors"
	"net/http"
)

var (
	// ErrUnExpectedError unexpected error
	ErrUnExpectedError = errors.New("unexpected error")
)

// HTTPGet http get
func HTTPGet(client *http.Client, uri string, accessToken string, v interface{}) error {

	req, err := http.NewRequest(http.MethodGet, uri, nil)

	if err != nil {
		return err
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+accessToken)
	req.Header.Set("Accept", "application/json")

	resp, err := client.Do(req)

	if err != nil {
		return err
	}

	defer resp.Body.Close()

	switch resp.StatusCode {
	case 200:
		if err := json.NewDecoder(resp.Body).Decode(v); err != nil {
			return err
		}
		return nil
	case 400, 401, 403:
		errMessage := ""

		if err := json.NewDecoder(resp.Body).Decode(&errMessage); err != nil {
			return err
		}
		return errors.New(errMessage)
	default:
		return ErrUnExpectedError
	}
}
