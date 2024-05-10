package main

import (
	"crypto/tls"
	"os"
	"strconv"

	"github.com/sirupsen/logrus"
)

var log = logrus.New()
var config = newConfig()

func init() {
	// Log as JSON instead of the default ASCII formatter.
	log.Formatter = new(logrus.TextFormatter)

	// Output to stdout instead of the default stderr
	log.Out = os.Stdout

	// Only log the warning severity or above.
	log.Level = logrus.InfoLevel
}

type Config struct {
	ServerPort   string
	TLSCertPath  string
	TLSKeyPath   string
	PullCacheURL string
	TLSEnabled   bool
}

func newConfig() *Config {
	tlsEnabled, err := strconv.ParseBool(getEnv("TLS_ENABLED", "true")) // Default to TLS enabled
	if err != nil {
		log.WithError(err).Fatal("Invalid TLS_ENABLED value")
	}
	return &Config{
		ServerPort:   getEnv("SERVER_PORT", "8443"),
		TLSCertPath:  getEnv("TLS_CERT_PATH", "/pki/tls.crt"),
		TLSKeyPath:   getEnv("TLS_KEY_PATH", "/pki/tls.key"),
		PullCacheURL: getEnv("PULL_CACHE_URL", ""),
		TLSEnabled:   tlsEnabled,
	}
}

func getEnv(key, fallback string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	log.Warnf("%s env var was not set, using default value %s", key, fallback)
	return fallback
}

func getTLSConfig() *tls.Config {
	return &tls.Config{
		MinVersion: tls.VersionTLS12,
	}
}
