package godis

import "testing"

func TestEXISTS(t *testing.T) {
	db := setUp()
	for _, c := range cases {
		db.SET(c.key, c.value)
	}

	// one existent key
	key := "key1"
	if got := db.EXISTS(key); got != 1 {
		t.Errorf("EXISTS(%q) == %v, want %d", key, got, 1)
	}

	// one non-existent key
	notEXISTSKey := "not-exists"
	if got := db.EXISTS(notEXISTSKey); got != 0 {
		t.Errorf("EXISTS(%q) == %v, want %v", notEXISTSKey, got, 0)
	}

	// all existent keys
	keys := []string{"key1", "test_num", "key 3"}
	if got := db.EXISTS(keys...); got != len(keys) {
		t.Errorf("EXISTS(%q) == %d, want %d", keys, got, len(keys))
	}

	// two existent keys and one non-existent keys
	keys = []string{"key1", "test_num", "foo"}
	if got := db.EXISTS(keys...); got != 2 {
		t.Errorf("EXISTS(%q) == %d, want %d", keys, got, 2)
	}

	// all non existent keys
	keys = []string{"foo", "bar", "baz"}
	if got := db.EXISTS(keys...); got != 0 {
		t.Errorf("EXISTS(%q) == %d, want %d", keys, got, 0)
	}

}

func TestDEL(t *testing.T) {
	db := setUp()
	for _, c := range cases {
		db.SET(c.key, c.value)
	}

	// DEL a key which exists
	key := "மொழி"
	got := db.DEL(key)
	if got != 1 {
		t.Errorf("DEL(%q) == %d, want %d", key, got, 1)
	}

	// DEL a key which doesn't exist
	key = "foo"
	got = db.DEL(key)
	if got != 0 {
		t.Errorf("DEL(%q) == %d, want %d", key, got, 0)
	}

	// DEL a list of keys which all exist
	removeKeys := []string{"key1", "test_num", "key 3"}
	got = db.DEL(removeKeys...)
	if got != len(removeKeys) {
		t.Errorf("DEL(%q) == %d, want %d", removeKeys, got, len(removeKeys))
	}

	// DEL a list of keys which has one non-existent key
	removeKeys = []string{"key2", "tested", "not-exists"}
	got = db.DEL(removeKeys...)
	if got != 2 {
		t.Errorf("DEL(%q) == %d, want %d", removeKeys, got, 2)
	}

	// DEL a list of keys which has all non-existent keys
	removeKeys = []string{"foo", "bar", "baz"}
	got = db.DEL(removeKeys...)
	if got != 0 {
		t.Errorf("DEL(%q) == %d, want %d", removeKeys, got, 0)
	}
}

// Test RENAME when key and newKey are same
func TestRENAMESameKeys(t *testing.T) {
	db := setUp()
	key := "myKey"
	newKey := "myKey"
	db.SET(key, "value")
	res := db.RENAME(key, newKey)
	if res != false {
		t.Errorf("RENAME(%s, %s) == %s, want %t", key, newKey, res, false)
	}
}

// Test RENAME when given key doesn't exist
func TestRENAMENonExistant(t *testing.T) {
	db := setUp()
	key := "myKey"
	newKey := "hisKey"
	res := db.RENAME(key, newKey)
	if res != false {
		t.Errorf("RENAME(%s, %s) == %s, want %t", key, newKey, res, false)
	}
}
