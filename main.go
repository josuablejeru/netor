package main

import (
	"net/http"

	log "github.com/sirupsen/logrus"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func main() {
	prometheus.MustRegister(pingMetric)
	http.Handle("/metrics", promhttp.Handler())

	log.Info("Beginning to serve on port :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
