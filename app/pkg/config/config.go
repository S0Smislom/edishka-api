package config

import (
	"log"

	"github.com/BurntSushi/toml"
)

// Config ...
type Config struct {
	BaseHost        string `toml:"base_host"`
	BaseAPIURL      string `toml:"base_api_url"`
	BaseAdminAPIURL string `toml:"base_admin_api_url"`

	BindAddr    string `toml:"bind_addr"`
	AdminAddr   string `toml:"admin_addr"`
	LogLevel    string `toml:"log_level"`
	DatabaseURL string `toml:"database_url"`
	SessionKey  string `toml:"session_key"`

	AccessTokenTTL  int    `toml:"access_token_ttl"`
	RefreshTokenTTL int    `toml:"refresh_token_ttl"`
	TokenSecret     string `toml:"token_secret"`

	AdminAccessTokenTTL int    `toml:"admin_access_token_ttl"`
	AdminTokenSecret    string `toml:"admin_token_secret"`

	MinioEndpoint  string `toml:"minio_endpoint"`
	MinioAccessKey string `toml:"minio_access_key"`
	MinioSecretKey string `toml:"minio_secret_key"`
	MinioUseSSL    bool   `toml:"minio_use_ssl"`
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

// Init test config
func InitTestConfig() (*Config, error) {
	config := &Config{
		AccessTokenTTL:  60,
		RefreshTokenTTL: 60,
		TokenSecret:     "test",
		DatabaseURL:     "host=localhost user=custom password=qwerty123 dbname=test port=5454 sslmode=disable",
	}
	return config, nil
}
