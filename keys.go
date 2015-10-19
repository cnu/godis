package godis

import (
	"errors"
)

// EXISTS returns in a key exists or not in the DB
func (g *Godis) EXISTS(keys ...string) int {
	count := 0
	for _, key := range keys {
		if _, ok := g.db[key]; ok {
			count++
		}
	}
	return count
}

// DEL removes all keys if it exists and returns the number of keys removed
func (g *Godis) DEL(keys ...string) int {
	count := g.EXISTS(keys...)
	for _, key := range keys {
		delete(g.db, key)
	}
	return count
}

// RENAME renames a key to newkey. Returns an error when the key
// and newkey are the same, or when key does not exist. If new key
// already exists it is overwritten.
func (g *Godis) RENAME(key, newKey string) (string, error) {
	if key == newKey {
		return "", errors.New("samekeys")
	}
	if g.EXISTS(key) == 0 {
		return "", errors.New("keynotexists")
	}
	if g.EXISTS(newKey) > 0 {
		g.DEL(newKey)
	}
	val, _ := g.GET(key)
	g.DEL(key)
	got, _ := g.SET(newKey, val)
	return got, nil
}

// RENAMENX is used to rename key to newkey if newkey does not yet exist.
// Returns an error under the same conditions as RENAME.
func (g *Godis) RENAMENX(key, newKey string) (string, error) {
	if key == newKey {
		return "", errors.New("samekeys")
	}
	if g.EXISTS(key) == 0 {
		return "", errors.New("keynotexists")
	}
	if g.EXISTS(newKey) != 0 {
		return "", errors.New("nosuchkey")
	}
	val, _ := g.GET(key)
	g.DEL(key)
	got, _ := g.SET(newKey, val)
	return got, nil

// RANDOMKEY returns a random key from the currently selected database.
func (g *Godis) RANDOMKEY() (*SDS, error) {
	keys := make([]string, len(g.db))
	i := 0
	// TODO : Move below logic to KEYS command when its implemented.
	for k := range g.db {
		keys[i] = k
		i++
	}
	if len(keys) == 0 {
		return nil, errors.New("emptydb")
	}
	rnum := g.generateRandnum(len(keys))
	return g.db[keys[rnum]], nil
}
