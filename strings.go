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

