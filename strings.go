package godis

// SET is used to assign a value to a key
func (g Godis) SET(key string, value interface{}) string {
	g.db[key] = value
	return key
}

// GET returns the value stored for a key
func (g Godis) GET(key string) interface{} {
	return g.db[key]
}

// MGET returns a slice of values for a input slice of keys
func (g Godis) MGET(keys ...string) []interface{} {
	var output []interface{}
	for _, key := range keys {
		value := g.GET(key)
		output = append(output, value)
	}
	return output
}

// MSET sets a slice of key-values
// pass a slice of keys and value alternating
// eg: "key1", "value1", "key2", "value2"
func (g Godis) MSET(items ...string) bool {
	for i, item := range items {
		if i%2 == 1 {
			continue
		}
		g.SET(item, items[i+1])
	}
	return true
}
