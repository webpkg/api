package main

import (
	"errors"
	"log"

	"github.com/webpkg/web"
	"github.com/webpkg/cmd"
)

const (
	_configFile        = "config.json"
	_addr              = "127.0.0.1:8443"
	_readTimeout       = 32
	_readHeaderTimeout = 8
	_writeTimeout      = 32
	_idleTimeout       = 8

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

// config struct
type config struct {
	web.Config
}

func runConfig(cmd *cmd.Command, args []string) {

	if len(args) != 0 {
		log.Fatal("Too many arguments given.\n")
	}

	writeConfig()
}

// writeConfig create new config.json at $configDir
func writeConfig() {

	cfg := &config{}

	cfg.Server = &web.ServerConfig{
		Addr:              _addr,
		ReadTimeout:       _readTimeout,
		ReadHeaderTimeout: _readHeaderTimeout,
		WriteTimeout:      _writeTimeout,
		IdleTimeout:       _idleTimeout,
	}

	cfg.Database = &web.DatabaseCluster{
		Driver:    _connection,
		Database:  _database,
		Username:  _username,
		Password:  _password,
		Charset:   _charset,
		Collation: _collation,
	}

	cfg.Database.Write = &web.DatabaseHostConfig{
		Host: _host,
		Port: _port,
	}

	cfg.Database.Read = &[]web.DatabaseHostConfig{
		{
			Host: _host,
			Port: _port,
		},
		{
			Host: _host,
			Port: _port,
		},
	}

	if err := web.WriteJSON(cfg, _configFile, _webForce); err != nil {
		log.Printf("config: %v", err)
	}
}

func readConfig() (*config, error) {

	c := &config{}

	err := web.ReadJSON(c, _configFile)

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
