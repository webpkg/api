// Copyright 2023 The GoStartKit Authors. All rights reserved.
// Use of this source code is governed by a AGPL
// license that can be found in the LICENSE file.
// https://gostartkit.com
package config

import (
	"errors"
	"fmt"
	"sort"
	"sync"
	"time"

	"github.com/gostartkit/api/helper"
)

const (
	_configFile = "config.json"
	_key        = "api"
)

var (
	_authUrl      string = ""
	_timeLocation *time.Location
	_webConfig    *webConfig
	_once         sync.Once
)

// Init config
func Init() error {

	var err error

	_once.Do(func() {
		_webConfig, err = readConfig()

		if err == nil {
			_timeLocation, err = time.LoadLocation(_webConfig.App.TimeLocation)
		}
	})

	return err
}

// App get AppConfig
func App() *AppConfig {
	return _webConfig.App
}

// Server get ServerConfig
func Server() *ServerConfig {
	return _webConfig.Server
}

// Database get DatabaseClusterConfig
func Database() *DatabaseCluster {
	return _webConfig.Database
}

// Rbac get RbacConfig
func Rbac() *RbacConfig {
	return _webConfig.Rbac
}

// Key get Key
func Key() string {
	return _key
}

// TimeLocation get time location
func TimeLocation() *time.Location {
	return _timeLocation
}

// TimeLayout get time layout
func TimeLayout() string {

	if _webConfig.App.TimeLayout == "" {
		return "2006-01-02 15:04:05"
	}

	return _webConfig.App.TimeLayout
}

// AuthUrl get auth url
func AuthUrl() string {

	if _authUrl == "" {

		env := _webConfig.App.AppEnv
		domain := _webConfig.App.Domain

		switch env {
		case "prod", "product":
			_authUrl = fmt.Sprintf("https://api.%s/auth/authorize/%s", domain, Key())
		default:
			_authUrl = fmt.Sprintf("https://%s-api.%s/auth/authorize/%s", env, domain, Key())
		}
	}

	return _authUrl
}

// webConfig struct
type webConfig struct {
	App      *AppConfig
	Server   *ServerConfig
	Database *DatabaseCluster
	Rbac     *RbacConfig
}

// WriteConfig create new config.json at $configDir
func WriteConfig(force bool) error {

	cfg := &webConfig{}

	cfg.App = CreateAppConfig()

	cfg.Server = CreateServerConfig()

	cfg.Database = CreateDatabaseClusterConfig()

	// cfg.Rbac = CreateRbacConfig()

	if err := helper.WriteJSON(_configFile, cfg, force); err != nil {
		return err
	}

	return nil
}

// readConfig read $configDir/config.json
func readConfig() (*webConfig, error) {

	c := &webConfig{}

	if helper.FileExist(_configFile) {

		if err := helper.ReadJSON(_configFile, c); err != nil {
			return nil, err
		}
	}

	if c.App == nil {
		c.App = CreateAppConfig()
	}

	if c.Server == nil {
		c.Server = CreateServerConfig()
	}

	if c.Database == nil {
		c.Database = CreateDatabaseClusterConfig()
	}

	if c.Database.Write == nil {
		return nil, errors.New("config.Database.Write is nil")
	}

	if c.Database.Read == nil {
		return nil, errors.New("config.Database.Read is nil")
	}

	if len(*c.Database.Read) == 0 {
		return nil, errors.New("config.Database.Read len is 0")
	}

	if c.Rbac == nil {
		c.Rbac = CreateRbacConfig()
	}

	sort.Sort(c.Rbac)

	return c, nil
}
