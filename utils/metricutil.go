package utils

import (
	"djforgo/config"
	l4g "github.com/alecthomas/log4go"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"net/http"
)

// Push metircs in background.
func MetricStart() {
	if len(config.QasConfig.Metric.Addr) == 0 {
		l4g.Info("disable Prometheus client")
		return
	}

	l4g.Info("start Prometheus client")

	http.Handle("/metrics", promhttp.Handler())
	l4g.Critical(http.ListenAndServe(config.QasConfig.Metric.Addr, nil))
}
