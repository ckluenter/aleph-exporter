package observe

import (
	"encoding/json"
	"fmt"
	"github.com/parnurzeal/gorequest"
)

type AlephStage struct {
	Job_id   string
	Stage    string
	Finished float64
	Running  float64
	Pending  float64
}
type AlephJob struct {
	Finished int
	Running  int
	Pending  int
	Stages   []AlephStage
}
type AlephCollection struct {
	CreatedAt    string `"json:created_at"`
	UpdatedAt    string `json:"updated_at"`
	Kind         string `json:"kind"`
	CollectionId string `json:"collection_id"`
	Label        string `json:"label"`
}
type AlephCollectionStatus struct {
	Running    int             `json:"running"`
	Finished   int             `json:"finished"`
	Pending    int             `json:"pending"`
	Jobs       []AlephJob      `json:"jobs"`
	Collection AlephCollection `json:"collection"`
}

type AlephStatus struct {
	Collections []AlephCollectionStatus `json:"results"`
	Total       int                     `json:"total"`
}

func GetAlephStatus(host string, token string) string {
	request := gorequest.New()
	_, body, _ := request.Get(host).
		Set("Authorization", "ApiKey "+token).
		End()
	return body
}

func ParseAlephStatus(body []byte) AlephStatus {
	var status = AlephStatus{}
	err := json.Unmarshal(body, &status)
	if err != nil {
		fmt.Printf("Error: %v",err)
	}
	return status
}
