// Copyright 2023 The GoStartKit Authors. All rights reserved.
// Use of this source code is governed by a AGPL
// license that can be found in the LICENSE file.
// https://gostartkit.com
package config

import "github.com/webpkg/api/helper"

// CreateAppConfig create app config
func CreateAppConfig() *AppConfig {
	cfg := &AppConfig{
		AppID:                      1,
		AppNum:                     4,
		AppName:                    "api",
		AppEnv:                     "local",
		AppKey:                     helper.CreateToken(),
		AppDebug:                   false,
		Domain:                     "gostartkit.com",
		PublicDir:                  "public",
		StorageDir:                 "storage",
		ResourceDir:                "resource",
		TimeLocation:               "Asia/Shanghai",
		TimeLayout:                 "2006-01-02 15:04:05",
		TokenExpireDuration:        3600 * 2,
		RefreshTokenExpireDuration: 3600 * 24 * 365,
	}

	return cfg
}

// AppConfig struct
type AppConfig struct {
	AppID                      uint64
	AppNum                     uint64
	AppName                    string
	AppEnv                     string
	AppKey                     string
	AppDebug                   bool
	Domain                     string
	PublicDir                  string
	StorageDir                 string
	ResourceDir                string
	TimeLocation               string
	TimeLayout                 string
	TokenExpireDuration        uint
	RefreshTokenExpireDuration uint
}
