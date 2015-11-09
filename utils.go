package godis

import (
	"math/rand"
	"time"
)

func (g *Godis) getSDS(key string) (*SDS, bool) {
	s, exists := g.db[key]
	return s, exists
}

// generateRandnum returns a random number within given range.
func (g *Godis) generateRandnum(n int) int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(n)
}

// Destroy a key after given time in seconds
func (g *Godis) destroyInSecs(key string, exp uint64) (int, error) {
	time.Sleep(time.Duration(exp) * time.Second)
	return g.DEL(key)
}

// Destroy a key after given time in milliseconds
func (g *Godis) destroyInMillis(key string, exp uint64) (int, error) {
	time.Sleep(time.Duration(exp) * time.Millisecond)
	return g.DEL(key)
}

// Get all the keys in the entire database
// TODO : Lock the whole db while getting its keys for consistant results?
func (g *Godis) getAllKeys() []string {
	db := g.db
	keys := make([]string, len(db))
	if len(db) > 0 {
		i := 0
		for k := range db {
			keys[i] = k //not using append as we know exact length of the db
			i++
		}
	}
	return keys
}
