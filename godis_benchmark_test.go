package godis

import "testing"

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
