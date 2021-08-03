package config

import "os"

// AppConfig - our app specific config
type AppConfig struct {
	Name    string
	Version string
	Port    string
}

// HTTPConfig - our Http config values
type HTTPConfig struct {
	Content string
	Problem string
}

// Config - our application global config struct
type Config struct {
	App  *AppConfig
	HTTP *HTTPConfig
}

func Load() *Config {
	return &Config{
		App: &AppConfig{
			Name:    env("APP_NAME", "Go App"),
			Version: env("APP_VERSION", "0.0.1"),
			Port:    env("APP_PORT", "8080"),
		},
		HTTP: &HTTPConfig{
			Content: env("HTTP_CONTENT_TYPE", "application/json"),
			Problem: env("HTTP_PROBLEM", "application/problem+json"),
		},
	}
}

func env(key string, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}

	return defaultValue
}
