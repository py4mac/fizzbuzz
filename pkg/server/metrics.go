package server

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/collectors"

	"github.com/py4mac/fizzbuzz/pkg/constants"
)

var (
	hitsTotal = prometheus.NewCounter(prometheus.CounterOpts{
		Name: constants.Name + "_hits_total",
	})
	hits = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: constants.Name + "_hits",
		},
		[]string{"status", "method", "path"},
	)
	times = prometheus.NewHistogramVec(
		prometheus.HistogramOpts{
			Name: constants.Name + "_times",
		},
		[]string{"status", "method", "path"},
	)
	info = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Name: constants.Name + "_info_gauge",
		},
		[]string{"name", "version", "revision", "built"},
	)
)

func init() {
	prometheus.MustRegister(hitsTotal)
	prometheus.MustRegister(hits)
	prometheus.MustRegister(times)
	prometheus.MustRegister(info)
	prometheus.MustRegister(collectors.NewBuildInfoCollector())
	info.With(
		prometheus.Labels{
			"name":     constants.Name,
			"version":  constants.Version,
			"revision": constants.Revision,
			"built":    constants.Built,
		},
	).Set(1)
}
