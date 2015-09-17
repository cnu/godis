package godis

import (
	"testing"
	"time"
)

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

var integers = []Case{
	{"int", 234},
	{"long", 223344231},
	//{"bool", true}, // boolean value
	{"negative", -554},
	{"zero", 0},
}

var floats = []Case{
	{"float64", 22423234.1223},
	{"float", 443.21},
}

var strings = []Case{
	{"மொழி", "தமிழ்"}, // value with string
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

// Test MGETting key-values from the DB
func TestMGET(t *testing.T) {
	testKeys := []string{"key1", "key2", "key 3"}
	want := []string{"new value 1", "value2", "value 3"}
	db := setUp()
	for _, c := range cases {
		db.SET(c.key, c.value)
	}
	got := db.MGET(testKeys...)
	for i, key := range testKeys {
		if got[i] != want[i] {
			t.Errorf("MGET(%q) == %q, want %q", key, got[i], want[i])
		}
	}
}

// Test MGETting non-existent key-values from the DB
func TestMGETNotExists(t *testing.T) {
	testKeys := []string{"non-key1", "non-key2", "non-key3"}
	db := setUp()
	for _, c := range cases {
		db.SET(c.key, c.value)
	}
	got := db.MGET(testKeys...)
	for i, key := range testKeys {
		if got[i] != nil {
			t.Errorf("MGET(%q) == %q, want %p", key, got[i], nil)
		}
	}

}

// Test MGETting few non-existent key-values from the DB
func TestMGETFewNotExists(t *testing.T) {
	testKeys := []string{"key1", "non-key2", "key 3"}
	want := []string{"new value 1", "value2", "value 3"}
	db := setUp()
	for _, c := range cases {
		db.SET(c.key, c.value)
	}
	got := db.MGET(testKeys...)
	for i, key := range testKeys {
		if i == 1 && got[i] != nil {
			t.Errorf("MGET(%q) == %q, want %p", key, got[i], nil)
		}
		if i != 1 && got[i] != want[i] {
			t.Errorf("MGET(%q) == %q, want %q", key, got[i], want[i])
		}
	}

}

// Test MSETing key-values pairs into the DB
func TestMSET(t *testing.T) {
	tests := []string{"key1", "value1", "key2", "value2", "key3", "value3"}
	db := setUp()
	got := db.MSET(tests...)
	if got != true {
		t.Errorf("MSET(%q) == %t, want %t", tests, got, true)
	}
}

// Test incrementing values for given key by 1
func TestINCR(t *testing.T) {
	db := setUp()
	for _, c := range integers {
		db.SET(c.key, c.value)
		got := db.INCR(c.key)
		if got != c.value.(int)+1 {
			t.Errorf("INCR(%q) == %d, want %d", c.key, got, c.value.(int)+1)
		}
	}
}

// Test incrementing non-existent keys
func TestINCRNonExists(t *testing.T) {
	db := setUp()
	got := db.INCR("non-incr-key")
	if got != 1 {
		t.Errorf("INCR(%q) == %d, want %d", "non-incr-key", got, 1)
	}
}

// Test decrementing values for given key by 1
func TestDECR(t *testing.T) {
	db := setUp()
	for _, c := range integers {
		db.SET(c.key, c.value)
		got := db.DECR(c.key)
		if got != c.value.(int)-1 {
			t.Errorf("DECR(%q) == %d, want %d", c.key, got, c.value.(int)-1)
		}
	}
}

// Test incrementing non-existent keys
func TestDECRNonExists(t *testing.T) {
	db := setUp()
	got := db.DECR("non-decr-key")
	if got != -1 {
		t.Errorf("DECR(%q) == %d, want %d", "non-decr-key", got, -1)
	}
}

// Test incrementing values for given key by n
func TestINCRBY(t *testing.T) {
	db := setUp()
	n := 3
	for _, c := range integers {
		db.SET(c.key, c.value)
		got := db.INCRBY(c.key, n)
		if got != c.value.(int)+n {
			t.Errorf("INCRBY(%q) == %d, want %d", c.key, got, c.value.(int)+n)
		}
	}
}

// Test decrementing values for given key by n
func TestDECRBY(t *testing.T) {
	db := setUp()
	n := 3
	for _, c := range integers {
		db.SET(c.key, c.value)
		got := db.DECRBY(c.key, n)
		if got != c.value.(int)-n {
			t.Errorf("DECRBY(%q) == %d, want %d", c.key, got, c.value.(int)-n)
		}
	}
}

/* Test incrementing values for given key by 1
func TestINCRmismatchs(t *testing.T) {
	db := setUp()
	n := 1
	for _, c := range strings {
		db.SET(c.key, c.value)
		got := db.INCR(c.key)
		_, ok := c.value.(int)
		//if ok && got != value+n {
		//	t.Errorf("INCR(%q) == %d, want %d", c.key, got, value+n)
		//}
		if !ok {
			t.Errorf("INCR(%q) got type %v, want type %v", c.key, reflect.TypeOf(got), reflect.TypeOf(n))
		}
	}
}*/

func TestSETEXWithinExp(t *testing.T) {
	// One second before expiry time
	key := "mykey"
	val := 25
	exp := 3
	db := setUp()
	db.SETEX(key, int64(exp), val)
	time.Sleep(time.Duration(exp-1) * time.Second)
	got := db.GET(key)
	if got != val {
		t.Errorf("SETEX(%q) == %d, want %d", key, got, val)
	}
}

func TestSETEXAfterExp(t *testing.T) {
	// One second before expiry time
	key := "mykey"
	val := 25
	exp := 3
	db := setUp()
	db.SETEX(key, int64(exp), val)
	time.Sleep(time.Duration(exp+1) * time.Second)
	got := db.GET(key)
	if got != nil {
		t.Errorf("SETEX(%q) == %d, want nil", key, got)
	}
}
