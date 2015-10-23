package godis

import (
	"strconv"
	"testing"
	"time"
)

// Test setting key-values to the DB
func TestSET(t *testing.T) {

	db := setUp()
	for _, c := range cases {
		got, err := db.SET(c.key, c.value)
		if got != c.key || err != nil {
			t.Errorf("SET(%q) == %q, want %q", c.key, got, c.key)
		}
	}
}

// Test getting key-values from the DB
func TestGET(t *testing.T) {
	db := setUp()
	for _, c := range cases {
		db.SET(c.key, c.value)
		got, err := db.GET(c.key)
		if err != nil && got != c.value {
			t.Errorf("GET(%q) == %q, %v want %q, <nil>", c.key, got, err, c.value)
		}
	}
}

// Test getting non-existent key from the DB
func TestGETNotEXISTS(t *testing.T) {
	db := setUp()
	key := "not exists"
	got, err := db.GET(key)
	if err.Error() != "keynotexists" {
		// Means the key exists
		t.Errorf("GET(%q) == %s, %v want %s, keynotexists", key, got, err, "")
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
		got, _ := db.INCR(c.key)
		want, _ := strconv.Atoi(c.value)
		wantStr := want + 1
		if got != wantStr {
			t.Errorf("INCR(%q) == %s, want %s", c.key, got, wantStr)
		}
	}
}

// Test incrementing non-existent keys
func TestINCRNonExists(t *testing.T) {
	db := setUp()
	got, _ := db.INCR("non-incr-key")
	if got != "1" {
		t.Errorf("INCR(%q) == %s, want %s", "non-incr-key", got, "1")
	}
}

// Test decrementing values for given key by 1
func TestDECR(t *testing.T) {
	db := setUp()
	for _, c := range integers {
		db.SET(c.key, c.value)
		got, _ := db.DECR(c.key)
		want, _ := strconv.Atoi(c.value)
		wantStr := strconv.Itoa(want - 1)
		if got != wantStr {
			t.Errorf("DECR(%q) == %s, want %s", c.key, got, wantStr)
		}
	}
}

// Test decrementing non-existent keys
func TestDECRNonExists(t *testing.T) {
	db := setUp()
	got, _ := db.DECR("non-decr-key")
	if got != "-1" {
		t.Errorf("DECR(%q) == %s, want %s", "non-decr-key", got, "-1")
	}
}

// Test incrementing values for given key by n
func TestINCRBY(t *testing.T) {
	db := setUp()
	n := 3
	for _, c := range integers {
		db.SET(c.key, c.value)
		got, _ := db.INCRBY(c.key, n)
		want, _ := strconv.Atoi(c.value)
		wantStr := strconv.Itoa(want + n)
		if got != wantStr {
			t.Errorf("INCRBY(%q) == %s, want %s", c.key, got, wantStr)
		}
	}
}

// Test incrementing values for a string value
func TestINCRBYString(t *testing.T) {
	db := setUp()
	n := 3
	key := "foo"
	db.SET(key, "string value")
	_, err := db.INCRBY(key, n)
	if !err {
		t.Errorf("INCRBY(%q, %d) == %t, want %t", key, n, err, true)
	}
}

// Test decrementing values for given key by n
func TestDECRBY(t *testing.T) {
	db := setUp()
	n := 3
	for _, c := range integers {
		db.SET(c.key, c.value)
		got, _ := db.DECRBY(c.key, n)
		want, _ := strconv.Atoi(c.value)
		wantStr := strconv.Itoa(want - n)
		if got != wantStr {
			t.Errorf("DECRBY(%q) == %s, want %s", c.key, got, wantStr)
		}
	}
}

// Test decrementing values for a string value
func TestDECRBYString(t *testing.T) {
	db := setUp()
	n := 3
	key := "foo"
	db.SET(key, "string value")
	_, err := db.DECRBY(key, n)
	if !err {
		t.Errorf("DECRBY(%q, %d) == %t, want %t", key, n, err, true)
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
        //  t.Errorf("INCR(%q) == %d, want %d", c.key, got, value+n)
        //}
        if !ok {
            t.Errorf("INCR(%q) got type %v, want type %v", c.key, reflect.TypeOf(got), reflect.TypeOf(n))
        }
    }
}*/

func TestSETEXWithinExp(t *testing.T) {
	// One second before expiry time
	key := "mykey"
	val := "some value"
	exp := 2
	db := setUp()
	got, _ := db.SETEX(key, uint64(exp), val)
	time.Sleep(time.Duration(exp-1) * time.Second)
	res, _ := db.EXISTS(key)
	if res != 1 {
		t.Errorf("SETEX(%q, %d, %v) == %d, <nil> want %d, <nil>", key, exp, val, got, 1)
	}
}

