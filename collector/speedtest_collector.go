package collector

import (
	"sync"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/showwin/speedtest-go/speedtest"
)

const (
	namespace = "netor"
	subsystem = "speed"
)

// SpeedTest type defines the collector struct
type SpeedTest struct {
	up                              prometheus.Gauge
	totalScrapes, jsonParseFailures prometheus.Counter
	targets                         speedtest.Servers
	metrics                         []*SpeedTestMetric
	mutex                           sync.RWMutex
}

// SpeedTestMetric implements the SpeedTestMetric Interface from Prometheus
type SpeedTestMetric struct {
	Desc *prometheus.Desc
	Type prometheus.ValueType

	Value func(server *speedtest.Server) float64
}

// NewSpeedTest returns a new Collector exposing Speedtest stats
func NewSpeedTest() *SpeedTest {
	targets := GetTargets()
	return &SpeedTest{
		up: prometheus.NewGauge(prometheus.GaugeOpts{
			Name: prometheus.BuildFQName(namespace, subsystem, "up"),
			Help: "Was the last scrape of the NETOR endpoint successful.",
		}),
		totalScrapes: prometheus.NewCounter(prometheus.CounterOpts{
			Name: prometheus.BuildFQName(namespace, subsystem, "total_scrapes"),
			Help: "Current total Speedtest scrapes.",
		}),
		jsonParseFailures: prometheus.NewCounter(prometheus.CounterOpts{
			Name: prometheus.BuildFQName(namespace, subsystem, "json_parse_failures"),
			Help: "Number of errors while parsing JSON.",
		}),
		targets: targets,
		metrics: []*SpeedTestMetric{
			{
				Type: prometheus.GaugeValue,
				Desc: prometheus.NewDesc(
					prometheus.BuildFQName(namespace, subsystem, "latency"),
					"The Latency (PING) of the network request to the server", nil, nil,
				),
				Value: func(server *speedtest.Server) float64 {
					return float64(server.Latency)
				},
			},
			{
				Type: prometheus.GaugeValue,
				Desc: prometheus.NewDesc(
					prometheus.BuildFQName(namespace, subsystem, "download"),
					"The Downloadspeed calculated by the server", nil, nil,
				),
				Value: func(server *speedtest.Server) float64 {
					return float64(server.DLSpeed)
				},
			},
			{
				Type: prometheus.GaugeValue,
				Desc: prometheus.NewDesc(
					prometheus.BuildFQName(namespace, subsystem, "upload"),
					"The Uploadspeed calculated by the server", nil, nil,
				),
				Value: func(server *speedtest.Server) float64 {
					return float64(server.ULSpeed)
				},
			},
		},
	}
}

func (s *SpeedTest) runAndFetchSpeedtestResult() *speedtest.Server {
	for _, t := range s.targets {
		go StartSpeedtesting(t)
		return t
	}
	return nil
}

// Describe sends the Desc to Metrix regristry.
// It implements prometheus.Describe
func (s *SpeedTest) Describe(ch chan<- *prometheus.Desc) {
	for _, metric := range s.metrics {
		ch <- metric.Desc
	}
	ch <- s.up.Desc()
	ch <- s.totalScrapes.Desc()
	ch <- s.jsonParseFailures.Desc()
}

// Collect fetches the data
// It implements prometheus.Collect
func (s *SpeedTest) Collect(ch chan<- prometheus.Metric) {
	s.mutex.Lock()
	s.totalScrapes.Inc()

	defer func() {
		ch <- s.up
		ch <- s.totalScrapes
		ch <- s.jsonParseFailures
		s.mutex.Unlock()
	}()

	serverResult := s.runAndFetchSpeedtestResult()

	for _, metric := range s.metrics {
		ch <- prometheus.MustNewConstMetric(
			metric.Desc,
			metric.Type,
			metric.Value(serverResult),
		)
	}
}
