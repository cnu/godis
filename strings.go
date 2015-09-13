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

// INCR increments the key by one
func (g Godis) INCR(key string) interface{} {
	return g.INCRBY(key, 1)
}

// DECR decrements the key by one
func (g Godis) DECR(key string) interface{} {
	return g.DECRBY(key, 1)
}

// INCRBY increments the key by given value
func (g Godis) INCRBY(key string, n int) interface{} {
	if g.EXISTS(key) == 0 {
		g.SET(key, 0)
	}
	val := g.GET(key).(int)
	g.SET(key, val+n)
	return g.GET(key)
}

// DECRBY decrements the key by given value
func (g Godis) DECRBY(key string, n int) interface{} {
	if g.EXISTS(key) == 0 {
		g.SET(key, 0)
	}
	val := g.GET(key).(int)
	g.SET(key, val-n)
	return g.GET(key)
}
