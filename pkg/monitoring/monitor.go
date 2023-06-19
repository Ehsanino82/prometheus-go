package monitoring

import (
	"context"
	"net/http"
	"prometheus-go/pkg/logging"

	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func StartPrometheusMetricServerOrPanic(listenAddr string) *http.Server {
	prometheusServer := &http.Server{
		Addr:    listenAddr,
		Handler: promhttp.Handler(),
	}

	go listenAndServePrometheusMetrics(prometheusServer)

	return prometheusServer
}

func listenAndServePrometheusMetrics(server *http.Server) {
	if err := server.ListenAndServe(); err != nil {
		logging.PanicWithError("failed to start prometheus http probe listener", err)
	}
}

func ShutdownPrometheusMetricServerOrPanic(server *http.Server) {
	if err := server.Shutdown(context.Background()); err != nil {
		logging.PanicWithError("Failed to shutdown prometheus metric server", err)
	}
}
