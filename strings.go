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
func (g Godis) INCR(key string, n int) interface{} {
	if g.Exists(key) == 1 {
		val := g.db[key].(int)
		g.db[key] = val + n
		return g.Get(key)
	} else {
		g.Set(key, n)
		return g.Get(key)
	}
}

// DECR decrements the key by one
func (g Godis) DECR(key string, n int) interface{} {
	if g.Exists(key) == 1 {
		val := g.db[key].(int)
		g.db[key] = val - n
		return g.Get(key)
	} else {
		g.Set(key, n)
		return g.Get(key)
	}
}

// INCRBY increments the key by given value
func (g Godis) INCRBY(key string, n int) interface{} {
	if g.Exists(key) == 1 {
		val := g.db[key].(int)
		g.db[key] = val + n
		return g.Get(key)
	} else {
		g.Set(key, n)
		return g.Get(key)
	}
}

// DECRBY decrements the key by given value
func (g Godis) DECRBY(key string, n int) interface{} {
	if g.Exists(key) == 1 {
		val := g.db[key].(int)
		g.db[key] = val - n
		return g.Get(key)
	} else {
		g.Set(key, n)
		return g.Get(key)
	}
}
