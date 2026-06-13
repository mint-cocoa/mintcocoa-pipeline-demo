package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
	"time"
)

var version = "dev"

type releaseResponse struct {
	Service  string `json:"service"`
	Version  string `json:"version"`
	Env      string `json:"env"`
	Hostname string `json:"hostname"`
	Time     string `json:"time"`
}

func main() {
	mux := http.NewServeMux()
	healthHandler := func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("content-type", "text/plain; charset=utf-8")
		_, _ = w.Write([]byte("ok\n"))
	}
	mux.HandleFunc("/healthz", healthHandler)
	mux.HandleFunc("/readyz", healthHandler)
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		hostname, _ := os.Hostname()
		appVersion := getenv("APP_VERSION", version)
		env := getenv("APP_ENV", "local")

		w.Header().Set("content-type", "application/json")
		_ = json.NewEncoder(w).Encode(releaseResponse{
			Service:  "pipeline-demo",
			Version:  appVersion,
			Env:      env,
			Hostname: hostname,
			Time:     time.Now().UTC().Format(time.RFC3339),
		})
	})

	addr := ":8080"
	log.Printf("pipeline-demo listening on %s", addr)
	if err := http.ListenAndServe(addr, mux); err != nil {
		log.Fatal(err)
	}
}

func getenv(name, fallback string) string {
	value := os.Getenv(name)
	if value == "" {
		return fallback
	}
	return value
}
