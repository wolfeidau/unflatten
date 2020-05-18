package unflatten

import (
	"bytes"
	"encoding/json"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

var m = map[string]interface{}{
	"cpu.usage.0.user": map[string]interface{}{
		"value": 2.3,
	},
	"cpu.usage.0.system": map[string]interface{}{
		"value": 1.2,
	},
	"cpu.usage.0.total": map[string]interface{}{
		"value": 1.1,
	},
	"cpu.usage.total": map[string]interface{}{
		"value": 2.3,
	},
	"cpu.usage.system": map[string]interface{}{
		"value": 1.2,
	},
	"cpu.usage.user": map[string]interface{}{
		"value": 1.1,
	},
	"memory.user": map[string]interface{}{
		"value": 1.1,
	},
	"memory.system": map[string]interface{}{
		"value": 2.1,
	},
}

var jsonResult = `{"cpu":{"usage":{"0":{"system":{"value":1.2},"total":{"value":1.1},"user":{"value":2.3}},"system":{"value":1.2},"total":{"value":2.3},"user":{"value":1.1}}},"memory":{"system":{"value":2.1},"user":{"value":1.1}}}`

func TestUnflatten(t *testing.T) {

	tree := Unflatten(m, func(k string) []string { return strings.Split(k, ".") })

	payload, _ := json.Marshal(tree)

	if string(payload) != jsonResult {
		t.Errorf("expected %s got %s", jsonResult, string(payload))
	}

}

var fromJSON = `{
  "Connections.Accepted": 4,
  "Connections.Open": 2,
  "Memory.Alloc": 682208,
  "Memory.Frees": 2567,
  "Memory.Lookups": 281,
  "Memory.Mallocs": 3326,
  "Memory.Sys": 5441784,
  "Memory.TotalAlloc": 1032488,
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

var expectedJSON = `{"Connections":{"Accepted":4,"Open":2},"Memory":{"Alloc":682208,"Frees":2567,"Lookups":281,"Mallocs":3326,"Sys":5441784,"TotalAlloc":1032488},"Peers":{"IPv6":{"Completed":0,"Current":0,"Joined":0,"Left":0,"Reaped":0,"Seeds":{"Current":0,"Joined":0,"Left":0,"Reaped":0}}},"ResponseTime":{"P50":0.045775,"P90":0.074299,"P95":0.096207}}`

func TestUnflattenConfig(t *testing.T) {
	assert := require.New(t)

	b := new(bytes.Buffer)
	config := make(map[string]interface{})

	b.WriteString(fromJSON)

	err := json.Unmarshal(b.Bytes(), &config)
	assert.NoError(err)

	tree := Unflatten(config, func(k string) []string { return strings.Split(k, ".") })

	payload, _ := json.Marshal(tree)

	if string(payload) != expectedJSON {
		t.Errorf("expected %s got %s", expectedJSON, string(payload))
	}

}

func BenchmarkNew(b *testing.B) {
	assert := require.New(b)
	buf := new(bytes.Buffer)
	config := make(map[string]interface{})

	buf.WriteString(fromJSON)

	err := json.Unmarshal(buf.Bytes(), &config)
	assert.NoError(err)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Unflatten(config, SplitByDot)
	}
}
