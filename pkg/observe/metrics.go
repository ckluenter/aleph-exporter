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

func UpdatePrometheus(status AlephStatus) {
	for _,job := range status.jobs {
		for _,stage := range job.stages {
			jobStatusMetric.WithLabelValues(job.collection.label,stage.stage,"finished").Set(stage.finished)
			jobStatusMetric.WithLabelValues(job.collection.label,stage.stage,"running").Set(stage.finished)
			jobStatusMetric.WithLabelValues(job.collection.label,stage.stage,"pending").Set(stage.finished)
		}
	}
}
