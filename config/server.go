package config

import "time"

// CreateServerConfig create server config
func CreateServerConfig() *ServerConfig {
	cfg := &ServerConfig{
		Addr:              "web:8443",
		ReadTimeout:       32,
		ReadHeaderTimeout: 8,
		WriteTimeout:      32,
		IdleTimeout:       8,
	}
	return cfg
}

// ServerConfig struct
type ServerConfig struct {
	Addr              string
	ReadTimeout       time.Duration
	ReadHeaderTimeout time.Duration
	WriteTimeout      time.Duration
	IdleTimeout       time.Duration
}
