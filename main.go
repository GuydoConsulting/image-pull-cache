package main

import (
	"net/http"

	"github.com/sirupsen/logrus"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/mutate", handleMutate)

	server := &http.Server{
		Addr:    ":" + config.ServerPort,
		Handler: mux,
	}

	log.WithFields(logrus.Fields{
		"service":     "webhook-server",
		"port":        config.ServerPort,
		"tls_enabled": config.TLSEnabled,
	}).Info("Starting webhook server")

	if config.TLSEnabled {
		server.TLSConfig = getTLSConfig() // Set the TLS configuration
		log.Fatal(server.ListenAndServeTLS(config.TLSCertPath, config.TLSKeyPath))
	} else {
		log.Fatal(server.ListenAndServe())
	}
}
