// Package metrics holds applicative metrics implementation
package metric

// Metrics holds applicative interface for the metrics
type Metrics interface {
	IncHits(status int, method, path string)
	ObserveResponseTime(status int, method, path string, observeTime float64)
}
