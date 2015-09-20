package godis

func (g *Godis) getSDS(key string) (*SDS, bool) {
	s, exists := g.db[key]
	return s, exists
}

// SET is used to assign a value to a key
func (g *Godis) SET(key string, value string) string {
	if g.EXISTS(key) == 1 {
		s, exists := g.getSDS(key)
		if exists {
			s.Lock()
			defer s.Unlock()
			s.value = value
		}
	} else {
		s := NewSDS(value)
		g.db[key] = s
	}
	return key
}

// GET returns the value stored for a key
func (g *Godis) GET(key string) (string, bool) {
	s, exists := g.getSDS(key)
	if exists {
		return s.get(), false
	}
	return "", true
}

// // INCR increments the key by one
// func (g Godis) INCR(key string) interface{} {
// 	return g.INCRBY(key, 1)
// }

// // DECR decrements the key by one
// func (g Godis) DECR(key string) interface{} {
// 	return g.DECRBY(key, 1)
// }

// // INCRBY increments the key by given value
// func (g Godis) INCRBY(key string, n int) interface{} {
// 	if g.EXISTS(key) == 0 {
// 		g.SET(key, 0)
// 	}
// 	val := g.GET(key).(int)
// 	g.SET(key, val+n)
// 	return g.GET(key)
// }

// // DECRBY decrements the key by given value
// func (g Godis) DECRBY(key string, n int) interface{} {
// 	if g.EXISTS(key) == 0 {
// 		g.SET(key, 0)
// 	}
// 	val := g.GET(key).(int)
// 	g.SET(key, val-n)
// 	return g.GET(key)
// }

// MGET returns a slice of values for a input slice of keys
func (g *Godis) MGET(keys ...string) []interface{} {
	var output []interface{} // will be strings or nils
	for _, key := range keys {
		value, err := g.GET(key)
		if !err {
			output = append(output, value)
		} else {
			// if the key isn't available, append a nil instead
			// TODO: how to do multi-return way if we have a slice to return?
			output = append(output, nil)
		}
	}
	return output
}

// MSET sets a slice of key-values
// pass a slice of keys and value alternating
// eg: "key1", "value1", "key2", "value2"
func (g *Godis) MSET(items ...string) bool {
	g.Lock()
	defer g.Unlock()
	for i, item := range items {
		if i%2 == 1 {
			continue
		}
		g.SET(item, items[i+1])
	}
	return true
}
