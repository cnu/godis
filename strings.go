package godis

import (
	"errors"
	"reflect"
	"strconv"
	"unicode/utf8"
)

// SET is used to assign a value to a key
func (g *Godis) SET(key string, value string) string {
	if g.EXISTS(key) == 1 {
		s, exists := g.getSDS(key)
		if exists {
			s.Lock()
			defer s.Unlock()
			s.value = value
		}
	} else {
		s := NewSDS(value)
		g.db[key] = s
	}
	return key
}

// GET returns the value stored for a key
func (g *Godis) GET(key string) (string, bool) {
	s, exists := g.getSDS(key)
	if exists {
		return s.get(), false
	}
	return "", true
}

// SETEX is used to assign a value to a key and destroy it within its given
//expiry time in seconds
func (g *Godis) SETEX(key string, exp uint64, value string) (string, error) {
	if exp <= 0 {
		return "", errors.New("invalid expire time in SETEX")
	}
	g.SET(key, value)
	go g.destroyInSecs(key, exp)
	return key, nil
}

// PSETEX is used to assign a value to a key and destroy it within its given
// expiry time in milliseconds
func (g *Godis) PSETEX(key string, exp uint64, value string) (string, error) {
	if exp <= 0 {
		return "", errors.New("invalid expire time in PSETEX")
	}
	g.SET(key, value)
	go g.destroyInMillis(key, exp)
	return key, nil
}

// INCR increments the key by one
func (g *Godis) INCR(key string) (string, bool) {
	return g.INCRBY(key, 1)
}

// DECR decrements the key by one
func (g *Godis) DECR(key string) (string, bool) {
	return g.DECRBY(key, 1)
}

// INCRBY increments the key by given value
func (g *Godis) INCRBY(key string, n int) (string, bool) {
	if g.EXISTS(key) == 0 {
		g.SET(key, "0")
	}
	val, err := g.GET(key)
	if !err {
		valInt, convErr := strconv.Atoi(val)
		if convErr == nil {
			valInt += n
			g.SET(key, strconv.Itoa(valInt))
			return strconv.Itoa(valInt), false
		}
	}
	return "", true
}

// DECRBY decrements the key by given value
func (g *Godis) DECRBY(key string, n int) (string, bool) {
	return g.INCRBY(key, -n)
}

// MGET returns a slice of values for a input slice of keys
func (g *Godis) MGET(keys ...string) []interface{} {
	var output []interface{} // will be strings or nils
	for _, key := range keys {
		value, err := g.GET(key)
		if !err {
			output = append(output, value)
		} else {
			// if the key isn't available, append a nil instead
			// TODO: how to do multi-return way if we have a slice to return?
			output = append(output, nil)
		}
	}
	return output
}

// MSET sets a slice of key-values
// pass a slice of keys and value alternating
// eg: "key1", "value1", "key2", "value2"
func (g *Godis) MSET(items ...string) bool {
	g.Lock()
	defer g.Unlock()
	for i, item := range items {
		if i%2 == 1 {
			continue
		}
		g.SET(item, items[i+1])
	}
	return true
}

//STRLEN returns the length of the string value stored at key.
//An error is returned when key holds a non-string value.
func (g *Godis) STRLEN(key string) (int64, error) {
	val, err := g.GET(key)
	if err {
		return 0, errors.New("keynotfound")
	}
	if reflect.ValueOf(val).Kind() != reflect.String {
		return 0, errors.New("typemismatch")
	}
	return int64(utf8.RuneCountInString(val)), nil
}
