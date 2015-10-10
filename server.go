package godis

import (
//"errors"
)

//DBSIZE returns the number of keys of type int64 in the currently selected database.
func (g *Godis) DBSIZE() (int64, error) {
	return int64(len(g.db)), nil
}
