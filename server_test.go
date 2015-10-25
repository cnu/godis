package godis

import (
	"testing"
)

// DBSIZE should return the count of keys present in current DB
func TestDBSIZE(t *testing.T) {
	db := setUp()
	db.MSET("mykey0", "0", "mykey1", "1", "mykey2", "2", "mykey3", "3", "mykey4", "4", "mykey5",
		"5", "mykey6", "6", "mykey7", "7", "mykey8", "8", "mykey9", "9")
	got, err := db.DBSIZE()
	if err != nil || got != 10 {
		t.Errorf("DBSIZE() == %v, %v want 10, <nil>", got, err)
	}
}
