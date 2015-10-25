package godis

import (
//"errors"
)

//DBSIZE returns the number of keys of type int64 in the currently selected database.
// TODO : DBSIZE scans the entire database, implement an efficient logic like
//maintaining separate db for statistics of keys.
func (g *Godis) DBSIZE() (int64, error) {
	return int64(len(g.db)), nil
}
