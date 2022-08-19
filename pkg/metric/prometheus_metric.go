package metric

import (
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/collectors"
	"github.com/prometheus/client_golang/prometheus/promhttp"

	"github.com/py4mac/fizzbuzz/pkg/constants"
)

// Prometheus holds applicative metrics
type PrometheusMetrics struct {
	Info      *prometheus.GaugeVec
	HitsTotal prometheus.Counter
	Hits      *prometheus.CounterVec
	Times     *prometheus.HistogramVec
}

// NewPrometheusMetric create a new prometheus metric
func NewPrometheusMetric(address, name string) (Metrics, error) {
	var metr PrometheusMetrics
	metr.HitsTotal = prometheus.NewCounter(prometheus.CounterOpts{
		Name: name + "_hits_total",
	})

	if err := prometheus.Register(metr.HitsTotal); err != nil {
		return nil, err
	}

	metr.Hits = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: name + "_hits",
		},
		[]string{"status", "method", "path"},
	)

	if err := prometheus.Register(metr.Hits); err != nil {
		return nil, err
	}

	metr.Times = prometheus.NewHistogramVec(
		prometheus.HistogramOpts{
			Name: name + "_times",
		},
		[]string{"status", "method", "path"},
	)

	if err := prometheus.Register(metr.Times); err != nil {
		return nil, err
	}

	metr.Info = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Name: name + "_info_gauge",
		},
		[]string{"name", "version", "revision", "built"},
	)
	metr.Info.With(
		prometheus.Labels{
			"name":     name,
			"version":  constants.Version,
			"revision": constants.Revision,
			"built":    constants.Built,
		},
	).Set(1)

	if err := prometheus.Register(metr.Info); err != nil {
		return nil, err
	}

	if err := prometheus.Register(collectors.NewBuildInfoCollector()); err != nil {
		return nil, err
	}

	go func() {
		router := echo.New()
		router.HideBanner = true
		router.GET("/metrics", echo.WrapHandler(promhttp.Handler()))
		logger.Printf("Metrics server is running on port: %s", address)

		if err := router.Start(address); err != nil {
			logger.Fatal(err)
		}
	}()

	return &metr, nil
}

// IncHits stores rest API counter
func (metr *PrometheusMetrics) IncHits(status int, method, path string) {
	metr.HitsTotal.Inc()
	metr.Hits.WithLabelValues(strconv.Itoa(status), method, path).Inc()
}

// ObserveResponseTime stores response time for rest API
func (metr *PrometheusMetrics) ObserveResponseTime(status int, method, path string, observeTime float64) {
	metr.Times.WithLabelValues(strconv.Itoa(status), method, path).Observe(observeTime)
}
