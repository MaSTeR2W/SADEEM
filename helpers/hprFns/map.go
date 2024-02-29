package hprFns

func GetExistedKeysVals(m map[string]any, keys ...string) ([]string, []any) {
	var ks = []string{}
	var vs = []any{}
	for _, key := range keys {
		if val, ok := m[key]; ok {
			ks = append(ks, key)
			vs = append(vs, val)
		}
	}
	return ks, vs
}
