package godis

// Godis is the actual in memory DB
type Godis struct {
	db map[string]interface{}
}

// New returns a reference to a new empty Godis object
func New() *Godis {
	return &Godis{make(map[string]interface{})}
}
