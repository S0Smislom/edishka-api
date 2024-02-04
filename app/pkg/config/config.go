package config

import (
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/BurntSushi/toml"
	"github.com/joho/godotenv"
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

func InitEnvConfig(configPath string) (*Config, error) {
	if err := godotenv.Load(configPath); err != nil {
		log.Print("No .env file found")
		return nil, err
	}
	return &Config{
		BaseHost:       getEnv("BASE_HOST", "localhost:8091"),
		BaseAPIURL:     getEnv("BASE_API_URL", "http://localhost:8091"),
		BindAddr:       getEnv("BIND_ADDR", ":8091"),
		LogLevel:       getEnv("LOG_LEVEL", "debug"),
		DatabaseURL:    getEnv("DATABASE_URL", ""),
		SessionKey:     getEnv("TOKEN_SECRET", ""),
		AccessTokenTTL: getEnvAsInt("ADMIN_ACCESS_TOKEN_TTL", 24),
		MinioEndpoint:  getEnv("MINIO_ENDPOINT", "localhost:9004"),
		MinioAccessKey: getEnv("MINIO_ACCESS_KEY", "root"),
		MinioSecretKey: getEnv("MINIO_SECRET_KEY", "changeme"),
		MinioUseSSL:    getEnvAsBool("MINIO_USE_SSL", false),
	}, nil
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

// Simple helper function to read an environment or return a default value
func getEnv(key string, defaultVal string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultVal
}

// Simple helper function to read an environment variable into integer or return a default value
func getEnvAsInt(name string, defaultVal int) int {
	valueStr := getEnv(name, "")
	if value, err := strconv.Atoi(valueStr); err == nil {
		return value
	}

	return defaultVal
}

// Helper to read an environment variable into a bool or return default value
func getEnvAsBool(name string, defaultVal bool) bool {
	valStr := getEnv(name, "")
	if val, err := strconv.ParseBool(valStr); err == nil {
		return val
	}

	return defaultVal
}

// Helper to read an environment variable into a string slice or return default value
func getEnvAsSlice(name string, defaultVal []string, sep string) []string {
	valStr := getEnv(name, "")

	if valStr == "" {
		return defaultVal
	}

	val := strings.Split(valStr, sep)

	return val
}
