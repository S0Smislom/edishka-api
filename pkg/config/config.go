package config

import (
	"log"

	"github.com/BurntSushi/toml"
)

// Config ...
type Config struct {
	BaseAPIURL      string `toml:"base_api_url"`
	BaseAdminAPIURL string `toml:"base_admin_api_url"`

	BindAddr    string `toml:"bind_addr"`
	AdminAddr   string `toml:"admin_addr"`
	LogLevel    string `toml:"log_level"`
	DatabaseURL string `toml:"database_url"`
	SessionKey  string `toml:"session_key"`

	AccessTokenTTL int    `toml:"access_token_ttl"`
	TokenSecret    string `toml:"token_secret"`

	AdminAccessTokenTTL int    `toml:"admin_access_token_ttl"`
	AdminTokenSecret    string `toml:"admin_token_secret"`
}

// initConfig ...
func InitConfig(configPath string) (*Config, error) {
	config := &Config{
		BindAddr: ":8080",
		LogLevel: "debug",
	}
	_, err := toml.DecodeFile(configPath, config)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	return config, nil
}
