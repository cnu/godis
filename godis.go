package godis

// Godis is the actual in memory DB
type Godis struct {
	db map[string]interface{}
}

// New returns a reference to a new empty Godis object
func New() *Godis {
	return &Godis{make(map[string]interface{})}
}

// Set is used to assign a value to a key
func (g Godis) Set(key string, value interface{}) string {
	g.db[key] = value
	return key
}

// Get returns the value stored for a key
func (g Godis) Get(key string) interface{} {
	return g.db[key]
}

// Exists returns in a key exists or not in the DB
func (g Godis) Exists(key string) bool {
	if _, ok := g.db[key]; ok {
		return true
	}
	return false

}
