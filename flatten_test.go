package unflatten

import (
	"encoding/json"
	"testing"
)

func TestFlattenConfig(t *testing.T) {
	var expectedJSON = `{
"Connections.Accepted": 4,
"Connections.Open": 2,
"Memory.Alloc": 682208,
"Memory.Frees": 2567,
"Memory.Lookups": 281,
"Memory.Mallocs": 3326,
"Memory.Sys": 54417,
"Memory.TotalAlloc": 10324,
"Peers.IPv6.Completed": 0,
"Peers.IPv6.Current": 0,
"Peers.IPv6.Joined": 0,
"Peers.IPv6.Left": 0,
"Peers.IPv6.Reaped": 0,
"Peers.IPv6.Seeds.Current": 0,
"Peers.IPv6.Seeds.Joined": 0,
"Peers.IPv6.Seeds.Left": 0,
"Peers.IPv6.Seeds.Reaped": 0,
"ResponseTime.P50": 0.045775,
"ResponseTime.P90": 0.074299,
"ResponseTime.P95": 0.096207
}`
	var fromJSON = `{"Connections":{"Accepted":4,"Open":2},"Memory":{"Alloc":682208,"Frees":2567,"Lookups":281,"Mallocs":3326,"Sys":54417,"TotalAlloc":10324},"Peers":{"IPv6":{"Completed":0,"Current":0,"Joined":0,"Left":0,"Reaped":0,"Seeds":{"Current":0,"Joined":0,"Left":0,"Reaped":0}}},"ResponseTime":{"P50":0.045775,"P90":0.074299,"P95":0.096207}}`

	config := make(map[string]interface{})
	json.Unmarshal([]byte(fromJSON), &config)

	tree := Flatten(config, JoinWithDot)
	payload, _ := json.MarshalIndent(tree, "", "")
	if string(payload) != expectedJSON {
		t.Errorf("expected %s got %s", expectedJSON, string(payload))
	}
}
