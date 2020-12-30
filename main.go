package main

import (
	"net/http"

	"github.com/josuablejeru/netor/collector"
	"github.com/sirupsen/logrus"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func main() {
	prometheus.MustRegister(collector.NewSpeedTest())
	http.Handle("/metrics", promhttp.Handler())

	logrus.Info("Beginning to serve on port :8080")
	logrus.Fatal(http.ListenAndServe(":8080", nil))
}
