package dependency

import (
	"fmt"
	"os"
	"strconv"

	"github.com/gin-gonic/gin"
)

type Config struct {
	AppMode   string
	Port      int
	Cors      string
	SentryDSN string
}

func GetEnvStrOrDefault(key string, defaultValue string) string {
	value := os.Getenv(key)

	if value == "" {
		return defaultValue
	}

	return value
}

func GetEnvStrOrError(key string) (string, error) {
	value := os.Getenv(key)

	if value == "" {
		return "", fmt.Errorf("environment variable %s is required but not set", key)
	}

	return value, nil
}

func GetEnvIntOrDefault(key string, defaultValue int) int {
	strValue := os.Getenv(key)

	intValue, err := strconv.Atoi(strValue)
	if err != nil {
		return defaultValue
	}

	return intValue
}

func LoadConfigFromEnv() (*Config, error) {
	cfg := &Config{
		AppMode:   GetEnvStrOrDefault("APP_MODE", "debug"),
		Port:      GetEnvIntOrDefault("PORT", 8080),
		Cors:      GetEnvStrOrDefault("CORS", "http://localhost:5173"),
		SentryDSN: GetEnvStrOrDefault("SENTRY_DSN", ""),
	}

	if cfg.AppMode == "production" {
		gin.SetMode(gin.ReleaseMode)
	} else {
		gin.SetMode(gin.DebugMode)
	}

	return cfg, nil
}
