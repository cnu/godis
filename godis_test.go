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
	for _, c := range cases {
		db.Set(c.key, c.value)
	}

	// one existent key
	key := "key1"
	if got := db.Exists(key); got != 1 {
		t.Errorf("Exists(%q) == %v, want %d", key, got, 1)
	}

	// one non-existent key
	notExistsKey := "not-exists"
	if got := db.Exists(notExistsKey); got != 0 {
		t.Errorf("Exists(%q) == %v, want %v", notExistsKey, got, 0)
	}

	// all existent keys
	keys := []string{"key1", "test_num", "key 3"}
	if got := db.Exists(keys...); got != len(keys) {
		t.Errorf("Exists(%q) == %d, want %d", keys, got, len(keys))
	}

	// two existent keys and one non-existent keys
	keys = []string{"key1", "test_num", "foo"}
	if got := db.Exists(keys...); got != 2 {
		t.Errorf("Exists(%q) == %d, want %d", keys, got, 2)
	}

	// all non existent keys
	keys = []string{"foo", "bar", "baz"}
	if got := db.Exists(keys...); got != 0 {
		t.Errorf("Exists(%q) == %d, want %d", keys, got, 0)
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

// Benchmark for Set method
func BenchmarkSet(b *testing.B) {
	db := setUp()
	b.ResetTimer()

	// run the Set method b.N times
	for n := 0; n < b.N; n++ {
		db.Set("key", "value")
	}
}

// Benchmark for Get method
func BenchmarkGet(b *testing.B) {
	db := setUp()
	db.Set("key", "value")
	b.ResetTimer()

	// run the Get method b.N times
	for n := 0; n < b.N; n++ {
		_ = db.Get("key")
	}
}

// Benchmark for Exists method
func BenchmarkExists(b *testing.B) {
	db := setUp()
	db.Set("key", "value")
	b.ResetTimer()

	// run the Exists method b.N times
	for n := 0; n < b.N; n++ {
		_ = db.Exists("key")
	}
}

// Benchmark for Del method
func BenchmarkDel(b *testing.B) {
	db := setUp()
	b.ResetTimer()

	// Set a key and Del method b.N times
	for n := 0; n < b.N; n++ {
		// Not an accurate measure as we can't Del a key without setting it
		db.Set("key", "value")
		_ = db.Del("key")
	}
}
