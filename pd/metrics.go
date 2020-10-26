package pd

import "github.com/prometheus/client_golang/prometheus"

var (
	namespace   = "client_go_pd_client"
	cmdDuration = prometheus.NewHistogramVec(
		prometheus.HistogramOpts{
			Namespace: namespace,
			Subsystem: "cmd",
			Name:      "handle_cmds_duration_seconds",
			Help:      "Bucketed histogram of processing time (s) of handled success cmds.",
			Buckets:   prometheus.ExponentialBuckets(0.0005, 2, 13),
		}, []string{"type"})

	cmdFailedDuration = prometheus.NewHistogramVec(
		prometheus.HistogramOpts{
			Namespace: namespace,
			Subsystem: "cmd",
			Name:      "handle_failed_cmds_duration_seconds",
			Help:      "Bucketed histogram of processing time (s) of failed handled cmds.",
			Buckets:   prometheus.ExponentialBuckets(0.0005, 2, 13),
		}, []string{"type"})

	requestDuration = prometheus.NewHistogramVec(
		prometheus.HistogramOpts{
			Namespace: namespace,
			Subsystem: "request",
			Name:      "handle_requests_duration_seconds",
			Help:      "Bucketed histogram of processing time (s) of handled requests.",
			Buckets:   prometheus.ExponentialBuckets(0.0005, 2, 13),
		}, []string{"type"})
)

func init() {
	prometheus.MustRegister(cmdDuration)
	prometheus.MustRegister(cmdFailedDuration)
	prometheus.MustRegister(requestDuration)
}
