package config

import (
	"errors"
	"log"

	"github.com/webpkg/api/helper"
)

const (
	_configFile   = "config.json"
)

// WebConfig struct
type WebConfig struct {
	App      *AppConfig
	Server   *ServerConfig
	Database *DatabaseCluster
	Auth     *AuthConfig
	Rbac     *RbacConfig
}

// WriteConfig create new config.json at $configDir
func WriteConfig() {

	cfg := &WebConfig{}

	cfg.App = CreateAppConfig()

	cfg.Server = CreateServerConfig()

	cfg.Database = CreateDatabaseClusterConfig()

	cfg.Auth = CreateAuthConfig()

	// cfg.Rbac = CreateRbacConfig()

	if err := helper.WriteJSON(cfg, _configFile, false); err != nil {
		log.Printf("config: %v", err)
	}
}

// ReadConfig read $configDir/config.json
func ReadConfig() (*WebConfig, error) {

	c := &WebConfig{}

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

	if c.Auth == nil {
		return nil, errors.New("config.Auth is nil")
	}

	if c.Rbac == nil {
		c.Rbac = CreateRbacConfig()
	}

	return c, nil
}
