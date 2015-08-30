package godis

import "testing"

func setUp() *Godis {
	return New()
}

type KV struct {
	key   string
	value interface{}
}

var cases = []KV{
	KV{"key1", "value1"},
	KV{"key2", "value2"},
	KV{"key 3", "value 3"},    // keys with spaces
	KV{"மொழி", "தமிழ்"},       // unicode
	KV{"key1", "new value 1"}, // overwrite a key
	KV{"tested", true},        // boolean value
	KV{"test_num", 7},         // int value
	KV{"PI", 3.14},            // float value
}

// Test setting key-values to the DB
func TestSet(t *testing.T) {

	db := setUp()
	for _, c := range cases {
		got := db.Set(c.key, c.value)
		if got != c.key {
			t.Errorf("Set(%q) == %q, want %q", c.key, got, c.key)
		}
	}
}

// Test getting key-values from the DB
func TestGet(t *testing.T) {
	db := setUp()
	for _, c := range cases {
		db.Set(c.key, c.value)
		got := db.Get(c.key)
		if got != c.value {
			t.Errorf("Get(%q) == %q, want %q", c.key, got, c.value)
		}
	}
}

// Test getting non-existent key from the DB
func TestGetNotExists(t *testing.T) {
	db := setUp()
	key := "not exists"
	got := db.Get(key)
	if got != nil {
		t.Errorf("Get(%q) == %v, want %v", key, got, nil)
	}
}

func TestExists(t *testing.T) {
	db := setUp()
	key := "exists"
	db.Set(key, "yes")
	if got := db.Exists(key); !got {
		t.Errorf("Exists(%q) == %v, want %v", key, got, true)
	}

	notExistsKey := "not-exists"
	if got := db.Exists(notExistsKey); got {
		t.Errorf("Exists(%q) == %v, want %v", notExistsKey, got, false)
	}
}

func TestDelete(t *testing.T) {
	db := setUp()
	for _, c := range cases {
		db.Set(c.key, c.value)
	}

	// Del a key which exists
	key := "மொழி"
	got := db.Del(key)
	if got != 1 {
		t.Errorf("Del(%q) == %d, want %d", key, got, 1)
	}

	// Del a key which doesn't exist
	key = "foo"
	got = db.Del(key)
	if got != 0 {
		t.Errorf("Del(%q) == %d, want %d", key, got, 0)
	}

	// Del a list of keys which all exist
	removeKeys := []string{"key1", "test_num", "key 3"}
	got = db.Del(removeKeys...)
	if got != len(removeKeys) {
		t.Errorf("Del(%q) == %d, want %d", removeKeys, got, len(removeKeys))
	}

	// Del a list of keys which has one non-existent key
	removeKeys = []string{"key2", "tested", "not-exists"}
	got = db.Del(removeKeys...)
	if got != 2 {
		t.Errorf("Del(%q) == %d, want %d", removeKeys, got, 2)
	}

	// Del a list of keys which has all non-existent keys
	removeKeys = []string{"foo", "bar", "baz"}
	got = db.Del(removeKeys...)
	if got != 0 {
		t.Errorf("Del(%q) == %d, want %d", removeKeys, got, 0)
	}

}
