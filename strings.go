package godis

import (
	"errors"
	"reflect"
	"strconv"
	"unicode/utf8"
)

// SET is used to assign a value to a key
func (g *Godis) SET(key string, value string) (string, error) {
	if got, _ := g.EXISTS(key); got == 1 {
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
	return key, nil
}

// GET returns the value stored for a key
func (g *Godis) GET(key string) (string, error) {
	s, exists := g.getSDS(key)
	if exists {
		return s.get(), nil
	}
	return "", errors.New("keynotexists")
}

// SETEX is used to assign a value to a key and destroy it within its given
//expiry time in seconds, returns the key,<nil> if set or "",<nil>
func (g *Godis) SETEX(key string, exp uint64, value string) (string, error) {
	if exp <= 0 {
		return "", errors.New("badexpiry")
	}
	g.SET(key, value)
	go g.destroyInSecs(key, exp)
	return key, nil
}

// PSETEX is used to assign a value to a key and destroy it within its given
// expiry time in milliseconds, returns the key,<nil> if set or "",<nil>
func (g *Godis) PSETEX(key string, exp uint64, value string) (string, error) {
	if exp <= 0 {
		return "", errors.New("badexpiry")
	}
	g.SET(key, value)
	go g.destroyInMillis(key, exp)
	return key, nil
}

// INCR increments the key by one
func (g *Godis) INCR(key string) (int, error) {
	return g.INCRBY(key, 1)
}

// DECR decrements the key by one
func (g *Godis) DECR(key string) (int, error) {
	return g.DECRBY(key, 1)
}

// INCRBY increments the key by given value
func (g *Godis) INCRBY(key string, n int) (int, error) {
	var val string
	var err error
	val, err = g.GET(key)
	if val == "" || err != nil {
		g.SET(key, "0")
		val, _ = g.GET(key)
	}
	valInt, convErr := strconv.Atoi(val)
	if convErr != nil {
		return 0, errors.New("typemismatch")
	}
	valInt += n
	g.SET(key, strconv.Itoa(valInt))
	return valInt, nil
}

// DECRBY decrements the key by given value
func (g *Godis) DECRBY(key string, n int) (int, error) {
	return g.INCRBY(key, -n)
}

// MGET returns a slice of values for a input slice of keys
func (g *Godis) MGET(keys ...string) ([]interface{}, error) {
	var output []interface{} // will be strings or nils
	for _, key := range keys {
		value, err := g.GET(key)
		if err == nil {
			output = append(output, value)
		} else {
			// if the key isn't available, append a nil instead
			// TODO: how to do multi-return way if we have a slice to return?
			output = append(output, nil)
		}
	}
	return output, nil
}

// MSET sets a slice of key-values
// pass a slice of keys and value alternating
// eg: "key1", "value1", "key2", "value2"
// returns true, <nil> as MSET never fails!
func (g *Godis) MSET(items ...string) (bool, error) {
	g.Lock()
	defer g.Unlock()
	for i, item := range items {
		if i%2 == 1 {
			continue
		}
		g.SET(item, items[i+1])
	}
	return true, nil
}

//STRLEN returns the length of the string value stored at key.
//An error is returned when key holds a non-string value.
func (g *Godis) STRLEN(key string) (int64, error) {
	val, err := g.GET(key)
	if err.Error() == "keynotfound" {
		return 0, errors.New("keynotfound")
	}
	if reflect.ValueOf(val).Kind() != reflect.String {
		return 0, errors.New("typemismatch")
	}
	return int64(utf8.RuneCountInString(val)), nil
}
