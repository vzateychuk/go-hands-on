package myhttp

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"time"
)

var (
	uptimeMetric = promauto.NewCounter(prometheus.CounterOpts{
		Name: "myapp_uptime_seconds_total",
		Help: "The total uptime",
	})

	invokeCount = promauto.NewCounter(prometheus.CounterOpts{
		Name: "myapp_invocation_total",
		Help: "The total invocation amount",
	})

	httpReqs = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "http_requests_total",
			Help: "How many HTTP requests processed, partitioned by status code and HTTP method.",
		},
		[]string{"code", "method"},
	)
)

func recordUptime() {
	go func() {
		for {
			uptimeMetric.Inc()
			time.Sleep(1 * time.Second)
		}
	}()
}
