package main

import (
	"github.com/prometheus/client_golang/prometheus"
	"sync"
)

// Collector implements the Collector Interface from Prometheus
type Collector struct {
	Desc *prometheus.Desc
	Type prometheus.ValueType

	mutex sync.RWMutex
}

var (
	pingMetric = New("current_ping", "Current internet ping", prometheus.GaugeValue, nil)
)

// New creates a Metric
func New(metricName string, docString string, t prometheus.ValueType, constLabels prometheus.Labels) *Collector {
	return &Collector{
		Desc: prometheus.NewDesc(
			prometheus.BuildFQName("netor", "netSpeed", metricName),
			docString,
			[]string{},
			constLabels,
		),
		Type: t,
	}
}

// Describe sends the Desc to Metrix regristry.
// It implements prometheus.Describe
func (c *Collector) Describe(ch chan<- *prometheus.Desc) {
	ch <- c.Desc
}

// Collect fetches the data
// It implements prometheus.Collect
func (c *Collector) Collect(ch chan<- prometheus.Metric) {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	ch <- prometheus.MustNewConstMetric(
		c.Desc, prometheus.GaugeValue, 1,
	)
}
