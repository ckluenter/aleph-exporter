package observe

import (
	"github.com/parnurzeal/gorequest"
)

type AlephStage struct {
	job_id string
	stage string
	finished float64
	running float64
	pending float64

}
type AlephJob struct {
	finished float64
	running float64
	pending float64
	stages []AlephStage
	collection AlephCollection
}
type AlephCollection struct {
	created_at string
	updated_at string
	kind string
	collection_id string
	label string
}
type AlephStatus struct {
	running float64
	finished float64
	pending float64
	jobs []AlephJob
}

func GetAlephStatus(host string, token string) string {
	request:=gorequest.New()
	_, body, _ := request.Get(host).
		Set("Authorization", "ApiKey " + token).
		End()
	return body
}
