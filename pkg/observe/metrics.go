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
			"collection_name",
			"job_id",
			"stage_task",
			"stage",
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
	for _,collection := range status.Collections {
		for _,job := range collection.Jobs {
			for _,stage := range job.Stages {
				jobStatusMetric.WithLabelValues(collection.Collection.Label, stage.Job_id, stage.Stage, "running").Set(float64(stage.Running))
				jobStatusMetric.WithLabelValues(collection.Collection.Label, stage.Job_id, stage.Stage, "pending").Set(float64(stage.Pending))
				jobStatusMetric.WithLabelValues(collection.Collection.Label, stage.Job_id, stage.Stage, "finished").Set(float64(stage.Finished))
			}
		}
	}
}
