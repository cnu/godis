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
