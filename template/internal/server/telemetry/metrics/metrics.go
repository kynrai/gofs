package metrics

import (
	"net/http"

	"module/placeholder/internal/config"

	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func Expose(r *http.ServeMux) *http.ServeMux {
	conf := config.New()
	if conf.Metrics {
		r.Handle("GET /metrics", promhttp.Handler())
	}
	return r
}
