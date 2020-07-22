package main

import (
	"errors"
	"log"
	"path/filepath"

	"github.com/webpkg/api/config"
	"github.com/webpkg/api/helper"
	"github.com/webpkg/cmd"
)

const (
	_configFile   = "config.json"
	_idConfigFile = "app/keys/id.json"
)

var (
	cmdConfig = &cmd.Command{
		Run:       runConfig,
		UsageLine: "config",
		Short:     "create config file",
		Long:      "create config.json file at current directory.\n",
	}
)

// WebConfig struct
type WebConfig struct {
	App      *config.AppConfig
	Server   *config.ServerConfig
	Database *config.DatabaseCluster
	Auth     *config.AuthConfig
	Rbac     *config.RbacConfig
}

func runConfig(cmd *cmd.Command, args []string) {

	if len(args) != 0 {
		log.Fatal("Too many arguments given.\n")
	}

	writeConfig()
}

// writeConfig create new config.json at $configDir
func writeConfig() {

	cfg := &WebConfig{}

	cfg.App = config.CreateAppConfig()

	cfg.Server = config.CreateServerConfig()

	cfg.Database = config.CreateDatabaseClusterConfig()

	cfg.Auth = config.CreateAuthConfig()

	cfg.Rbac = config.CreateRbacConfig()

	if err := helper.WriteJSON(cfg, _configFile, _webForce); err != nil {
		log.Printf("config: %v", err)
	}
}

func readConfig() (*WebConfig, error) {

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
		return nil, errors.New("config.Rbac is nil")
	}

	return c, nil
}

func readIDConfig(dir string) (*config.IDConfig, error) {

	filename := filepath.Join(dir, _idConfigFile)

	c := &config.IDConfig{}

	if helper.FileExist(filename) {

		err := helper.ReadJSON(c, filename)

		if err != nil {
			return nil, err
		}
	}

	return c, nil
}

func writeIDConfig(dir string, kc *config.IDConfig) error {

	filename := filepath.Join(dir, _idConfigFile)

	if err := helper.WriteJSON(kc, filename, true); err != nil {
		return err
	}

	return nil
}