func TestSETEXAfterExp(t *testing.T) {
	// One second before expiry time
	key := "mykey"
	val := "some value"
	exp := 1
	db := setUp()
	db.SETEX(key, uint64(exp), val)
	time.Sleep(time.Duration(exp+1) * time.Second)
	got, _ := db.EXISTS(key)
	if got != 0 {
		t.Errorf("SETEX(%q, %d) == %d, want %d", key, exp, got, 0)
	}
}

func TestSETEXWithZero(t *testing.T) {
	// Zero as expiry time
	key := "mykey"
	val := "some value"
	exp := 0
	db := setUp()
	_, err := db.SETEX(key, uint64(exp), val)
	if err == nil {
		t.Errorf("SETEX(%q, %d) == nil, want Error(\"invalid expire time in SETEX\")", key, exp)
	}
}

func TestPSETEXWithinExp(t *testing.T) {
	// One second before expiry time
	key := "mykey"
	val := "some value"
	exp := 1000
	db := setUp()
	db.PSETEX(key, uint64(exp), val)
	time.Sleep(time.Duration(exp-10) * time.Millisecond)
	got, _ := db.EXISTS(key)
	if got != 1 {
		t.Errorf("PSETEX(%q, %d) == %d, want %d", key, exp, got, 1)
	}
}

func TestPSETEXAfterExp(t *testing.T) {
	// One second before expiry time
	key := "mykey"
	val := "some value"
	exp := 1000
	db := setUp()
	db.PSETEX(key, uint64(exp), val)
	time.Sleep(time.Duration(exp+10) * time.Millisecond)
	got, _ := db.EXISTS(key)
	if got != 0 {
		t.Errorf("PSETEX(%q) == %d, want %d", key, got, 0)
	}
}

func TestPSETEXWithZero(t *testing.T) {
	// Zero as expiry time
	key := "mykey"
	val := "some value"
	exp := 0
	db := setUp()
	_, err := db.PSETEX(key, uint64(exp), val)
	if err == nil {
		t.Errorf("PSETEX(%q, %d) == nil, want Error(\"invalid expire time in PSETEX\")", key, exp)
	}
}

// STRLEN should return length of a int
func TestSTRLENWithInt(t *testing.T) {
	db := setUp()
	key := "mykey"
	val := "12345"
	db.SET(key, val)
	got, err := db.STRLEN(key)
	if err != nil || got != 5 {
		t.Errorf("STRLEN(%q) == %d, want %d", key, got, 5)
	}
}

// STRLEN should return length of a float
func TestSTRLENWithFloat(t *testing.T) {
	db := setUp()
	key := "mykey"
	val := "12345.54321"
	db.SET(key, val)
	got, err := db.STRLEN(key)
	if err != nil || got != 11 {
		t.Errorf("STRLEN(%q) == %d, want %d", key, got, 11)
	}
}

// STRLEN should return length of a long
func TestSTRLENWithLong(t *testing.T) {
	db := setUp()
	key := "mykey"
	val := "12345.5432123453453463423434344652123234235"
	db.SET(key, val)
	got, err := db.STRLEN(key)
	if err != nil || got != 43 {
		t.Errorf("STRLEN(%q) == %d, want %d", key, got, 43)
	}
}

// STRLEN should return length of a string
func TestSTRLENWithStr(t *testing.T) {
	db := setUp()
	key := "mykey"
	val := "A quick brown fox jumped over a lazy dog and broke its leg"
	db.SET(key, val)
	got, err := db.STRLEN(key)
	if err != nil || got != 58 {
		t.Errorf("STRLEN(%q) == %d, want %d", key, got, 58)
	}
}

// STRLEN should return Zero if given key is empty.
func TestSTRLENWithoutVal(t *testing.T) {
	db := setUp()
	key := "mykey"
	val := ""
	db.SET(key, val)
	got, err := db.STRLEN(key)
	if err != nil || got != 0 {
		t.Errorf("STRLEN(%q) == %d, want %d", key, got, 0)
	}
}

// STRLEN should return 0, keynotfound error if given key does not exist.
func TestSTRLENWithoutKey(t *testing.T) {
	db := setUp()
	key := "mykey"
	got, err := db.STRLEN(key)
	if err.Error() != "keynotfound" {
		t.Errorf("STRLEN(%q) == %v,%v want 0,keynotfound", key, got, err)
	}
}

// TODO : Write test cases for STRLEN in type mismatch after data structs are done.
