package api

import (
	"prometheus-go/metrics"
	"time"

	"github.com/prometheus/client_golang/prometheus"
)

type Api interface {
	StopWithError(error)
}

type ResponseTime struct {
	metric    *prometheus.HistogramVec
	startTime time.Time
}

func NewApiTimer() *ResponseTime {
	return &ResponseTime{
		metric:    metrics.ApiResponseTime,
		startTime: time.Now(),
	}
}

func (r *ResponseTime) StopWithMethod(method string) {
	labels := prometheus.Labels{"method": method}

	r.metric.With(labels).Observe(time.Since(r.startTime).Seconds())
}
