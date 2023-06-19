package metrics

import (
	"sync"

	"github.com/prometheus/client_golang/prometheus"
)

var (
	ApiResponseTime = prometheus.NewHistogramVec(
		prometheus.HistogramOpts{
			Name: "api_response_time_seconds",
			Help: "Histogram to observe api response time",
		}, []string{"method"},
	)
	ApiTotalRequests = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "api_requests_total",
			Help: "Number of get requests",
		},
		[]string{"method"},
	)
)

var once sync.Once

func InitApiPrometheusMetrics() {
	once.Do(func() {
		prometheus.MustRegister(ApiResponseTime)
		prometheus.MustRegister(ApiTotalRequests)
	})
}
