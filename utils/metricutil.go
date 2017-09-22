package utils

import (
	"djforgo/system"
	l4g "github.com/alecthomas/log4go"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"net/http"
)

// Push metircs in background.
func PrometheusMonitorStart() {
	if len(system.SysConfig.Metric.Addr) == 0 {
		l4g.Info("disable Prometheus client")
		return
	}

	l4g.Info("start Prometheus client")

	http.Handle("/metrics", promhttp.Handler())
	l4g.Critical(http.ListenAndServe(system.SysConfig.Metric.Addr, nil))
}
