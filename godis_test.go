package godis

import "testing"

func setUp() *Godis {
	return New()
}

type Case struct {
	key   string
	value interface{}
}

var cases = []Case{
	{"key1", "value1"},
	{"key2", "value2"},
	{"key 3", "value 3"},    // keys with spaces
	{"மொழி", "தமிழ்"},       // unicode
	{"key1", "new value 1"}, // overwrite a key
	{"tested", true},        // boolean value
	{"test_num", 7},         // int value
	{"PI", 3.14},            // float value
}

// Test setting key-values to the DB
func TestSET(t *testing.T) {

	db := setUp()
	for _, c := range cases {
		got := db.SET(c.key, c.value)
		if got != c.key {
			t.Errorf("SET(%q) == %q, want %q", c.key, got, c.key)
		}
	}
}

// Test getting key-values from the DB
func TestGET(t *testing.T) {
	db := setUp()
	for _, c := range cases {
		db.SET(c.key, c.value)
		got := db.GET(c.key)
		if got != c.value {
			t.Errorf("GET(%q) == %q, want %q", c.key, got, c.value)
		}
	}
}

// Test getting non-existent key from the DB
func TestGETNotEXISTS(t *testing.T) {
	db := setUp()
	key := "not exists"
	got := db.GET(key)
	if got != nil {
		t.Errorf("GET(%q) == %v, want %v", key, got, nil)
	}
}

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
