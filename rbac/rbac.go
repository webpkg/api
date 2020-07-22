package rbac

import (
	"errors"
	"sort"
	"strings"
	"sync"

	"github.com/webpkg/api/config"
)

var (
	// ErrInvalidToken parse token error
	ErrInvalidToken = errors.New("invalid token")
	// ErrTokenEmpty token value is empty
	ErrTokenEmpty = errors.New("token value is empty")
	// ErrPermissionDenied do not have permission access this resource
	ErrPermissionDenied = errors.New("permission denied")

	_rights *config.RbacConfig
	_once   sync.Once
)

// Init config
func Init(rbacConfig *config.RbacConfig) {
	_once.Do(func() {
		_rights = rbacConfig
		sort.Sort(_rights)
	})
}

// TryParseBearerToken return token
func TryParseBearerToken(auth string) (string, error) {
	const prefix = "Bearer "

	if len(auth) < len(prefix) || !strings.EqualFold(auth[:len(prefix)], prefix) {
		return "", ErrInvalidToken
	}

	token := auth[len(prefix):]

	if token == "" {
		return "", ErrTokenEmpty
	}

	return token, nil
}

// Check check right
func Check(userRight int64, keys ...string) bool {
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
