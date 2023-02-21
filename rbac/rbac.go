// Copyright 2023 The GoStartKit Authors. All rights reserved.
// Use of this source code is governed by a AGPL
// license that can be found in the LICENSE file.
// https://gostartkit.com
package rbac

import (
	"strings"
	"sync"

	"github.com/gostartkit/api/config"
)

var (
	_rights *config.RbacConfig
	_once   sync.Once
)

// Init config
func Init() {
	_once.Do(func() {
		_rights = config.Rbac()
	})
}

// ParseBearerToken return token
func ParseBearerToken(auth string) string {
	const prefix = "Bearer "
	l := len(prefix)

	if len(auth) < l || !strings.EqualFold(auth[:l], prefix) {
		return ""
	}

	return auth[l:]
}

// Check check right
func Check(userRight int64, keys ...string) bool {

	if len(keys) == 0 {
		return userRight > 0
	}

	return checkKeys(userRight, keys...)
}

// checkKeys check right by keys
func checkKeys(userRight int64, keys ...string) bool {

	for _, key := range keys {
		if !checkKey(userRight, key) {
			return false
		}
	}

	return true
}

// checkKey check right by key
func checkKey(userRight int64, key string) bool {
	orKeys := strings.Split(key, "|")

	for _, orKey := range orKeys {

		if orKey == "" {
			return userRight > 0
		}

		val := getVal(orKey)

		if val > 0 {
			return val&userRight > 0
		}
	}

	return false
}

// getVal get rights value by key
func getVal(key string) int64 {
	return _rights.Search(key)
}
