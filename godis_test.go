package godis

import "testing"

func setUp() *Godis {
	return New()
}

type KV struct {
	key   string
	value interface{}
}

// Test setting key-values to the DB
func TestSet(t *testing.T) {
	cases := []struct {
		in  KV
		out interface{}
	}{
		{KV{"key1", "value1"}, "key1"},
		{KV{"key2", "value2"}, "key2"},
		{KV{"key 3", "value 3"}, "key 3"},   // keys with spaces
		{KV{"மொழி", "தமிழ்"}, "மொழி"},       // unicode
		{KV{"key1", "new value 1"}, "key1"}, // overwrite a key
		{KV{"tested", true}, "tested"},      // boolean value
		{KV{"test_num", 7}, "test_num"},     // int value
		{KV{"PI", 3.14}, "PI"},              // float value
	}
	db := setUp()
	for _, c := range cases {
		got := db.Set(c.in.key, c.in.value)
		if got != c.out {
			t.Errorf("Set(%q) == %q, want %q", c.in.key, got, c.out)
		}
	}
}

// Test getting key-values to the DB
func TestGet(t *testing.T) {
	cases := []struct {
		in  KV
		out interface{}
	}{
		{KV{"key1", "value1"}, "value1"},
		{KV{"key2", "value2"}, "value2"},
		{KV{"key 3", "value 3"}, "value 3"},
		{KV{"மொழி", "தமிழ்"}, "தமிழ்"},
		{KV{"key1", "new value 1"}, "new value 1"},
		{KV{"tested", true}, true},
		{KV{"test_num", 7}, 7},
		{KV{"PI", 3.14}, 3.14},
	}
	db := setUp()
	for _, c := range cases {
		db.Set(c.in.key, c.in.value)
		got := db.Get(c.in.key)
		if got != c.out {
			t.Errorf("Get(%q) == %q, want %q", c.in.key, got, c.out)
		}
	}
}
