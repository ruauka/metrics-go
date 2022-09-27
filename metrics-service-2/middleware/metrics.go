package middleware

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

// Metrics prometheus + registry.
type Metrics struct {
	// RequestsTotal metric http_requests_total.
	Total *prometheus.CounterVec
	// RequestsDuration metric http_request_duration_seconds.
	Duration *prometheus.HistogramVec
}

// NewRegistry create new store metrics prometheus.
func (m Metrics) NewRegistry() *prometheus.Registry {
	registry := prometheus.NewRegistry()

	registry.MustRegister(
		m.Duration,
		m.Total,
	)
	return registry
}

// NewMetrics returns all metrics for prometheus.
func NewMetrics() *Metrics {
	var metrics Metrics
	labels := []string{"status_code"}

	metrics.Total = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "requests_total",
			Help: "Total number http requests",
		},
		labels,
	)

	metrics.Duration = prometheus.NewHistogramVec(prometheus.HistogramOpts{
		Name: "request_duration_seconds",
		Help: "Latency http requests.",
		// buckets from 10ms to 5s
		Buckets: []float64{.01, .02, .03, .04, .05, .1, .2, .3, .5, 1.0, 1.5, 2.0, 2.5, 3.0, 4.0, 5.0},
	},
		labels,
	)

	return &metrics
}

// NewMetricsHandler returns Handler for endpoint `/metrics`.
func NewMetricsHandler(metrics *Metrics) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		promhttp.HandlerFor(metrics.NewRegistry(), promhttp.HandlerOpts{}).ServeHTTP(w, r)
	}
}
