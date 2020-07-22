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
