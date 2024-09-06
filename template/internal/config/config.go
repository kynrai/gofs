package config

import (
	"cmp"
	"fmt"
	"os"
	"strings"
)

type Config struct {
	ServiceName    string
	Host           string
	Port           string
	AllowedOrigins []string
	Env            Environment
	DSN            string
	ICN            string
	Tracing        string // url to tracing server, e.g. zipkin
	Metrics        bool   // enable prometheus metrics
}

func New() Config {
	host := getEnvDefault("HOST", "0.0.0.0")
	port := getEnvDefault("PORT", "8080")
	return Config{
		ServiceName: "ttz-app",
		Host:        host,
		Port:        port,
		AllowedOrigins: strings.Split(
			getEnvDefault("ALLOWED_ORIGINS", fmt.Sprintf("http://%s:%s,https://%s:%s", host, port, host, port)), ",",
		),
		Env:     Environment(getEnvDefault("ENV", "prod")),
		DSN:     os.Getenv("DSN"),
		ICN:     os.Getenv("ICN"),
		Tracing: os.Getenv("TRACING"),
		Metrics: os.Getenv("METRICS") == "true",
	}
}

func getEnvDefault(key, def string) string {
	return cmp.Or(os.Getenv(key), def)
}
