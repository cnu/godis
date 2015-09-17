package godis

// EXISTS returns in a key exists or not in the DB
func (g Godis) EXISTS(keys ...string) int {
	count := 0
	for _, key := range keys {
		if _, ok := g.db[key]; ok {
			count++
		}
	}
	return count
}

// DEL removes all keys if it exists and returns the number of keys removed
func (g Godis) DEL(keys ...string) int {
	count := g.EXISTS(keys...)
	for _, key := range keys {
		delete(g.db, key)
	}
	return count
}

// RENAME renames a key to newkey. Returns an error when the source
// and destination names are the same, or when key does not exist. If new key
// already exists it is overwritten.
func (g Godis) RENAME(key, newKey string) interface{} {
	if key != newKey || g.EXISTS(key) != 0 {
		if g.EXISTS(newKey) > 0 {
			g.DEL(newKey)
		}
		val := g.GET(key)
		g.DEL(key)
		return g.SET(newKey, val)
	} else {
		return false
	}
}
