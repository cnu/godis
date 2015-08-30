package godis

// Set is used to assign a value to a key
func (g Godis) Set(key string, value interface{}) string {
	g.db[key] = value
	return key
}

// Get returns the value stored for a key
func (g Godis) Get(key string) interface{} {
	return g.db[key]
}
