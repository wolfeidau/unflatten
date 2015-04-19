// +build old

package unflatten

import (
	"bytes"
	"encoding/json"
	"strings"
	"testing"
)

func BenchmarkOld(b *testing.B) {
	buf := new(bytes.Buffer)
	config := make(map[string]interface{})

	buf.WriteString(fromJSON)

	json.Unmarshal(buf.Bytes(), &config)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		OldUnflatten(config, func(k string) []string { return strings.Split(k, ".") })
	}
}
