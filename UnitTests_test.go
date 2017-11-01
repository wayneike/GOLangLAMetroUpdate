package MetroAPI

import (
	"bytes"
	"encoding/json"
	"sort"
	"testing"
)

func TestConstructInfo(t *testing.T) {
	var buf bytes.Buffer

	constructInfo(&buf, 12.0)

	if buf.String() != "12 minutes\n" {
		t.Errorf("ConstructInfo was incorrect. Returned: %s", buf.String())
	}
}

func TestReqHTTPGet(t *testing.T) {
	url := "http://api.metro.net/agencies/lametro-rail/stops/80212/predictions/"

	body := reqHTTPGet(url)
	if len(body) <= 0 {
		t.Errorf("No body was returned for %s", url)
	}
}

func TestSort(t *testing.T) {
	const jsonStr = `{
		"items": [
			{
				"seconds": 1055.0,
				"block_id": "214",
				"route_id": "802",
				"is_departing": true,
				"run_id": "802_1_var0",
				"minutes": 17.0
			},
			{
				"seconds": 79.0,
				"block_id": "214",
				"route_id": "802",
				"is_departing": true,
				"run_id": "802_0_var0",
				"minutes": 1.0
			},
			{
				"seconds": 1314.0,
				"block_id": "219",
				"route_id": "802",
				"is_departing": true,
				"run_id": "802_0_var0",
				"minutes": 21.0
			},
			{
				"seconds": 413.0,
				"block_id": "202",
				"route_id": "805",
				"is_departing": true,
				"run_id": "805_1_var0",
				"minutes": 6.0
			}
		]
	}`

	jsonByte := []byte(jsonStr)

	pred := items{}

	jsonErr := json.Unmarshal(jsonByte, &pred)
	if jsonErr != nil {
		t.Errorf("%s", jsonErr)
	}

	sort.Sort(byMinutes(pred.Items))

	if pred.Items[1].Minutes != 6.0 {
		t.Errorf("The second element should have 6 minutes")
	}

	if pred.Items[2].Minutes != 17.0 {
		t.Errorf("The third element should have 17 minutes")
	}
}
