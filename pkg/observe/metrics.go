package observe

import (
	"github.com/gorilla/mux"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

var (
	jobStatusMetric = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Name: "aleph_job_status",
			Help: "The number index/ingest jobs ",
		},
		[]string{
			"job_name",
			"stage",
			"status",
		},
	)
)

// RegisterPrometheus adds the prometheus handler to the mux router
// Note you must register every metric with prometheus for it show up
// when the /metrics route is hit.
func RegisterPrometheus(m *mux.Router) *mux.Router {
	prometheus.MustRegister(jobStatusMetric)

	m.Handle("/metrics", promhttp.Handler())
	return m
}

func UpdatePrometheus(status AlephCollectionStatus) {
	jobStatusMetric.WithLabelValues(status.Collection.Label, "finished").Set(float64(status.Finished))
	jobStatusMetric.WithLabelValues(status.Collection.Label, "pending").Set(float64(status.Pending))
	jobStatusMetric.WithLabelValues(status.Collection.Label, "running").Set(float64(status.Running))
}
