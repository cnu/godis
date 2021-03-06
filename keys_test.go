package godis

import (
	"reflect"
	"sort"
	"testing"
)

func TestEXISTS(t *testing.T) {
	db := setUp()
	for _, c := range cases {
		db.SET(c.key, c.value)
	}

	// one existent key
	key := "key1"
	if got, err := db.EXISTS(key); got != 1 {
		t.Errorf("EXISTS(%q) == %v, %v want %d, <nil>", key, got, err, 1)
	}

	// one non-existent key
	notEXISTSKey := "not-exists"
	if got, err := db.EXISTS(notEXISTSKey); got != 0 {
		t.Errorf("EXISTS(%q) == %v, %v want %v, <nil>", notEXISTSKey, got, err, 0)
	}

	// all existent keys
	keys := []string{"key1", "key2", "key 3"}
	if got, err := db.EXISTS(keys...); got != len(keys) {
		t.Errorf("EXISTS(%q) == %d, %v want %d, <nil>", keys, got, err, len(keys))
	}

	// two existent keys and one non-existent keys
	keys = []string{"key1", "key2", "foo"}
	if got, err := db.EXISTS(keys...); got != 2 {
		t.Errorf("EXISTS(%q) == %d, %v want %d, <nil>", keys, got, err, 2)
	}

	// all non existent keys
	keys = []string{"foo", "bar", "baz"}
	if got, err := db.EXISTS(keys...); got != 0 {
		t.Errorf("EXISTS(%q) == %d, %v want %d, <nil>", keys, got, err, 0)
	}

}

func TestDEL(t *testing.T) {
	db := setUp()
	for _, c := range cases {
		db.SET(c.key, c.value)
	}

	// DEL a key which exists
	key := "மொழி"
	got, err := db.DEL(key)
	if got != 1 {
		t.Errorf("DEL(%q) == %d, %v want %d, <nil>", key, got, err, 1)
	}

	// DEL a key which doesn't exist
	key = "foo"
	got, err = db.DEL(key)
	if got != 0 {
		t.Errorf("DEL(%q) == %d, %v want %d, <nil>", key, got, err, 0)
	}

	// DEL a list of keys which all exist
	removeKeys := []string{"key1", "key 3"}
	got, err = db.DEL(removeKeys...)
	if got != len(removeKeys) {
		t.Errorf("DEL(%q) == %d, %v want %d, <nil>", removeKeys, got, err, len(removeKeys))
	}

	// DEL a list of keys which has one non-existent key
	removeKeys = []string{"key2", "not-exists"}
	got, err = db.DEL(removeKeys...)
	if got != 1 {
		t.Errorf("DEL(%q) == %d, %v want %d, <nil>", removeKeys, got, err, 1)
	}

	// DEL a list of keys which has all non-existent keys
	removeKeys = []string{"foo", "bar", "baz"}
	got, err = db.DEL(removeKeys...)
	if got != 0 {
		t.Errorf("DEL(%q) == %d, %v, want %d, <nil>", removeKeys, got, err, 0)
	}
}

// Test RENAME with different, key and newKey
func TestRENAME(t *testing.T) {
	db := setUp()
	key := "myKey"
	newKey := "hisKey"
	db.SET(key, "value")
	db.RENAME(key, newKey)
	got, err := db.GET(newKey)
	if got != "value" || err != nil {
		t.Errorf("GET(%q) == %q, %v want %q, <nil>", newKey, got, err, "value")
	}
}

// Test RENAME when key and newKey are same
func TestRENAMESameKeys(t *testing.T) {
	db := setUp()
	key := "myKey"
	newKey := "myKey"
	db.SET(key, "value")
	got, err := db.RENAME(key, newKey)
	// TODO : Check whether the error is returned as simple string or enclosed by
	// error(samekeys)
	if err.Error() != "samekeys" {
		t.Errorf("RENAME(%q, %q) == %q, %v want %q, samekeys", key, newKey, got,
			err, got)
	}
}

// Test RENAME when given key doesn't exist
func TestRENAMENonExistant(t *testing.T) {
	db := setUp()
	key := "myKey"
	newKey := "hisKey"
	got, err := db.RENAME(key, newKey)
	// TODO : Check whether the error is returned as simple string or enclosed by
	// error(keynotexists)
	if err.Error() != "keynotexists" {
		t.Errorf("RENAME(%q, %q) == %q, %v want %q, keynotexists", key, newKey,
			got, err, got)
	}
}

