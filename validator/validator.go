// Copyright 2023 The GoStartKit Authors. All rights reserved.
// Use of this source code is governed by a AGPL
// license that can be found in the LICENSE file.
// https://gostartkit.com
package validator

import (
	"errors"
	"fmt"
	"regexp"
)

const (
	strAlphabet = "^([a-zA-Z]+)$"
)

var (
	// ErrEmailOrPasswordInvalid email or password invalid
	ErrEmailOrPasswordInvalid = errors.New("email or password invalid")
	// ErrEmailNotVerified email not verified
	ErrEmailNotVerified = errors.New("email not verified")
	// ErrPhoneOrPasswordInvalid phone or password invalid
	ErrPhoneOrPasswordInvalid = errors.New("phone or password invalid")
	// ErrPhoneNotVerified phone not verified
	ErrPhoneNotVerified = errors.New("phone not verified")

	regexAlphabet = regexp.MustCompile(strAlphabet)
	regexEmail    = regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")
	regexPhone    = regexp.MustCompile(`^[1-9][0-9]{10}$`)
)

// CreateRequiredError create required error
func CreateRequiredError(name string) error {
	return createValidationError(name, "required")
}

// CreateInvalidError create invalid error
func CreateInvalidError(name string) error {
	return createValidationError(name, "invalid")
}

// ValidateModelName validate modelName
func ValidateModelName(modelName ...string) error {

	for i := 0; i < len(modelName); i++ {

		if !MatchAlphabet(modelName[i]) {
			return fmt.Errorf("modelName '%s' invalid, must match regex /%s/", modelName[i], strAlphabet)
		}
	}

	return nil
}

// ValidateEmail validate email
func ValidateEmail(email ...string) error {

	for i := 0; i < len(email); i++ {

		if !MatchEmail(email[i]) {
			return fmt.Errorf("email '%s' invalid", email[i])
		}
	}

	return nil
}

// ValidatePhone validate phone number
func ValidatePhone(phone ...string) error {

	for i := 0; i < len(phone); i++ {

		if !MatchPhone(phone[i]) {
			return fmt.Errorf("phone '%s' invalid", phone[i])
		}
	}

	return nil
}

// MatchAlphabet check alphabet
func MatchAlphabet(val string) bool {
	return regexAlphabet.MatchString(val)
}

// MatchEmail check email
func MatchEmail(val string) bool {
	return regexEmail.MatchString(val)
}

// MatchPhone check phone
func MatchPhone(val string) bool {
	return regexPhone.MatchString(val)
}

// createValidationError create validation error with name and message
func createValidationError(name string, message string) error {
	return fmt.Errorf("%s: %s", name, message)
}
