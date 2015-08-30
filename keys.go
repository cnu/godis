package godis

// Exists returns in a key exists or not in the DB
func (g Godis) Exists(keys ...string) int {
	count := 0
	for _, key := range keys {
		if _, ok := g.db[key]; ok {
			count++
		}
	}
	return count
}

// Del removes all keys if it exists and returns the number of keys removed
func (g Godis) Del(keys ...string) int {
	count := g.Exists(keys...)
	for _, key := range keys {
		delete(g.db, key)
	}
	return count
}
