package observe

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestAlephParsing(t *testing.T) {
	data := []byte(`
{
  "results": [
    {
      "finished": 2
    },
    {
      "finished": 6196,
      "running": 174,
      "pending": 0,
      "jobs": [
        {
          "finished": 6196,
          "running": 174,
          "pending": 0,
          "stages": [
            {
              "job_id": "9:76c04d8a-321d-49b7-95a9-245c5522e559",
              "stage": "index",
              "pending": 0,
              "running": 1,
              "finished": 3099
            },
            {
              "job_id": "9:76c04d8a-321d-49b7-95a9-245c5522e559",
              "stage": "ingest",
              "pending": 0,
              "running": 173,
              "finished": 3097
            }
          ],
          "start_time": "2019-12-20T08:05:16.179229",
          "end_time": null
        }
      ],
      "collection": {
        "created_at": "2019-12-20T08:05:14.474451",
        "updated_at": "2019-12-20T09:11:35.005722",
        "category": "other",
        "kind": "source",
        "id": "50",
        "collection_id": "50",
        "foreign_id": "some_collection_foreign_id",
        "label": "some_collection_label",
        "casefile": false,
        "secret": true,
        "writeable": true
      },
      "id": "some_collection_id"
    }
  ],
  "total": 2
}
	`)
	parsedData := ParseAlephStatus(data)
	assert.Equal(t, 2, parsedData.Total)
	assert.Equal(t, 2, len(parsedData.Collections))
	assert.Equal(t, 6196, parsedData.Collections[1].Finished)
	assert.Equal(t, "some_collection_label", parsedData.Collections[1].Collection.Label)
	assert.Equal(t, 174, parsedData.Collections[1].Jobs[0].Running)
	assert.Equal(t, "index", parsedData.Collections[1].Jobs[0].Stages[0].Stage)
}
