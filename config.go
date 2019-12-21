package main

import (
	"errors"
	"log"
	"time"

	"github.com/webpkg/api/config"
	"github.com/webpkg/api/helper"
	"github.com/webpkg/cmd"
)

const (
	_configFile = "config.json"

	_connection = "mysql"
	_host       = "127.0.0.1"
	_port       = 3306
	_database   = "webgo"
	_username   = "webgo"
	_password   = "webgo@8443"
	_charset    = "utf8"
	_collation  = "utf8_general_ci"
)

var (
	cmdConfig = &cmd.Command{
		Run:       runConfig,
		UsageLine: "config",
		Short:     "create config file",
		Long:      "create config.json file at current directory.\n",
	}
)

// Config struct
type Config struct {
	Server   *config.ServerConfig
	Database *config.DatabaseCluster
	Redis    *config.RedisConfig
}

func runConfig(cmd *cmd.Command, args []string) {

	if len(args) != 0 {
		log.Fatal("Too many arguments given.\n")
	}

	writeConfig()
}

// writeConfig create new config.json at $configDir
func writeConfig() {

	cfg := &Config{}

	cfg.Server = &config.ServerConfig{
		Addr:              "127.0.0.1:8443",
		ReadTimeout:       32 * time.Second,
		ReadHeaderTimeout: 8 * time.Second,
		WriteTimeout:      32 * time.Second,
		IdleTimeout:       8 * time.Second,
	}

	cfg.Database = &config.DatabaseCluster{
		Driver:    _connection,
		Database:  _database,
		Username:  _username,
		Password:  _password,
		Charset:   _charset,
		Collation: _collation,
	}

	cfg.Database.Write = &config.DatabaseHostConfig{
		Host: _host,
		Port: _port,
	}

	cfg.Database.Read = &[]config.DatabaseHostConfig{
		{
			Host: _host,
			Port: _port,
		},
		{
			Host: _host,
			Port: _port,
		},
	}

	if err := helper.WriteJSON(cfg, _configFile, _webForce); err != nil {
		log.Printf("config: %v", err)
	}
}

func readConfig() (*Config, error) {

	c := &Config{}

	err := helper.ReadJSON(c, _configFile)

	if err != nil {
		return nil, err
	}

	if c.Server == nil {
		return nil, errors.New("config.Server is nil")
	}

	if c.Database == nil {
		return nil, errors.New("config.Database is nil")
	}

	if c.Database.Write == nil {
		return nil, errors.New("config.Database.Write is nil")
	}

	return c, nil
}
