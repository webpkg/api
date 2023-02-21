// Copyright 2023 The GoStartKit Authors. All rights reserved.
// Use of this source code is governed by a AGPL
// license that can be found in the LICENSE file.
// https://gostartkit.com
package config

// CreateAuthConfig return *AuthConfig
func CreateAuthConfig() *AuthConfig {

	cfg := &AuthConfig{
		AppKey:    "",
		AppSecret: "",
		Addr: []string{
			"http://127.0.0.1",
			"http://127.0.0.1",
		},
	}

	return cfg
}

// AuthConfig struct
type AuthConfig struct {
	AppKey    string
	AppSecret string
	Addr      []string
}