// Test RENAME when newKey exists
func TestRENAMENewKeyExist(t *testing.T) {
	db := setUp()
	key := "myKey"
	newKey := "hisKey"
	db.SET(key, "value")
	db.SET(newKey, "somevalue")
	db.RENAME(key, newKey)
	got, err := db.GET(newKey)
	if got != "value" || err != nil {
		t.Errorf("GET(%q) == %q, %v want %q, <nil>", newKey, got, err, "value")
	}
}

// Test RENAMENX with different key and newKey
func TestRENAMENX(t *testing.T) {
	db := setUp()
	key := "myKey"
	newKey := "hisKey"
	db.SET(key, "value")
	db.RENAMENX(key, newKey)
	got, err := db.GET(newKey)
	if got != "value" || err != nil {
		t.Errorf("GET(%q) == %q, %v want %q, <nil>", newKey, got, err, "value")
	}
}

// Test RENAMENX when key and newKey are same
func TestRENAMENXSameKeys(t *testing.T) {
	db := setUp()
	key := "myKey"
	newKey := "myKey"
	db.SET(key, "value")
	got, err := db.RENAMENX(key, newKey)
	// TODO : Check whether the error is returned as simple string or enclosed by
	// error(samekeys)
	if err.Error() != "samekeys" {
		t.Errorf("RENAMENX(%q, %q) == %q, %v want %q, samekeys", key, newKey, got,
			err, got)
	}
}

// Test RENAMENX when given key doesn't exist
func TestRENAMENXNonExistant(t *testing.T) {
	db := setUp()
	key := "myKey"
	newKey := "hisKey"
	got, err := db.RENAMENX(key, newKey)
	// TODO : Check whether the error is returned as simple string or enclosed by
	// error(keynotexists)
	if err.Error() != "keynotexists" {
		t.Errorf("RENAMENX(%q, %q) == %q, %v want %q, keynotexists", key, newKey,
			got, err, got)
	}
}

// Test RENAMENX when newKey exists
func TestRENAMENXNewKeyExist(t *testing.T) {
	db := setUp()
	key := "myKey"
	newKey := "hisKey"
	db.SET(key, "value")
	db.SET(newKey, "somevalue")
	got, err := db.RENAMENX(key, newKey)
	if err.Error() != "newkeyexists" {
		t.Errorf("RENAMENX(%q, %q) == %q, %v want \"\", newkeyexists", key, newKey,
			got, err)
	}
}

// Test RANDOMKEY for existing db
func TestRANDOMKEY(t *testing.T) {
	db := setUp()
	db.MSET("key1", "val1", "key2", "val2", "key3", "val3", "key4", "val4",
		"key5", "val5", "key6", "val6", "key7", "val7", "key8", "val8")
	got, err := db.RANDOMKEY()
	if err != nil {
		t.Errorf("RANDOMKEY() == %q, %v want %q, <nil>", got, err, got)
	}
}

// Test RANDOMKEY for non-existant db
// TODO : RANDOMKEY should return "" instead of <nil>
// with emptydb error
func TestRANDOMKEYNonExistant(t *testing.T) {
	db := setUp()
	got, err := db.RANDOMKEY()
	if err.Error() != "emptydb" {
		t.Errorf("RANDOMKEY() == %q, %v want %q, emptydb", got, err, got)
	}
}

// Test KEYS for given patterns
func TestKEYS(t *testing.T) {
	db := setUp()
	db.MSET("hello", "1", "hallo", "2", "hullo", "3", "hyllo", "4", "hvllo", "5",
		"heeeello", "6")
	for _, p := range patterns {
		pattern := p.regex
		want := p.result
		got, err := db.KEYS(pattern)
		sort.Strings(want)
		sort.Strings(got)
		if !reflect.DeepEqual(got, want) || err != nil {
			t.Errorf("KEYS(%q) == %v, %v want %v, <nil>", pattern, got, err, want)
		}
	}
}

// Test KEYS for non-existant db,
// should return empty slice
func TestKEYSEmptyDB(t *testing.T) {
	db := setUp()
	pattern := "h?llo"
	got, err := db.KEYS(pattern)
	if got != nil || err != nil {
		t.Errorf("KEYS(%q) == %v, %v want %v, <nil>", pattern, got, err, got)
	}
}

// Test KEYS for invalid regex patterns,
// should return "invalidregex" error
func TestKEYSInvalidRegex(t *testing.T) {
	db := setUp()
	pattern := "h?++llo"
	got, err := db.KEYS(pattern)
	if got != nil || err.Error() != "invalidregex" {
		t.Errorf("KEYS(%q) == %v, %v want %v, invalidregex", pattern, got, err, got)
	}
}
