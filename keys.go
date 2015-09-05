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
