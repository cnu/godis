package godis

import "time"

// SET is used to assign a value to a key
func (g Godis) SET(key string, value interface{}) string {
	g.db[key] = value
	return key
}

// GET returns the value stored for a key
func (g Godis) GET(key string) interface{} {
	return g.db[key]
}

// Internal function to destroy a key after given time in seconds
func (g Godis) destroyInSecs(key string, exp int64) int {
	var i int64 = 0
	for i = 0; i < exp; i++ {
		time.Sleep(1000 * time.Millisecond)
	}
	return g.DEL(key)
}

/* SETEX is used to assign a value to a key and destroy it within its given
expiry time in seconds*/
func (g Godis) SETEX(key string, exp int64, value interface{}) string {
	g.db[key] = value
	go g.destroyKey(key, exp)
	return key
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
