package main

import (
	"github.com/gin-gonic/gin"
	"math/rand"
	"net/http"
	"prometheus-go/metrics"
	"prometheus-go/metrics/api"
	"prometheus-go/pkg/monitoring"
	"time"
)

const listenAddr = ":9000"

func main() {
	metrics.InitApiPrometheusMetrics()
	prometheusMetricServer := monitoring.StartPrometheusMetricServerOrPanic(listenAddr)
	defer monitoring.ShutdownPrometheusMetricServerOrPanic(prometheusMetricServer)

	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		metrics.ApiTotalRequests.WithLabelValues("counter").Inc()
		at := api.NewApiTimer()
		defer func() { at.StopWithMethod("timer") }()
		t := rand.Intn(1000) + 1
		time.Sleep(time.Duration(t) * time.Millisecond)
		c.JSON(http.StatusOK, gin.H{
			"message": "count up",
		})
	})
	r.Run()
}
