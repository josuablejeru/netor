package main

import (
	"net/http"

	"github.com/josuablejeru/netor/collector"
	log "github.com/sirupsen/logrus"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func main() {
	prometheus.MustRegister(collector.NewSpeedTest())
	http.Handle("/metrics", promhttp.Handler())

	log.Info("Beginning to serve on port ':8080'")
	log.Info("Serve prometheus exporter on '/metrics'")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
