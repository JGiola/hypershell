package config

import (
	"fmt"
	"os"
	"strings"
)

type Config struct {
	GRPCServerAddr string
	APIServerURL   string
	Namespace      string
	LogLevel       string
}

func Load() (*Config, error) {
	cfg := &Config{
		GRPCServerAddr: getEnv("HYPERSHELL_GRPC_SERVER_ADDR", "localhost:9000"),
		APIServerURL:   getEnv("HYPERSHELL_API_SERVER_URL", "http://localhost:8000"),
		Namespace:      getEnv("HYPERSHELL_NAMESPACE", "hypershell"),
		LogLevel:       strings.ToLower(getEnv("HYPERSHELL_LOG_LEVEL", "info")),
	}

	if cfg.GRPCServerAddr == "" {
		return nil, fmt.Errorf("HYPERSHELL_GRPC_SERVER_ADDR is required")
	}

	return cfg, nil
}

func getEnv(key, fallback string) string {
	if v := os.Getenv(key); v != "" {
		return v
	}
	return fallback
}
