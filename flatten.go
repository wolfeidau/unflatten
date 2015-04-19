package unflatten

func Flatten(m map[string]interface{}, tokenizer func([]string) string) map[string]interface{} {
	var r = make(map[string]interface{})
	flattenrecurse(m, []string{}, func(ks []string, v interface{}) {
		r[tokenizer(ks)] = v
	})
	return r
}

func flattenrecurse(m map[string]interface{}, ks []string, cb func([]string, interface{})) {
	for k, v := range m {
		newks := append(ks, k)
		if newm, ok := v.(map[string]interface{}); ok {
			flattenrecurse(newm, newks, cb)
		} else {
			cb(newks, v)
		}
	}
}
