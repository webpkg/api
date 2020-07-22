package config

import "github.com/webpkg/api/helper"

// CreateAppConfig create app config
func CreateAppConfig() *AppConfig {
	cfg := &AppConfig{
		AppID:       1,
		AppNum:      2,
		AppName:     "web.go",
		AppEnv:      "local",
		AppKey:      helper.CreateToken(),
		AppDebug:    true,
		AppURL:      "http://127.0.0.1",
		PublicDir:   "public",
		StorageDir:  "storage",
		ResourceDir: "resource",
	}

	return cfg
}

// AppConfig struct
type AppConfig struct {
	AppID       uint64
	AppNum      uint64
	AppName     string
	AppEnv      string
	AppKey      string
	AppDebug    bool
	AppURL      string
	PublicDir   string
	StorageDir  string
	ResourceDir string
}
