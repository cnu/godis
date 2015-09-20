package godis

import "sync"

// Godis is the actual in memory DB
type Godis struct {
	db           map[string]*SDS
	sync.RWMutex // DB level lock
}

// NewGodis returns a reference to a new empty Godis object
func NewGodis() *Godis {
	return &Godis{db: make(map[string]*SDS)}
}
