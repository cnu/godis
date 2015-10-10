package godis

import (
	"math/rand"
	"time"
)

// generateRandnum returns a random number within given range.
func (g *Godis) generateRandnum(n int) int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(n)
}

// Destroy a key after given time in seconds
func (g *Godis) destroyInSecs(key string, exp uint64) int {
	time.Sleep(time.Duration(exp) * time.Second)
	return g.DEL(key)
}

// Destroy a key after given time in milliseconds
func (g *Godis) destroyInMillis(key string, exp uint64) int {
	time.Sleep(time.Duration(exp) * time.Millisecond)
	return g.DEL(key)
}
