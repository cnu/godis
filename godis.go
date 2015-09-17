package godis

// Godis is the actual in memory DB
type Godis struct {
	db map[string]*SDS
}

// NewGodis returns a reference to a new empty Godis object
func NewGodis() *Godis {
	return &Godis{make(map[string]*SDS)}
}
