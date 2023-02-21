// Copyright 2023 The GoStartKit Authors. All rights reserved.
// Use of this source code is governed by a AGPL
// license that can be found in the LICENSE file.
// https://gostartkit.com
package config

import (
	"fmt"
	"time"
)

// CreateServerConfig create server config
func CreateServerConfig() *ServerConfig {
	cfg := &ServerConfig{
		Network:           "unix",
		Addr:              fmt.Sprintf("./log/%s.sock", Key()),
		ReadTimeout:       32,
		ReadHeaderTimeout: 8,
		WriteTimeout:      32,
		IdleTimeout:       8,
	}
	return cfg
}

// ServerConfig struct
type ServerConfig struct {
	Network           string
	Addr              string
	ReadTimeout       time.Duration
	ReadHeaderTimeout time.Duration
	WriteTimeout      time.Duration
	IdleTimeout       time.Duration
}
